package mysql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var sgbd string = "mysql"
var username string = "root"
var password string = "secret"
var dbname string = "naturalYSalvaje"

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
