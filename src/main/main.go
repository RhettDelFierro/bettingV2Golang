package main

import (
	"github.com/gorilla/mux"
	"github.com/RhettDelFierro/bettingGolangv2/src/model"
	"github.com/RhettDelFierro/bettingGolangv2/src/common"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"database/sql"
)

var db *sql.DB
var err error

type EveryTeam struct {
	Results model.NBA_League `json:"results"`
}


func GetPredictions(w http.ResponseWriter, req *http.Request) {
	//ch := make(chan []byte)
	key := os.Getenv("NBA_API_KEY")
	//fmt.Println("here's the api key:", key)

	//nba_league,err := extra_query.GetNBATeams(key,db)

	if err != nil {
		common.DisplayAppError(w, err, "Error parsing team JSON.", 500)
	}

	if j, err := json.Marshal(EveryTeam{Results: nba_league}); err != nil {
		common.DisplayAppError(w, err, "error in GetPredictions json.Marshal", 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}


}

func main() {
	db, err = common.CreateDbSession()
	if err != nil {
		log.Panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/predictions", GetPredictions).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
