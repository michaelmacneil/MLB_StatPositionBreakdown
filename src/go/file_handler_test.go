package main

import (
  "testing"
)

func Test_getPlayerIdByPlayerNameFromFile(t *testing.T) {
  expectedPlayerId := "gardb001"
  playerName := "Brett Gardner"
  playerId := getPlayerIdByPlayerNameFromFile(playerName, "../data/test/rosters/")
  if playerId != expectedPlayerId {
    t.Errorf("getPlayerIdByPlayerNameFromFile is incorrect, playerId should have been %d but was %d", expectedPlayerId, playerId)
  }
}

func Test_getPlayerPositionBaseStatsiticsFromFileForAllPlayers(t *testing.T) {
  expectedPlayerMapSize := 1
  expectedPlayerId := "gardb001"
  expectedPlayerHomeRuns := 1
  expectedPlayerPositionName := "Left Field"
  var playerMap map[string]PlayerContainer
  playerMap = getPlayerPositionBaseStatisticsFromFileForAllPlayers("../data/test/stats/")
  if len(playerMap) != expectedPlayerMapSize {
    t.Errorf("playerMapSize is incorrect, should have gotten %d from file but got %d", expectedPlayerMapSize, len(playerMap))
  }
  for playerId, player := range playerMap {
    if playerId != expectedPlayerId {
      t.Errorf("playerId is incorrect, got: %d, want: %d.", playerId, expectedPlayerId)
    }
    if player.positionStats[0].PositionName != expectedPlayerPositionName {
      t.Errorf("PositionName is incorrect, got: %d, want: %d.", player.positionStats[0].PositionName, expectedPlayerPositionName)
    }
    if player.positionStats[0].HomeRuns != expectedPlayerHomeRuns {
      t.Errorf("HomeRuns is incorrect, got: %d, want: %d.", player.positionStats[0].HomeRuns, expectedPlayerHomeRuns)
    }
  }
}

func Test_getPlayerPositionBaseStatisticsFromFile(t *testing.T) {
  positionStats := getPlayerPositionBaseStatisticsFromFile("gardb001", "../data/test/stats/")
  expectedPlayerHomeRuns := 1
  expectedPlayerPositionName := "Left Field"
  if positionStats[0].PositionName != expectedPlayerPositionName {
    t.Errorf("PositionName is incorrect, got: %d, want %d.", positionStats[0].PositionName, expectedPlayerPositionName)
  }
  if positionStats[0].HomeRuns != expectedPlayerHomeRuns {
    t.Errorf("HomeRuns is incorrect, got: %d, want %d.", positionStats[0].HomeRuns, expectedPlayerHomeRuns)
  }
}

func Test_getMLBTeamsFromFile(t *testing.T) {
  mlbTeams := getMLBTeamsFromFile("../data/test/teams/")
  expectedTeamName := "Peaches"
  expectedCityName := "Rockford"
  expectedAbbrev := "ROC"
  if mlbTeams[0].TeamName != expectedTeamName {
    t.Errorf("TeamName is incorrect, got: %d, want %d.", mlbTeams[0].TeamName, expectedTeamName)
  }
  if mlbTeams[0].CityName != expectedCityName {
    t.Errorf("CityName is incorrect, got: %d, want %d.", mlbTeams[0].CityName, expectedCityName)
  }
  if mlbTeams[0].Abbrev != expectedAbbrev {
    t.Errorf("TeamAbbrev is incorrect, got: %d, want %d.", mlbTeams[0].Abbrev, expectedAbbrev)
  }
}

func Test_getPlayersFromTeamFromFile(t *testing.T) {
  players := getPlayersFromTeamFromFile("FAKE", "../data/test/rosters/")
  expectedPlayerCount := 1
  expectedPlayerId := "gardb001"
  expectedPlayerName := "Brett Gardner"
  if len(players) != expectedPlayerCount {
    t.Errorf("Player Count is incorrect, got: %d, want %d.", len(players), expectedPlayerCount)
  }
  if players[0].PlayerId != expectedPlayerId {
    t.Errorf("PlayerId is incorrect, got: %d, want %d.", players[0].PlayerId, expectedPlayerId)
  }
  if players[0].PlayerName != expectedPlayerName {
    t.Errorf("PlayerName is incorrect, got: %d, want %d.", players[0].PlayerName, expectedPlayerName)
  }
}

func Test_getAllPlayersFromFile(t *testing.T) {
  players := getAllPlayersFromFile("../data/test/rosters/")
  expectedPlayerCount := 1
  expectedPlayerId := "gardb001"
  expectedPlayerName := "Brett Gardner"
  if len(players) != expectedPlayerCount {
    t.Errorf("Player Count is incorrect, got: %d, want %d.", len(players), expectedPlayerCount)
  }
  if players[0].PlayerId != expectedPlayerId {
    t.Errorf("PlayerId is incorrect, got: %d, want %d.", players[0].PlayerId, expectedPlayerId)
  }
  if players[0].PlayerName != expectedPlayerName {
    t.Errorf("PlayerName is incorrect, got: %d, want %d.", players[0].PlayerName, expectedPlayerName)
  }
}
