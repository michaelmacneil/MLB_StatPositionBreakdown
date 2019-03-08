package main

func getStatisticsByPositionForPlayer(playerId string, readFromFile bool, statsDirectoryPath string, mysqlConnection string) (positionStats []PositionStats) {
  if playerId != "" {

    if(readFromFile) {
      positionStats = getPlayerPositionBaseStatisticsFromFile(playerId, statsDirectoryPath)
    } else {
      positionStats = getPlayerPositionBaseStatisticsFromDB(playerId, mysqlConnection)
    }
  }

  // Sum all position stats and add
  positionStats = append(positionStats, sumStats(positionStats))

  // Calculate percentage statistics from base numbers
  for num := range positionStats {
    positionStats[num] = calculateExtraStats(positionStats[num])
  }

  return positionStats
}

func getMLBTeams(readFromFile bool, directoryPath string, mysqlConnection string) []Team {
  if(readFromFile) {
    return getMLBTeamsFromFile(directoryPath)
  } else {
    return getMLBTeamsFromDB(mysqlConnection)
  }
}

func getAllPlayers(readFromFile bool, directoryPath string, mysqlConnection string) []Player {
  if readFromFile {
    return getAllPlayersFromFile(directoryPath)
  } else {
    return getAllPlayersFromDB(mysqlConnection)
  }
}

func getPlayersFromTeam(teamAbbrev string, readFromFile bool, directoryPath string, mysqlConnection string) []Player {
  if readFromFile {
    return getPlayersFromTeamFromFile(teamAbbrev, directoryPath)
  } else {
    return getPlayersFromTeamFromDB(teamAbbrev, mysqlConnection)
  }
}

func getPlayerIdByPlayerName(playerName string, readFromFile bool, directoryPath string, mysqlConnection string) (playerId string) {
  if readFromFile {
    return getPlayerIdByPlayerNameFromFile(playerName, directoryPath)
  } else {
    return getPlayerIdByPlayerNameFromDB(playerName, mysqlConnection)
  }
}
