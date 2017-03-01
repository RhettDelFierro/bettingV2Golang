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
const sportsVuUrl string = "http://api.probasketballapi.com/sportsvu/team"  //send season and game_id
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
	fmt.Println("here are boxscores:", boxScores)
	err = dbInsertBoxScores(boxScores)
	if err != nil {
		common.DisplayAppError(w, err, "error in geBoxScoresHTTP", http.StatusInternalServerError)
	}

}

func dbInsertBoxScores(boxScores model.BoxScores) (err error) {

	for _, v := range boxScores {
		err = getTeamVu(v)
		if err != nil {
			return
		}

		_, err = db.Exec("INSERT INTO boxscores (id, team_id, opponent_id, min, fgm, fga, fg3m, fg3a, ftm, fta, oreb, dreb, ast, blk, stl, tno, pf, pts, plus_minus) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)",
			v.GameID, v.TeamID, v.OpponentID, v.Min, v.Fgm, v.Fga, v.Fg3M, v.Fg3A, v.Ftm, v.Fta, v.Oreb, v.Dreb, v.Ast, v.Blk, v.Stl, v.To, v.Pf, v.Pts, v.PlusMinus)
		if err != nil {
			return
		}
	}
	return
}

func getTeamVu(boxScore model.BoxScore) (err error) {
	teamVu, err := getTeamVuHTTP(boxScore)
	if err != nil {
		return
	}

	err = dbInsertTeamVu(teamVu)
	if err != nil {
		return
	}

	return
}

func dbInsertTeamVu(t model.TeamVu) (err error) {
	_, err = db.Exec("INSERT INTO vus (id, team_id, opponent_id, orbc, drbc, sast, ftast, cfgm, cfga, ufgm, ufga, dfgm, dfga) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)",
		t.GameID, t.TeamID, t.OpponentID, t.Orbc, t.Drbc, t.Sast, t.Ftast, t.Cfgm, t.Cfga, t.Ufgm, t.Ufga, t.Dfgm, t.Dfga)
	if err != nil {
		return
	}

	return
}

func getTeamVuHTTP(boxScore model.BoxScore) (teamVu model.TeamVu, err error) {
	var teamVus model.TeamVus
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
		fmt.Println("error in Do getTeamVuHTTP", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	//err = json.NewDecoder(resp.Body).Decode(&games)
	err = json.Unmarshal(body, &teamVus)
	teamVu = teamVus[0]
	return
}

func getBoxScoresHTTP() (boxscores model.BoxScores, err error) {
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
