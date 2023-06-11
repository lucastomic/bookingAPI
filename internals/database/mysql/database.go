package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var sgbd string = "mysql"
var username string = os.Getenv("MYSQL_USER")
var password string = os.Getenv("MYSQL_PASSWORD")
var dbname string = os.Getenv("MYSQL_DATABASE")

var lock = &sync.Mutex{}

var instance *sql.DB

func GetInstance() *sql.DB {
	lock.Lock()

	defer lock.Unlock()

	if instance == nil {
		instance = initDatabase()
	}

	return instance
}

func initDatabase() *sql.DB {
	var dataSource string = fmt.Sprintf("%s:%s@(mysql)/%s", username, password, dbname)

	db, err := sql.Open(sgbd, dataSource)

	if err != nil {
		panic(err)
	}

	return db
}
