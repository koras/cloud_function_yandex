package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Body       interface{} `json:"body"`
}

type Features struct {
	gorm.Model
	ID     int64  `gorm:"column:id" json:"id"`
	Screen string `gorm:"column:screen" json:"screen"`
	Name   string `gorm:"column:name" json:"name"`
}

var dbSettings struct {
	host     string
	port     int16
	user     string
	password string
	dbname   string
}

func Connect() (*sql.DB, error) {

	dbSettings.user = os.Getenv("CLOUD_PG_USER")
	dbSettings.host = os.Getenv("CLOUD_PG_HOST")
	dbSettings.dbname = os.Getenv("CLOUD_PG_DBNAME")
	dbSettings.port = 6432
	dbSettings.password = os.Getenv("CLOUD_PG_DBPASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s statement_cache_mode=describe search_path=public  sslmode=require binary_parameters=yes",
		dbSettings.host, dbSettings.port, dbSettings.user, dbSettings.password, dbSettings.dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}
	return db, nil
}
