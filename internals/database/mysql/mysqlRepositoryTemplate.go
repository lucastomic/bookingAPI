package mysql

// MysqlRepository is a template for the interactions with the MYSQL databse.
// It expects two generics parameters, T is the object the repository will handle
// and I is the type of the object Id
// For example, for BoatRepository it would be [T=Boat, I=int] (knowing that the Id
// of the Boat object is an integer)
type MysqlRepository[T any, I any] struct {
	concretRepository ConcretRepository[T, I]
}

// Save persists the object specified as argument into the database. If it exists,
// it updates the modified values, if it doesn't it creates a new register
func (repo MysqlRepository[T, I]) Save(object T) error {
	if repo.alreadyExists(object) {
		err := repo.update(object)
		return err
	}
	err := repo.insertNew(object)
	return err
}

// insertNew creates a new DB register given an object
func (repo MysqlRepository[T, I]) insertNew(object T) error {
	db := getInstance()
	stmt, err := db.Prepare(repo.concretRepository.insertStmt())
	if err != nil {
		return err
	}
	defer stmt.Close()
	values := repo.concretRepository.persistenceValues(object)
	_, err = stmt.Exec(values...)
	return err
}

// alreadyExists checks whether an object has already been inserted into the DB
// Returns true if exists a row with the id of the object passed as argument, or false otherwise
func (repo MysqlRepository[T, I]) alreadyExists(object T) bool {
	id := repo.concretRepository.id(object)
	_, err := repo.FindById(id)
	return err == nil
}

// update updates the modified values of an object in the database
// If an error ocurrs it returns it, otherwise return nil
func (repo MysqlRepository[T, I]) update(object T) error {
	db := getInstance()
	stmt, err := db.Prepare(repo.concretRepository.updateStmt())
	if err != nil {
		return err
	}
	values := repo.concretRepository.persistenceValues(object)
	_, err = stmt.Exec(values...)
	defer stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

// FindById returns an object given its Id.
// If an error ocurrs, it returns an empty object with the error as second
// value. If no error ocurrs, it returns the object as first parameter and nil as second
func (repo MysqlRepository[T, I]) FindById(id I) (T, error) {
	var response T
	db := getInstance()
	stmt, err := db.Prepare(repo.concretRepository.findByIdStmt())
	if err != nil {
		return *repo.concretRepository.empty(), err
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		return *repo.concretRepository.empty(), err
	}
	rows.Scan(response)
	return response, nil
}
