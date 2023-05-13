package mysql

import (
	"database/sql"
	"fmt"
)

// CommonMysqlLogic is a template for all the interactions with the database wihch are common in all
// the repositories, regardless of the domain class they handle.
// It expects two generics parameters, T is the object the repository will handle
// and I is the type of the object Id
// For example, for BoatRepository it would be [T=Boat, I=int] (knowing that the Id
// of the Boat object is an integer)
type CommonMysqlLogic[T any, I any] struct {
	IPrimitiveRepoBehaivor[T, I]
}

// Save persists the object specified as argument into the database. If it exists,
// it updates the modified values, if it doesn't it creates a new register
// It also persists the changes/insertions of all its childs.
// For example, given
//
//	type A struct{
//		name string
//		childs []B
//	}
//
// Save(A) would persist the object A and its field "name", but it also would persist,
// all the changes/insertions in their childs (the slice of B objects)
func (repo CommonMysqlLogic[T, I]) Save(object T) error {
	var err error
	if repo.alreadyExists(object) {
		err = repo.update(object)
	} else {
		err = repo.insertNew(object)
	}
	if err != nil {
		return err
	}
	err = repo.SaveChildsChanges(&object)

	return err
}

// insertNew creates a new DB register given a T object
func (repo CommonMysqlLogic[T, I]) insertNew(object T) error {
	db := GetInstance()
	stmt, err := db.Prepare(repo.InsertStmt())
	if err == nil {
		defer stmt.Close()
		values := repo.PersistenceValues(object)
		_, err = stmt.Exec(values...)
	}
	return err

}

// alreadyExists checks whether an object has already been inserted into the DB
// Returns true if exists a row with the id of the object passed as argument, or false otherwise
func (repo CommonMysqlLogic[T, I]) alreadyExists(object T) bool {
	id := repo.Id(object)
	dbObject, err := repo.FindById(id...)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return !repo.IsZero(dbObject)
}

// update updates the modified values of an object in the database
// If an error ocurrs it returns it, otherwise return nil
func (repo CommonMysqlLogic[T, I]) update(object T) error {
	db := GetInstance()
	stmt, err := db.Prepare(repo.UpdateStmt())
	if err != nil {
		return err
	}
	persistenceValues := repo.PersistenceValues(object)
	ids := repo.Id(object)
	persistenceValues = repo.concatId(persistenceValues, ids)

	_, err = stmt.Exec(persistenceValues...)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// concatId takes a slice and concatenates the elements inside a IDs slice
func (repo CommonMysqlLogic[T, I]) concatId(originalSlice []any, ids []I) []any {
	var response []any = originalSlice
	for _, id := range ids {
		response = append(response, id)
	}
	return response
}

// FindById returns an object given its Id.
// If the wanted row has more than one id (a compound ID), both IDs must be passed as argument
// (for example, FindById(id1,id2))
// If an error ocurrs, it returns an empty object with the error as second
// value. If no error ocurrs, it returns the object as first parameter and nil as second
func (repo CommonMysqlLogic[T, I]) FindById(id ...I) (T, error) {
	idsParsed := repo.parseISliceToAnySlice(id)
	stmt := repo.FindByIdStmt()
	response, err := repo.Query(stmt, idsParsed)
	if err != nil || len(response) == 0 {
		return *repo.Empty(), err
	}
	return response[0], nil
}

// GetAll retrieves all the T objects form the database
// In case of error, it returns an empty slice and the error.
func (repo CommonMysqlLogic[T, I]) FindAll() ([]T, error) {
	stmt := repo.FindAllStmt()
	response, err := repo.Query(stmt, []any{})
	if err != nil || len(response) == 0 {
		return []T{}, err
	}
	return response, nil

}

// parseISliceToAnySlice parses a slice []I into a []any slice
func (repo CommonMysqlLogic[T, I]) parseISliceToAnySlice(slice []I) []any {
	response := make([]any, len(slice))
	for i, val := range slice {
		response[i] = val
	}
	return response
}

// Query retrieves an slice of T objects given a Query statement and the params to execute it.
// In case of error, it returns an empty slice and the error.
// It's important that the Query Statement passed as argument, must return all the columns of the entity

// For example, given the next statement
// SELECT * FROM cats WHERE color = ? AND age = ?
// And the next params
// []any{"Orange", 4}
// Then, it would return a slice with all the orange cats with 4 yeats old
func (repo CommonMysqlLogic[T, I]) Query(queryStmt string, queryParams []any) ([]T, error) {
	var response []T
	db := GetInstance()
	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		return []T{}, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(queryParams...)
	if err != nil {
		return []T{}, err
	}
	for rows.Next() {
		newValue, err := repo.getEntityFromRow(rows)
		if err != nil {
			return []T{}, err
		}
		response = append(response, newValue)
	}

	if err != nil {
		return []T{}, err
	}

	return response, nil
}

// getEntityFromRow retrieves the corrspondient object in the rows and update
// its relations (OneToMany,ManyToMany,etc.)
// If there is an error it returns an empty object and the error
func (repo CommonMysqlLogic[T, I]) getEntityFromRow(rows *sql.Rows) (T, error) {
	response, err := repo.Scan(rows)
	if err != nil {
		return *repo.Empty(), err
	}
	err = repo.UpdateRelations(&response)
	if err != nil {
		return *repo.Empty(), err
	}
	return response, nil
}

// Remove removes a T object from the database
// If an error ocurrs it returns it, otherwise return nil
func (repo CommonMysqlLogic[T, I]) Remove(object T) error {
	db := GetInstance()
	stmt, err := db.Prepare(repo.RemoveStmt())
	if err != nil {
		return err
	}
	ids := repo.Id(object)
	idsParsed := repo.parseISliceToAnySlice(ids)
	_, err = stmt.Exec(idsParsed...)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil

}
