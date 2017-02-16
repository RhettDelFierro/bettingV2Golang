package main

import (
	"github.com/gorilla/mux"

	"github.com/RhettDelFierro/bettingGolangv2/src/model"
	"github.com/RhettDelFierro/bettingGolangv2/src/common"
	"log"
	"net/http"

	//"fmt"
	"time"
	//"net/url"
	"os"
	//"bytes"
	"encoding/json"
	"database/sql"
	"fmt"
	//"io/ioutil"
	//"bytes"
	//"strings"
	"bytes"
	//"github.com/RhettDelFierro/rhett_memory_match/src/data"
	"net/url"
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

	nba_league,err := GetNBATeams(key)

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

func GetNBATeams(api_key string) (model.NBA_League,error) {
	league_url := "http://api.probasketballapi.com/team"

	client := &http.Client{
		Timeout: time.Second * 100,
	}
	//fmt.Println("here's the api key:", api_key)

	data := url.Values{}
	data.Set("api_key", api_key)
	fmt.Println("data being sent", api_key)
	req, err := http.NewRequest("POST", league_url, bytes.NewBufferString(data.Encode()))
	//req, err:= http.NewRequest("POST", league_url, strings.NewReader())
	if err != nil {
		fmt.Println("error in new request")
		return model.NBA_League{}, err
	}
	//req.Close = true

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in Do", err)
		return model.NBA_League{}, err
	}
	//body,_ := ioutil.ReadAll(resp.Body)
	//fmt.Println("here's response: ",string(body))


	defer resp.Body.Close()

	var nba_league model.NBA_League
	if err := json.NewDecoder(resp.Body).Decode(&nba_league); err != nil {
		fmt.Println("error in Decode", err)
		return model.NBA_League{}, err
	}

	err = insertNBATeams(nba_league)
	if err != nil {
		fmt.Println("error in insertNBATeams", err)
		return model.NBA_League{}, err
	}

	return nba_league,nil
}

func insertNBATeams(nba_league model.NBA_League) error {
	for _,val := range nba_league {
		_, err = db.Exec("INSERT INTO NBA_Teams(id, city, team_name, abbreviation) VALUES(?, ?, ?, ?)",
			val.Team_ID, val.City, val.Team_Name, val.Abbreviation)
		if err != nil {
			fmt.Println("db error insertNBATeams", err)
			return err
		}
	}

	return err
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
