package database

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var sgbd string = "mysql"
var password string = "root"
var username string = "root"
var dbname string = "naturalYSalvajeRent"

var lock = &sync.Mutex{}

var instance *sql.DB

func getInstance() *sql.DB {
	lock.Lock()

	defer lock.Unlock()

	if instance == nil {
		instance = initDatabase()
	}

	return instance
}

func initDatabase() *sql.DB {
	var dataSource string = fmt.Sprintf("%s:%s@/%s", username, password, dbname)

	db, err := sql.Open(sgbd, dataSource)

	if err != nil {
		panic(err)
	}

	return db
}
