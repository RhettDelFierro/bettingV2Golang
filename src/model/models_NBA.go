package model

import "github.com/RhettDelFierro/bettingGolang/src/stats"

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
		Date   string `json:"date"`
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
		Stl        string `json:"stl"`
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

//type (
//	NBA_League []NBA_Team
//
//	NBA_Team struct {
//		Team_ID      int `json:"id"`
//		City         string `json:"city"`
//		Team_Name    string `json:"team_name"`
//		Abbreviation string `json:"abbreviation"`
//	}
//
//	NBA_GameLog []NBA_Boxscore
//
//	NBA_Boxscore struct {
//		Game_ID int `json:"game_id"`
//		Team_ID int `json:"team_id"`
//		Opponent_ID int `json:"opponent_id"`
//		Season string `json:"season"`
//		Min string `json:"min"`
//		FGM int `json:"fmg"`
//		FGA int `json:"fga"`
//		FG3M int `json:"fg3m"`
//		FG3A int `json:"fg3m"`
//		FTM int `json:"ftm"`
//		FTA int `json:"fta"`
//		Oreb int `json:"oreb"`
//		Dreb int `json:"dreb"`
//		Ast int `json:"ast"`
//		Blk int `json:"blk"`
//		Stl string `json:"stl"`
//		TO int `json:"to"`
//		PF int `json:"pf"`
//		PTS int `json:"pts"`
//		Plus_Minus int `json:"plus_minus"`
//	}
//
//	NBA_Team_SVU []NBA_Game_SVU
//
//	NBA_Game_SVU struct {
//		Game_ID int `json:"game_id"`
//		Team_ID int `json:"team_id"`
//		Opponent_ID int `json:"opponent_id"`
//		Season string `json:"season"`
//		Speed string
//		Obrc int `json:"orbc"` //offensive rebound chances if a player is within 3.5ft
//		Drbc int `json:"drbc"`
//		Sec_Ast int `json:"sast"`
//		FT_Ast int `json:"ast"`
//		Contested_FGM int `json:"cfgm"`
//		Contested_FGA int `json:"cfga"`
//		UnContested_FGM int `json:"ufgm"`
//		UnContested_FGA int `json:"ufga"`
//		RimDef_OppFGM int `json:"dfgm"` //field goals opponent made when team defended rim
//		RimDef_OppFGA int `json:"dfga"`
//	}
//)
