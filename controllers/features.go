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

type Features struct {
	gorm.Model
	ID             int64  `gorm:"column:id" json:"id"`
	Screen         string `gorm:"column:screen" json:"screen"`
	Name           string `gorm:"column:name" json:"name"`
	Feature        string `gorm:"column:feature" json:"feature"`
	SectionID      string `gorm:"column:section_id" json:"section_id"`
	Comment        string `gorm:"column:comment" json:"comment"`
	Scenario       string `gorm:"column:scenario" json:"scenario"`
	Code           string `gorm:"column:code" json:"code"`
	AndroidVersion string `gorm:"column:android_version" json:"android_version"`
	ViIosVersion   string `gorm:"column:vi_ios_version" json:"vi_ios_version"`

	AndroidApplicable *bool  `gorm:"column:android_applicable" json:"android_applicable"`
	AndroidActive     *bool  `gorm:"column:android_active" json:"android_active"`
	ViIosApplicable   *bool  `gorm:"column:vi_ios_applicable" json:"vi_ios_applicable"`
	ViIosActive       *bool  `gorm:"column:vi_ios_active" json:"vi_ios_active"`
	Organization      string `gorm:"column:organization" json:"organization"`

	SectionsCode string `gorm:"column:sections_code" json:"sections_code"`
	SectionsName string `gorm:"column:sections_name" json:"sections_name"`
}

// save
func FeatureSave(db *sql.DB, query *http.Request) Features {

	lastInsertId := 0

	sqlStatementUpdate := `UPDATE "features" SET  "screen"=$2,"name"=$3,"feature"=$4,"comment"=$5, "section_id"=$6,"scenario"=$7,"code"=$8,"android_version"=$9, "vi_ios_version"=$10, "android_applicable"=$11,"android_active"=$12,"vi_ios_applicable"=$13,"vi_ios_active"=$14,"organization"=$15  WHERE  "id"=$1`

	id := query.FormValue("id")
	name := query.FormValue("name")
	screen := query.FormValue("screen")
	feature := query.FormValue("feature")
	sectionID := query.FormValue("section_id")
	comment := query.FormValue("comment")
	organization := query.FormValue("organization")
	androidApplicable, _ := strconv.ParseBool(query.FormValue("android_applicable"))
	androidActive, _ := strconv.ParseBool(query.FormValue("android_active"))
	viIosApplicable, _ := strconv.ParseBool(query.FormValue("vi_ios_applicable"))
	viIosActive, _ := strconv.ParseBool(query.FormValue("vi_ios_active"))

	scenario := query.FormValue("scenario")
	code := query.FormValue("code")
	android_version := query.FormValue("android_version")
	vi_ios_version := query.FormValue("vi_ios_version")

	fmt.Printf("start info  %s \n\n", "start info")
	if id != "" {
		_, err := db.Exec(sqlStatementUpdate, id, screen, name, feature, comment, sectionID, scenario, code, android_version, vi_ios_version, androidApplicable, androidActive, viIosApplicable, viIosActive, organization)
		if err != nil {
			fmt.Printf("err  %s \n\n", err)
			panic(err)
		}
		intVar, errs := strconv.Atoi(id)
		if errs != nil {
			panic(errs)
		}
		return FeatureGet(db, intVar)
	} else {
		sqlStatement := `INSERT INTO "features" ("screen","name","feature","comment", "section_id","scenario","code","android_version",	"vi_ios_version", "android_applicable","android_active","vi_ios_applicable","vi_ios_active","organization") 	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id`

		err := db.QueryRow(sqlStatement, screen, name, feature, comment, sectionID, scenario, code, android_version, vi_ios_version, androidApplicable, androidActive, viIosApplicable, viIosActive, organization).Scan(&lastInsertId)

		//err := db2.Scan(&lastInsertId)
		if err != nil {
			fmt.Printf("e12312312312rr :  %s \n\n", err)
			panic(err)
		}
		return FeatureGet(db, lastInsertId)
	}
}

func FeatureDelete(db *sql.DB, query *http.Request) {
	sqlStatementUpdate := `DELETE FROM "features" WHERE  "id"=$1`
	id := query.FormValue("id")
	intId, errs := strconv.Atoi(id)
	if errs != nil {
		panic(errs)
	}

	fmt.Printf("start info  %s \n\n", "start info")
	if id != "" {
		_, err := db.Exec(sqlStatementUpdate, intId)
		if err != nil {
			fmt.Printf("err  %s \n\n", err)
			panic(err)
		}
	}
}

// save
func FeatureUpdate(db *sql.DB, query *http.Request) Features {
	// Get the values from the request query
	id := query.FormValue("id")
	key := query.FormValue("key")
	value := query.FormValue("value")

	// Convert the id string to integer
	intID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	// Build the SQL statement with the provided key
	sqlStatementUpdate := fmt.Sprintf("UPDATE features SET %s=$1 WHERE id=$2", key)

	// Execute the SQL statement and update the feature
	if id != "" {
		_, err = db.Exec(sqlStatementUpdate, value, intID)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			panic(err)
		}
	}

	// Return the updated feature
	return FeatureGet(db, intID)
}

// func FeaturesList(db *sql.DB) (*services.Response, error) {
func FeaturesList(db *sql.DB) []Features {
	features := []Features{}
	rows, err := db.Query("select features.id, features.name, screen, feature, comment, section_id , android_applicable, android_active, vi_ios_applicable, vi_ios_active, scenario, features.code, android_version, vi_ios_version , organization, sections.name as sections_name, sections.code as sections_code from features left join sections on sections.id = features.section_id")
	if err != nil {
		panic(err)
	}
	// Обработка результатов запроса
	for rows.Next() {
		var featuresSingle Features
		err = rows.Scan(&featuresSingle.ID,
			&featuresSingle.Name,
			&featuresSingle.Screen,
			&featuresSingle.Feature,
			&featuresSingle.Comment,
			&featuresSingle.SectionID,
			&featuresSingle.AndroidApplicable,
			&featuresSingle.AndroidActive,
			&featuresSingle.ViIosApplicable,
			&featuresSingle.ViIosActive,
			&featuresSingle.Scenario,
			&featuresSingle.Code,
			&featuresSingle.AndroidVersion,
			&featuresSingle.ViIosVersion,
			&featuresSingle.Organization,

			&featuresSingle.SectionsCode,
			&featuresSingle.SectionsName,
		)
		if err != nil {
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
	sqlStatement := `select id, name, screen, feature, comment, section_id,  android_applicable, android_active, vi_ios_applicable, vi_ios_active, scenario, code, android_version, vi_ios_version, organization  from features  where id = $1`

	featuresSingle := Features{}

	err := db.QueryRow(sqlStatement, id).Scan(&featuresSingle.ID,
		&featuresSingle.Name,
		&featuresSingle.Screen,
		&featuresSingle.Feature,
		&featuresSingle.Comment,
		&featuresSingle.SectionID,
		&featuresSingle.AndroidApplicable,
		&featuresSingle.AndroidActive,
		&featuresSingle.ViIosApplicable,
		&featuresSingle.ViIosActive,
		&featuresSingle.Scenario,
		&featuresSingle.Code,
		&featuresSingle.AndroidVersion,
		&featuresSingle.ViIosVersion,
		&featuresSingle.Organization,
	)
	if err != nil {
		fmt.Printf("Get :  %s \n\n", err)
		log.Fatal(err)
	}

	return featuresSingle
}
