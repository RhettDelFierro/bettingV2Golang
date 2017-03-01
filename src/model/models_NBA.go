package model

import "time"

type (
	Teams []Team
	Team struct {
		ID           int `json:"id"`
		DkID         int `json:"dk_id"`
		City         string `json:"city"`
		TeamName     string `json:"team_name"`
		Abbreviation string `json:"abbreviation"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
	}

	Games []Game
	Game []struct {
		ID     int `json:"id"`
		HomeID int `json:"home_id"`
		AwayID int `json:"away_id"`
		Season string `json:"season"`
		Date   time.Time `json:"date,string"`
		Final  int `json:"final"`
	}

	BoxScores []BoxScore
	BoxScore struct {
		GameID     int `json:"game_id"`
		TeamID     int `json:"team_id"`
		OpponentID int `json:"opponent_id"`
		Period     string `json:"period"`
		Season     string `json:"season"`
		Min        string `json:"min"`
		Fgm        int `json:"fgm"`
		Fga        int `json:"fga"`
		Fg3M       int `json:"fg3m"`
		Fg3A       int `json:"fg3a"`
		Ftm        int `json:"ftm"`
		Fta        int `json:"fta"`
		Oreb       int `json:"oreb"`
		Dreb       int `json:"dreb"`
		Ast        int `json:"ast"`
		Blk        int `json:"blk"`
		Stl        int `json:"stl,string"`
		To         int `json:"to"`
		Pf         int `json:"pf"`
		Pts        int `json:"pts"`
		PlusMinus  int `json:"plus_minus"`
	}

	TeamVus []TeamVu
	TeamVu []struct {
		GameID     int `json:"game_id"`
		TeamID     int `json:"team_id"`
		OpponentID int `json:"opponent_id"`
		Period     string `json:"period"`
		Season     string `json:"season"`
		Spd        string `json:"spd"`
		Dist       string `json:"dist"`
		Orbc       int `json:"orbc"`
		Drbc       int `json:"drbc"`
		Tchs       int `json:"tchs"`
		Sast       int `json:"sast"`
		Ftast      int `json:"ftast"`
		Pass       int `json:"pass"`
		Cfgm       int `json:"cfgm"`
		Cfga       int `json:"cfga"`
		Ufgm       int `json:"ufgm"`
		Ufga       int `json:"ufga"`
		Dfgm       int `json:"dfgm"`
		Dfga       int `json:"dfga"`
	}
)