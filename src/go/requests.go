package main

import(
        "strings"
        "net/http"
        "encoding/json"
)

type StatisticsResponse struct {
  PlayerStatsByPosition []PositionStats
  Message, PlayerId, PlayerName string
  Success bool
}

type TeamResponse struct {
  Teams []Team
  Success bool
}

type PlayersResponse struct {
  Players []Player
  Success bool
}

func (s *server) getPlayersFromTeamRequest() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  	w.WriteHeader(http.StatusCreated)
    if r.URL.Path != "/favicon.ico" && r.Method != "OPTIONS" {
  		r.ParseForm()
      var response PlayersResponse
      var teamAbbrev string
      response.Success = false
      for k, v := range r.Form {
        val := strings.Join(v, "")
        if k == "teamAbbrev" {
          teamAbbrev = val
        }
      }

      response.Players = getPlayersFromTeam(teamAbbrev, s.configuration.ReadFromFile, s.configuration.RosterPath, s.configuration.MysqlConnection)

      if len(response.Players) > 0 {
        response.Success = true
      } else {
        response.Success = false
      }
      json.NewEncoder(w).Encode(response)
      return
    }
  }
}

func (s *server) getStatisticsByPositionForPlayerRequest() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // Return JSON based on player data request
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  	w.WriteHeader(http.StatusCreated)
    var response StatisticsResponse
    response.Success = false
    if r.URL.Path != "/favicon.ico" && r.Method != "OPTIONS" {
  		r.ParseForm()
      var playerStatsByPosition []PositionStats
      playerId := ""
  		for k, v := range r.Form {
  			val := strings.Join(v, "")
  			if k == "playerName" {
          playerId = getPlayerIdByPlayerName(val, s.configuration.ReadFromFile, s.configuration.RosterPath, s.configuration.MysqlConnection)
          if playerId != "" {
            playerStatsByPosition = getStatisticsByPositionForPlayer(playerId, s.configuration.ReadFromFile, s.configuration.StatsPath, s.configuration.MysqlConnection)
            if len(playerStatsByPosition) > 1 {
              response.Success = true
              response.Message = "Player stats by position returned"
              response.PlayerStatsByPosition = playerStatsByPosition
              response.PlayerName = val
              response.PlayerId = playerId
              json.NewEncoder(w).Encode(response)
              return
            } else {
              response.Message = "Player does not have hitting information"
              response.PlayerName = ""
              response.PlayerId = playerId
              response.PlayerStatsByPosition = []PositionStats{}
              json.NewEncoder(w).Encode(response)
              return
            }
          } else {
            response.Message = "Player does not have hitting information"
            response.PlayerName = ""
            response.PlayerId = playerId
            response.PlayerStatsByPosition = []PositionStats{}
            json.NewEncoder(w).Encode(response)
            return
          }
  			}
  		}
  		if response.Success == false {
        response.Message = "Player could not be found"
        response.PlayerName = ""
        response.PlayerId = playerId
        response.PlayerStatsByPosition = []PositionStats{}
        json.NewEncoder(w).Encode(response)
        return
  		}
  	} else {
      response.Success = false
      response.Message = "non-get method called"
      json.NewEncoder(w).Encode(response)
      return
    }
  }
}

func (s *server) getMLBTeamsRequest() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    // Return JSON based on player data request
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  	w.WriteHeader(http.StatusCreated)
    if r.URL.Path != "/favicon.ico" && r.Method != "OPTIONS" {
  		r.ParseForm()
      var response TeamResponse

      response.Teams = getMLBTeams(s.configuration.ReadFromFile, s.configuration.TeamPath, s.configuration.MysqlConnection)

      // If a team was found send success
      if (len(response.Teams) > 0) {
        response.Success = true
        json.NewEncoder(w).Encode(response)
      } else {
        response.Success = false
        json.NewEncoder(w).Encode(response)
      }
      return
    }
  }
}
