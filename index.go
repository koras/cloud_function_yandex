package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"work/controllers"
	"work/services"

	_ "github.com/lib/pq"
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
	db = services.Connect()

	// по какому урлу стучит клиент

	path := req.URL.Path

	fmt.Printf("path %s \n", path)
	rw.Header().Set("X-Custom-Header", "Test")
	rw.WriteHeader(200)

	if path == "/sections/list" {
		data := controllers.SectorsList(db)
		rw.Header().Set("Content-Type", "application/json")
		jsonSectors, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Отправляем данные в браузер
		rw.Write(jsonSectors)
		return
	}

	if path == "/section/save" {
		//	query := req.URL.Query()

		data := controllers.SectionSave(db, req)
		rw.Header().Set("Content-Type", "application/json")
		jsonSection, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Отправляем данные в браузер
		rw.Write(jsonSection)
		return
	}

	if path == "/features/list" {
		//	name := req.URL.Query().Get("name")
		//	io.WriteString(rw, fmt.Sprintf("Hello, %s!", name))
		data := controllers.FeaturesList(db)
		rw.Header().Set("Content-Type", "application/json")
		jsonFeatures, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Отправляем данные в браузер
		rw.Write(jsonFeatures)
		return
	}

	if path == "/feature/save" {
		query := req.URL.Query()

		//	io.WriteString(rw, fmt.Sprintf("Hello, %s!", name))
		data := controllers.FeatureSave(db, query)
		rw.Header().Set("Content-Type", "application/json")
		jsonFeatures, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		// Отправляем данные в браузер
		rw.Write(jsonFeatures)
		return
	}
	rw.WriteHeader(400)
	io.WriteString(rw, "not found")
}
