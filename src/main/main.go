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
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"net/url"
	"bytes"
	"io/ioutil"
	"strconv"
)

var db *sql.DB
var api_key string

const gamesUrl string = "http://api.probasketballapi.com/game"              //send season: 2016
const teamsUrl string = "http://api.probasketballapi.com/team"              //send only api key
const sportsVuUrl string = "http://api.probasketballapi.com/sportsvu/team"     //send season and game_id
const boxscoresUrl string = "http://api.probasketballapi.com/boxscore/team" //send season

func init() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("NBA_DB"))
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")

	api_key = os.Getenv("NBA_API_KEY")
}

func GetData(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//check if the DB has been populated:

	//check for games

	//err := getGames()
	//common.DisplayAppError(w, err, "error in getRecentGamesHTTP", http.StatusInternalServerError)

	boxScores, err := getBoxScoresHTTP()
	if err != nil {
		common.DisplayAppError(w, err, "error in geBoxScoresHTTP", http.StatusInternalServerError)
	}

	err := dbInsertBoxScores(boxScores)
}

func dbInsertBoxScores(boxScores model.BoxScores) (err error) {
	for _,v := range boxScores {
		err = getTeamVus(v)
		if err != nil {
			return
		}
	}

	return
}

func getTeamVus(boxScore model.BoxScore) (err error) {
	teamVUs, err := getTeamVusHTTP(boxScore)
}

func getTeamVusHTTP(boxScore model.BoxScore) (teamVus model.TeamVus, err error){
	client := &http.Client{
		Timeout: time.Second * 100,
	}

	data := url.Values{}
	data.Set("api_key", api_key)
	data.Set("season", "2016")
	data.Set("team_id", strconv.Itoa(boxScore.TeamID))
	data.Set("game_id", strconv.Itoa(boxScore.GameID))
	req, err := http.NewRequest("POST", sportsVuUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("error in new request")
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in Do", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&games)
	err = json.Unmarshal(body, &boxScores)
	return
}

func getBoxScoresHTTP() (boxscores model.BoxScores, err error){
	client := &http.Client{
		Timeout: time.Second * 100,
	}

	data := url.Values{}
	data.Set("api_key", api_key)
	data.Set("season", "2016")
	req, err := http.NewRequest("POST", boxscoresUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("error in new request")
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in Do", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&games)
	err = json.Unmarshal(body, &boxscores)
	return
}

func getGames() error {
	games, err := getRecentGamesHTTP()
	if err != nil {
		return err
	}

	err = dbInsertGames(games)
	if err != nil {
		return err
	}

	return nil
}

func getRecentGames() (rows *sql.Rows, err error) {
	//get all the games
	//remember, want to query for recent game.
	rows, err = db.Query("SELECT * FROM games")
	defer rows.Close()
	return
}

func getRecentGamesHTTP() (games model.Games, err error) {
	client := &http.Client{
		Timeout: time.Second * 100,
	}

	data := url.Values{}
	data.Set("api_key", api_key)
	data.Set("season", "2016")
	req, err := http.NewRequest("POST", gamesUrl, bytes.NewBufferString(data.Encode()))
	if err != nil {
		fmt.Println("error in new request")
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in Do", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&games)
	err = json.Unmarshal(body, &games)
	return
}

//should make a query to check if the game exists.
func dbInsertGames(games model.Games) (err error) {
	for _, v := range games {
		_, err = db.Exec("INSERT INTO games (id, home_id, away_id, season, game_date, final) VALUES ($1, $2, $3, $4, $5, $6)",
			v.ID, v.HomeID, v.AwayID, v.Season, v.Date.Time, v.Final)
		if err != nil {
			return
		}
	}
	return
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/predictions", GetData).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
