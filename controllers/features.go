package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type Features struct {
	gorm.Model
	ID        int64  `gorm:"column:id" json:"id"`
	Screen    string `gorm:"column:screen" json:"screen"`
	Name      string `gorm:"column:name" json:"name"`
	Feature   string `gorm:"column:feature" json:"feature"`
	SectionID string `gorm:"column:section_id" json:"section_id"`
	Comment   string `gorm:"column:comment" json:"comment"`
}

// save
func FeatureSave(db *sql.DB, query url.Values) Features {
	lastInsertId := 0
	sqlStatement := `INSERT INTO "features" ("screen", "name", "feature", "comment", "section_id") values ($1, $2, $3, $4, $5) RETURNING id`

	name := query.Get("name")
	screen := query.Get("screen")
	feature := query.Get("feature")
	section_id := query.Get("section_id")
	comment := query.Get("comment")
	fmt.Printf("comment %s,%s,%s,%s,%s \n\n", screen, name, feature, comment, section_id)

	err := db.QueryRow(sqlStatement, screen, name, feature, comment, section_id).Scan(&lastInsertId)
	if err != nil {
		panic(err)
	}
	return FeatureGet(db, lastInsertId)
}

// func FeaturesList(db *sql.DB) (*services.Response, error) {
func FeaturesList(db *sql.DB) []Features {

	features := []Features{}
	rows, err := db.Query("select id, name, screen, feature, comment, section_id  from features")
	//rows, err := db.Query("select id, name, screen from features")
	if err != nil {
		panic(err)
	}
	// Обработка результатов запроса
	for rows.Next() {
		var featuresSingle Features
		err = rows.Scan(&featuresSingle.ID, &featuresSingle.Name, &featuresSingle.Screen, &featuresSingle.Feature, &featuresSingle.Comment, &featuresSingle.SectionID)
		if err != nil {
			fmt.Printf("err")
			log.Fatal(err)
		}
		features = append(features, featuresSingle)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return features
}

// func FeaturesList(db *sql.DB) (*services.Response, error) {
func FeatureGet(db *sql.DB, id int) Features {
	sqlStatement := `select id, name, screen, feature, section_id, comment from features where id = $1`

	featuresSingle := Features{}

	err := db.QueryRow(sqlStatement, id).Scan(&featuresSingle.ID, &featuresSingle.Name, &featuresSingle.Screen, &featuresSingle.Feature, &featuresSingle.Comment, &featuresSingle.SectionID)
	if err != nil {

		fmt.Printf("err 22")
		fmt.Printf("err")
		log.Fatal(err)
	}

	return featuresSingle
}
