package model

type (
	NBA_League []NBA_Team

	NBA_Team struct {
		Team_ID      int `json:"id"`
		City         string `json:"city"`
		Team_Name    string `json:"team_name"`
		Abbreviation string `json:"abbreviation"`
	}
)
