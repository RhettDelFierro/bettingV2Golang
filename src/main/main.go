package main

import (
	"github.com/gorilla/mux"

	"github.com/RhettDelFierro/bettingGolang/src/models"

	"log"
	"net/http"

)

//func unMarshalRequest(jsonStr []byte) models.GameResults {
//	gameLogs := []models.GameResults{}
//	var data map[string][]models.ResultSets
//
//	err := json.Unmarshal(jsonStr,&data) {
//
//	}
//
//}

func GetPredictions(w http.ResponseWriter, req *http.Request) {

}

func main() {
	models.InitTeams()
	router := mux.NewRouter()
	router.HandleFunc("/predictions", GetPredictions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
