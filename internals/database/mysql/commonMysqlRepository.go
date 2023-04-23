package mysql

import (
	"fmt"

	"github.com/lucastomic/naturalYSalvajeRent/internals/domain/boat"
)

// CommonMysqlRepository is a template for all the interactions with the database wihch are common in all
// the repositories, regardless of the domain class they handle.
// It expects two generics parameters, T is the object the repository will handle
// and I is the type of the object Id
// For example, for BoatRepository it would be [T=Boat, I=int] (knowing that the Id
// of the Boat object is an integer)
type CommonMysqlRepository[T any, I any] struct {
	IPrimitiveRepoBehaivor[T, I]
}

// Save persists the object specified as argument into the database. If it exists,
// it updates the modified values, if it doesn't it creates a new register
func (repo CommonMysqlRepository[T, I]) Save(object T) error {
	var err error
	if repo.alreadyExists(object) {
		err = repo.update(object)
	} else {
		err = repo.insertNew(object)
	}
	return err
}

// insertNew creates a new DB register given a T object
func (repo CommonMysqlRepository[T, I]) insertNew(object T) error {
	db := getInstance()
	stmt, err := db.Prepare(repo.insertStmt())
	if err == nil {
		defer stmt.Close()
		values := repo.persistenceValues(object)
		_, err = stmt.Exec(values...)
	}
	return err
}

// alreadyExists checks whether an object has already been inserted into the DB
// Returns true if exists a row with the id of the object passed as argument, or false otherwise
func (repo CommonMysqlRepository[T, I]) alreadyExists(object T) bool {
	id := repo.id(object)
	dbObject, err := repo.FindById(id...)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return !repo.isZero(dbObject)
}

// update updates the modified values of an object in the database
// If an error ocurrs it returns it, otherwise return nil
func (repo CommonMysqlRepository[T, I]) update(object T) error {
	db := getInstance()
	stmt, err := db.Prepare(repo.updateStmt())
	if err != nil {
		return err
	}
	persistenceValues := repo.persistenceValues(object)
	ids := repo.id(object)
	persistenceValues = repo.concatId(persistenceValues, ids)

	_, err = stmt.Exec(persistenceValues...)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// concatId takes a slice and concatenates the elements inside a IDs slice
func (repo CommonMysqlRepository[T, I]) concatId(originalSlice []any, ids []I) []any {
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
func (repo CommonMysqlRepository[T, I]) FindById(id ...I) (T, error) {
	var response T
	db := getInstance()
	stmt, err := db.Prepare(repo.findByIdStmt())
	if err != nil {
		return *repo.empty(), err
	}
	defer stmt.Close()
	idsParsed := repo.parseISliceToAnySlice(id)
	rows, err := stmt.Query(idsParsed...)
	if err != nil {
		return *repo.empty(), err
	}
	if rows.Next() {
		response, err = repo.scan(rows)
		if err != nil {
			return *repo.empty(), err
		}
	}

	err = repo.updateRelations(&response)
	if err != nil {
		return *repo.empty(), err
	}

	return response, nil
}

// parseISliceToAnySlice parses a slice []I into a []any slice
func (repo CommonMysqlRepository[T, I]) parseISliceToAnySlice(slice []I) []any {
	response := make([]any, len(slice))
	for i, val := range slice {
		response[i] = val
	}
	return response
}

func newBoatCommonMysqlRepository() CommonMysqlRepository[boat.Boat, int] {
	return CommonMysqlRepository[boat.Boat, int]{boatPrimitiveRepoBehaivor{}}
}
func newStateRoomCommonMysqlRepository() CommonMysqlRepository[boat.StateRoom, int] {
	return CommonMysqlRepository[boat.StateRoom, int]{stateRoomPrimitiveRepoBehaivor{}}
}
