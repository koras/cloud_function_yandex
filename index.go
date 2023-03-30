package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"work/controllers"
	"work/services"

	_ "github.com/lib/pq"
)

const (
	SectionsListURL = "/sections/list"
	SectionSaveURL  = "/section/save"
	FeaturesListURL = "/features/list"
	FeatureSaveURL  = "/feature/save"
	FeatureGetURL   = "/feature/get"
)

var db *sql.DB

// вызывается локально для разработки
func main() {

	fmt.Printf("start localhost:8099")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8099", nil)

}
func handler(w http.ResponseWriter, req *http.Request) {
	Handler(w, req)
}

func Start(w http.ResponseWriter, req *http.Request) {
	Handler(w, req)
}

// вызывается в облаке
func Handler(rw http.ResponseWriter, req *http.Request) {
	db, errDb := services.Connect()

	if errDb != nil {
		http.Error(rw, errDb.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	// Извлечение параметров запроса
	path := req.URL.Path
	query := req.URL.Query()

	// по какому урлу стучит клиент
	switch path {
	case SectionsListURL:
		data := controllers.SectorsList(db)
		jsonResponse(rw, data)
	case SectionSaveURL:
		data := controllers.SectionSave(db, req)
		jsonResponse(rw, data)
	case FeaturesListURL:
		data := controllers.FeaturesList(db)
		jsonResponse(rw, data)
	case FeatureSaveURL:
		data := controllers.FeatureSave(db, req)
		jsonResponse(rw, data)
	case FeatureGetURL:
		id := getIntParam(query, "id")
		data := controllers.FeatureGet(db, id)
		jsonResponse(rw, data)
	default:
		rw.WriteHeader(http.StatusNotFound)
		io.WriteString(rw, "not found")
	}
}

func jsonResponse(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(data); err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getIntParam(query url.Values, key string) int {
	if val, err := strconv.Atoi(query.Get(key)); err == nil {
		return val
	}
	return 0
}
