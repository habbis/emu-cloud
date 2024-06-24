package main

import (
	"database/sql"
	"emu-cloud/internal/dbconfig"
	"encoding/json"
	"fmt"
	//"emu-cloud/internal/dbconfig"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
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

//db, err := dbconfig.Dbconn()

type Vlan struct {
	Id      int    `json:"id"` // Database Id, can be safely ignored. Will be removed.
	Site    string `json:"site"`
	Vlan_id int    `json:"vlan_id"`
	Name    string `json:"name"` // Human readable name
	// prefix    types.ipnet
	prefix  string `json:"prefix"`
	Comment string `json:"comment"`
	Zone    string `json:"zone"` // Z1/Z2 etc
	State   string `json:"state"`
	//Updated   *types.Time
}

func Get_vlan(conn *gin.Context) {
	db, err := dbconfig.Dbconn()

	var (
		id_vlan int
		site    string
		vlan_id string
		name    string
		prefix  string
		comment string
		zone    string
		state   string
	)

	rows, err := db.Query("SELECT id,site,vlan_id,name,prefix,comment,zone,state FROM vlan")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {

		if err := rows.Scan(&id_vlan, &site, &vlan_id, &name, &prefix, &comment, &zone, &state); err != nil {
			log.Fatal(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		var vlan = []Vlan{
			{Id: id_vlan, Site: site, Vlan_id: 22, Name: name, prefix: prefix, Comment: comment, Zone: zone, State: state},
		}

		//conn.IndentedJSON(http.StatusOK, vlan)
		conn.PureJSON(http.StatusOK, vlan)

	}

}

func main() {
	router := gin.Default()
	router.GET("/api/vlan/", Get_vlan)
	//router.GET("/api/vlan/", Get_vlan)
	router.Run("localhost:8080")

}
