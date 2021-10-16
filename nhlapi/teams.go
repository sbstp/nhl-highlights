package nhlapi

import (
	_ "embed" // teams.json
	"encoding/json"
)

type Team struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Abbrev   string `json:"abbreviation"`
	TeamName string `json:"teamName"`
	Location string `json:"locationName"`
}

//go:embed teams.json
var rawTeamsJSON []byte

var Teams []*Team
var TeamsByID map[int32]*Team

func init() {
	type teamsRoot struct {
		Teams []*Team `json:"teams"`
	}

	root := teamsRoot{
		Teams: nil,
	}

	if err := json.Unmarshal(rawTeamsJSON, &root); err != nil {
		panic(err)
	}

	Teams = root.Teams
	TeamsByID = make(map[int32]*Team)
	for _, team := range Teams {
		TeamsByID[team.ID] = team
	}
}
