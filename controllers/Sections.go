package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Sections struct {
	gorm.Model
	ID       int64  `gorm:"column:id" json:"id"`
	ParentId string `gorm:"column:parent_id" json:"parent_id"`
	Name     string `gorm:"column:name" json:"name"`
	Code     string `gorm:"column:code" json:"code"`
}

func SectorsList(db *sql.DB) []Sections {

	features := []Sections{}
	rows, err := db.Query("select id, name, code, parent_id from sections")
	if err != nil {
		panic(err)
	}
	// Обработка результатов запроса
	for rows.Next() {
		var sectionSingle Sections
		err = rows.Scan(&sectionSingle.ID, &sectionSingle.Name, &sectionSingle.Code, &sectionSingle.ParentId)
		if err != nil {
			fmt.Printf("err")
			log.Fatal(err)
		}
		features = append(features, sectionSingle)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return features
}

// данные которые приходят с фронта
type SectionInput struct {
	//	EventID      string `json:"event_id"`
	Name     string `form:"name" binding:"required"`
	Code     string `form:"code" binding:"required"`
	ParentId string `form:"parent_id" binding:"required"`
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// save
func SectionSave(db *sql.DB, req *http.Request) Sections {

	lastInsertId := 0
	sqlStatement := `INSERT INTO "sections" ("parent_id", "name", "code") values ($1, $2, $3) RETURNING id`
	sqlStatementUpdate := `UPDATE "sections" SET "code"=$2 , "name"=$3 , "parent_id"=$4 WHERE  "id"=$1`
	id := req.FormValue("id")
	name := req.FormValue("name")
	code := req.FormValue("code")
	parent_id := req.FormValue("parent_id")

	fmt.Printf("comment %s,%s,%s \n\n", parent_id, name, code)
	if id != "" {

		_, err := db.Exec(sqlStatementUpdate, id, code, name, parent_id)

		if err != nil {
			fmt.Printf("err  %s \n\n", err)
			panic(err)
		}
		intVar, errs := strconv.Atoi(id)
		if errs != nil {
			panic(errs)
		}
		return SectionGet(db, intVar)
	} else {
		err := db.QueryRow(sqlStatement, parent_id, name, code).Scan(&lastInsertId)
		if err != nil {
			panic(err)
		}

		return SectionGet(db, lastInsertId)
	}
}

// func FeaturesList(db *sql.DB) (*services.Response, error) {
func SectionGet(db *sql.DB, id int) Sections {
	sqlStatement := `select id, "parent_id", "name", "code" from sections where id = $1`

	SectionsSingle := Sections{}

	err := db.QueryRow(sqlStatement, id).Scan(&SectionsSingle.ID, &SectionsSingle.Name, &SectionsSingle.Code, &SectionsSingle.ParentId)
	if err != nil {
		log.Fatal(err)
	}

	return SectionsSingle
}
