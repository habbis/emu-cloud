package dbconfig

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Jsondata struct {
	Db_host     string
	Db_port     int
	Db_user     string
	Db_password string
	Database    string
}

func Dbconn() (db *sql.DB, err error) {

	dirpath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	var config = ".config.json"

	homedirconf := fmt.Sprintf("%s/%s", dirpath, config)

	content, err := os.ReadFile(homedirconf)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var payload Jsondata
	err = json.Unmarshal(content, &payload)

	mysqlconn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", payload.Db_user, payload.Db_password, payload.Db_host, payload.Db_port, payload.Database)
	db, err = sql.Open("mysql", mysqlconn)

	return

}
