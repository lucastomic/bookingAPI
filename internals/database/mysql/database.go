package mysql

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucastomic/naturalYSalvajeRent/internals/enviroment"
)

var sgbd string = enviroment.GetSGBD()
var username string = enviroment.GetDatabaseUser()
var password string = enviroment.GetDatabasePassword()
var dbname string = enviroment.GetDatabaseName()

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
	var dataSource string = fmt.Sprintf("%s:%s@()/%s", username, password, dbname)

	db, err := sql.Open(sgbd, dataSource)

	if err != nil {
		panic(err)
	}

	return db
}
