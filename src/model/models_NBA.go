package model

type (
	NBA_League []NBA_Team

	NBA_Team struct {
		Team_ID      int `json:"id"`
		City         string `json:"city"`
		Team_Name    string `json:"team_name"`
		Abbreviation string `json:"abbreviation"`
	}

	NBA_GameLog []NBA_Boxscore

	NBA_Boxscore struct {
		Game_ID int `json:"game_id"`
		Team_ID int `json:"team_id"`
		Opponent_ID int `json:"opponent_id"`
		Season string `json:"season"`
		Min string `json:"min"`
		FGM int `json:"fmg"`
		FGA int `json:"fga"`
		FG3M int `json:"fg3m"`
		FG3A int `json:"fg3m"`
		FTM int `json:"ftm"`
		FTA int `json:"fta"`
		Oreb int `json:"oreb"`
		Dreb int `json:"dreb"`
		Ast int `json:"ast"`
		Blk int `json:"blk"`
		Stl string `json:"stl"`
		TO int `json:"to"`
		PF int `json:"pf"`
		PTS int `json:"pts"`
		Plus_Minus int `json:"plus_minus"`
	}

	NBA_Team_SVU []NBA_Game_SVU

	NBA_Game_SVU struct {
		Game_ID int `json:"game_id"`
		Team_ID int `json:"team_id"`
		Opponent_ID int `json:"opponent_id"`
		Season string `json:"season"`
		Speed string
		Obrc int `json:"orbc"` //offensive rebound chances if a player is within 3.5ft
		Drbc int `json:"drbc"`
		Sec_Ast int `json:"sast"`
		FT_Ast int `json:"ast"`
		Contested_FGM int `json:"cfgm"`
		Contested_FGA int `json:"cfga"`
		UnContested_FGM int `json:"ufgm"`
		UnContested_FGA int `json:"ufga"`
		RimDef_OppFGM int `json:"dfgm"` //field goals opponent made when team defended rim
		RimDef_OppFGA int `json:"dfga"`
	}
)
