package extra_query

import (
	"github.com/RhettDelFierro/bettingGolangv2/src/model"
	"net/http"
	"time"
	"net/url"
	"fmt"
	"bytes"
	"encoding/json"
	"database/sql"
)


func GetNBATeams(api_key string, db *sql.DB) (model.NBA_League, error) {
	league_url := "http://api.probasketballapi.com/team"

	client := &http.Client{
		Timeout: time.Second * 100,
	}
	//fmt.Println("here's the api key:", api_key)

	data := url.Values{}
	data.Set("api_key", api_key)
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

	err = insertNBATeams(nba_league, db)
	if err != nil {
		fmt.Println("error in insertNBATeams", err)
		return model.NBA_League{}, err
	}

	return nba_league, nil
}

func insertNBATeams(nba_league model.NBA_League, db *sql.DB) error {
	var err error
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
