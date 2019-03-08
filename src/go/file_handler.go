package main

import (
  "io/ioutil"
  "strings"
)

// Input a player's name, this method will scan through the retrosheet storage for a player with that name and return their playerId
func getPlayerIdByPlayerNameFromFile(playerName string, directoryPath string) (playerId string) {
  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  for _, fileInfo := range files {
    //Read file and store character buffer in fileContent
    fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
    check(err)

    var current string
    var list []string

    // Loop over character buffer of file content
    for i:=0 ; i < len(fileContent) ; i++ {
      // Process a single line from the file
      if string(fileContent[i]) != "\n" {
        if string(fileContent[i]) != "\r" {
          current = current + string(fileContent[i])
        }
      } else {
        // This is where we have a complete line - split on commas here
        list = strings.Split(current, ",")

        if len(list) > 2 {
          // See if current player is the player we're searching for
          if strings.Compare(cleanString(playerName),(list[2]+list[1])) == 0 {
            // Return the player's id
            return list[0]
          }
        }
        current = ""
      }
    }
  }
  return ""
}

// Process all files in stats directory to get all position stats for all players
func getPlayerPositionBaseStatisticsFromFileForAllPlayers(directoryPath string) (playerMap map[string]PlayerContainer) {
  start := "start"
  sub := "sub"
  play := "play"

  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  playerMap = make(map[string]PlayerContainer)

  // Loop over all data files in provided directory
  for _, fileInfo := range files {
    //Read file and store character buffer in fileContent
    fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
    check(err)

    var current string
    var list []string

    // Loop over character buffer of file content
    for i:=0 ; i < len(fileContent) ; i++ {
      // Process a single line from the file
      if string(fileContent[i]) != "\n" {
        if string(fileContent[i]) != "\r" {
          current = current + string(fileContent[i])
        }
      } else {
        // This is where we have a complete line - split on commas here
        list = strings.Split(current, ",")
        var playerId string
        var playerContainer PlayerContainer

        // Player start event - ensure player is in map, if not create new player
        if list[0] == start {
          playerId = list[1]
          if _, exists := playerMap[playerId]; !exists {
            playerMap[playerId] = getEmptyPlayerContainer()
          }
          playerContainer = playerMap[playerId]
          playerContainer.currentPosition = list[5]
          playerContainer.currentPositionKey = checkIfPositionIncluded(playerContainer.positionStats, playerContainer.currentPosition)
          if playerContainer.currentPositionKey == -1 {
            // -1 is returned when the value isn't included in the array
            // create a new empty set of stats and set it as the current pos
            playerContainer.positionStats = append(playerContainer.positionStats, getEmptyPositionStats(playerContainer.currentPosition))
            playerContainer.currentPositionKey = checkIfPositionIncluded(playerContainer.positionStats, playerContainer.currentPosition)
          }
          playerMap[playerId] = playerContainer
        }

        if list[0] == sub {
          playerId = list[1]
          if _, exists := playerMap[playerId]; !exists {
            playerMap[playerId] = getEmptyPlayerContainer()
          }
          playerContainer = playerMap[playerId]
          playerContainer.currentPosition = sub
          playerContainer.currentPositionKey = checkIfPositionIncluded(playerContainer.positionStats, sub)
          if playerContainer.currentPositionKey == -1 {
            playerContainer.positionStats = append(playerContainer.positionStats, getEmptyPositionStats(sub))
            playerContainer.currentPositionKey = checkIfPositionIncluded(playerContainer.positionStats, sub)
          }
          playerMap[playerId] = playerContainer
        }

        // Play event - process for player's current position
        if list[0] == play {
          playerId = list[3]
          playerContainer = playerMap[playerId]
          if playerContainer.currentPositionKey != -1 {
            currentPosition := playerContainer.positionStats[playerContainer.currentPositionKey]
            currentPosition.PlateAppearances += 1

            if strings.Contains(list[6], "NP") || strings.Contains(list[6], "CS") || strings.Contains(list[6], "SB") || strings.Contains(list[6], "WP") {
              // Remove no-pitches, runners caught stealing, stolen bases, and wild pitches
              currentPosition.PlateAppearances -= 1
            } else if strings.Contains(list[6], "IW") {
              currentPosition.IntentionalWalks += 1
            } else if strings.Contains(list[6], "W") {
              currentPosition.Walks += 1
            } else if strings.Contains(list[6], "HP") {
              currentPosition.HitByPitch += 1
            } else if strings.Contains(list[6], "K") {
              currentPosition.Strikeouts += 1
            } else if strings.Contains(list[6], "SF") {
              currentPosition.SacFlys += 1
            } else if strings.Contains(list[6], "SH") {
              currentPosition.SacHits += 1
            } else if strings.Contains(list[6], "HR") {
              currentPosition.HomeRuns += 1
            } else if strings.Contains(list[6], "GDP") {
              currentPosition.GroundDoublePlay += 1
            } else if list[6][0] == 'S' {
              currentPosition.Singles += 1
            } else if list[6][0] == 'D' {
              currentPosition.Doubles += 1
            } else if list[6][0] == 'T' {
              currentPosition.Triples += 1
            }
            playerContainer.positionStats[playerContainer.currentPositionKey] = currentPosition
            playerMap[playerId] = playerContainer
          }
        }

        // Reset current string - line has already been processed
        current = ""
      }
    }
  }
  return playerMap
}

// Process all files in stats directory to get all position stats for an individual players
func getPlayerPositionBaseStatisticsFromFile(playerId string, directoryPath string) (positionStats []PositionStats) {
  start := "start"
  sub := "sub"
  play := "play"

  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  currentPositionKey := -1

  // Loop over all data files in provided directory
  for _, fileInfo := range files {
    //Read file and store character buffer in fileContent
    fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
    check(err)

    var current string
    var list []string
    var currentPos string

    // Loop over character buffer of file content
    for i:=0 ; i < len(fileContent) ; i++ {
      // Process a single line from the file
      if string(fileContent[i]) != "\n" {
        if string(fileContent[i]) != "\r" {
          current = current + string(fileContent[i])
        }
      } else {
        // This is where we have a complete line - split on commas here
        list = strings.Split(current, ",")

        // If player start message
        if list[0] == start && list[1] == playerId {
          currentPos = list[5]
          currentPositionKey = checkIfPositionIncluded(positionStats, currentPos)
          if currentPositionKey == -1 {
            // -1 is returned when the value isn't included in the array
            // create a new empty set of stats and set it as the current pos
            positionStats = append(positionStats, getEmptyPositionStats(currentPos))
            currentPositionKey = checkIfPositionIncluded(positionStats, currentPos)
          }
        }

        // If player is substituted into the game
        if list[0] == sub  && list[1] == playerId {
          currentPos = sub
          currentPositionKey = checkIfPositionIncluded(positionStats, sub)
          if currentPositionKey == -1 {
            positionStats = append(positionStats, getEmptyPositionStats(sub))
            currentPositionKey = checkIfPositionIncluded(positionStats, sub)
          }
        }

        // Update position stats for current position for current player for one play
        if list[0] == play && list[3] == playerId {
          positionStats[currentPositionKey].PlateAppearances += 1

          if strings.Contains(list[6], "NP") || strings.Contains(list[6], "CS") || strings.Contains(list[6], "SB") || strings.Contains(list[6], "WP") {
            // Remove no-pitches, runners caught stealing, stolen bases, and wild pitches
            positionStats[currentPositionKey].PlateAppearances -= 1
          } else if strings.Contains(list[6], "IW") {
            positionStats[currentPositionKey].IntentionalWalks += 1
          } else if strings.Contains(list[6], "W") {
            positionStats[currentPositionKey].Walks += 1
          } else if strings.Contains(list[6], "HP") {
            positionStats[currentPositionKey].HitByPitch += 1
          } else if strings.Contains(list[6], "K") {
            positionStats[currentPositionKey].Strikeouts += 1
          } else if strings.Contains(list[6], "SF") {
            positionStats[currentPositionKey].SacFlys += 1
          } else if strings.Contains(list[6], "SH") {
            positionStats[currentPositionKey].SacHits += 1
          } else if strings.Contains(list[6], "HR") {
            positionStats[currentPositionKey].HomeRuns += 1
          } else if strings.Contains(list[6], "GDP") {
            positionStats[currentPositionKey].GroundDoublePlay += 1
          } else if list[6][0] == 'S' {
            positionStats[currentPositionKey].Singles += 1
          } else if list[6][0] == 'D' {
            positionStats[currentPositionKey].Doubles += 1
          } else if list[6][0] == 'T' {
            positionStats[currentPositionKey].Triples += 1
          }
        }

        // Reset current string - line has already been processed
        current = ""
      }
    }
  }
  return positionStats
}

func getMLBTeamsFromFile(directoryPath string) (teams []Team) {
  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  // Loop over all data files in provided directory
  for _, fileInfo := range files {
    //Read file and store character buffer in fileContent
    fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
    check(err)

    var current string
    var list []string

    // Loop over character buffer of file content
    for i:=0 ; i < len(fileContent) ; i++ {
      // Process a single line from the file
      if string(fileContent[i]) != "\n" {
        if string(fileContent[i]) != "\r" {
          current = current + string(fileContent[i])
        }
      } else {
        list = strings.Split(current, ",")

        var team Team
        team.Abbrev = list[0]
        team.CityName = list[2]
        team.TeamName = list[3]
        teams = append(teams, team)

        current = ""
      }
    }
  }
  return teams
}

func getPlayersFromTeamFromFile(teamAbbrev string, directoryPath string) (players []Player) {
  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  // Loop over all data files in provided directory
  for _, fileInfo := range files {
    if strings.Contains(fileInfo.Name(), teamAbbrev) {
      //Read file and store character buffer in fileContent
      fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
      check(err)

      var current string
      var list []string

      // Loop over character buffer of file content
      for i:=0 ; i < len(fileContent) ; i++ {
        // Process a single line from the file
        if string(fileContent[i]) != "\n" {
          if string(fileContent[i]) != "\r" {
            current = current + string(fileContent[i])
          }
        } else {
          list = strings.Split(current, ",")
          var currentPlayer Player
          currentPlayer.TeamAbbrev = list[5]
          if (currentPlayer.TeamAbbrev == teamAbbrev) {
            currentPlayer.PlayerId = list[0]
            currentPlayer.PlayerName = cleanString(list[2]) + " " + cleanString(list[1])
            currentPlayer.ListedPosition = list[6]
            players = append(players, currentPlayer)
          }
          current = "";
        }
      }
    }
  }
  return players
}

func getAllPlayersFromFile(directoryPath string) (players []Player) {
  files, err := ioutil.ReadDir(directoryPath)
  check(err)

  // Loop over all data files in provided directory
  for _, fileInfo := range files {
    //Read file and store character buffer in fileContent
    fileContent, err := ioutil.ReadFile(directoryPath + fileInfo.Name())
    check(err)

    var current string
    var list []string

    // Loop over character buffer of file content
    for i:=0 ; i < len(fileContent) ; i++ {
      // Process a single line from the file
      if string(fileContent[i]) != "\n" {
        if string(fileContent[i]) != "\r" {
          current = current + string(fileContent[i])
        }
      } else {
        list = strings.Split(current, ",")

        var currentPlayer Player
        currentPlayer.TeamAbbrev = list[5]
        currentPlayer.PlayerId = list[0]
        currentPlayer.PlayerName = cleanString(list[2]) + " " + cleanString(list[1])
        currentPlayer.ListedPosition = list[6]
        players = append(players, currentPlayer)

        current = "";
      }
    }
  }
  return players
}
