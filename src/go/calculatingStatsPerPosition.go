package main

import (
        "log"
        "io/ioutil"
        "os"
        "strings"
        "strconv"
        "net/http"
        "encoding/json"
)

type PositionStats struct {
  Position, PositionName string
  Hits, AtBats, GroundDoublePlay, SacHits, SacFlys, IntentionalWalks, Walks, HitByPitch, Strikeouts, HomeRuns, PlateAppearances, Singles,  Doubles, Triples int
  BattingAverage, OnBasePercentage float64
}

func getPositionName(position string) (positionName string) {
  if strings.Contains(position, "0") {
    return "Pitcher"
  } else if strings.Contains(position, "1") {
    return "DH"
  } else if strings.Contains(position, "2") {
    return "Catcher"
  } else if strings.Contains(position, "3") {
    return "First Base"
  } else if strings.Contains(position, "4") {
    return "Second Base"
  } else if strings.Contains(position, "5") {
    return "Third Base"
  } else if strings.Contains(position, "6") {
    return "Shortstop"
  } else if strings.Contains(position, "7") {
    return "Left Field"
  } else if strings.Contains(position, "8") {
    return "Center Field"
  } else if strings.Contains(position, "9") {
    return "Right Field"
  } else if strings.Contains(position, "sub") {
    return "Substitute"
  } else if strings.Contains(position, "sum") {
    return "All Positions"
  }
  return "Other"
}

func getEmptyPositionStats(position string) (stats PositionStats) {
    stats.Position = position
    stats.PositionName = getPositionName(position)
    stats.IntentionalWalks = 0
    stats.Walks = 0
    stats.HitByPitch = 0
    stats.Strikeouts = 0
    stats.HomeRuns = 0
    stats.PlateAppearances = 0
    stats.Singles = 0
    stats.Doubles = 0
    stats.Triples = 0
    stats.SacFlys = 0
    stats.SacHits = 0
    stats.GroundDoublePlay = 0
    stats.Hits = 0
    stats.AtBats = 0
    stats.BattingAverage = 0
    stats.OnBasePercentage = 0
    return stats
}

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func checkIfPositionIncluded(stats []PositionStats, position string) (key int) {
  for i := range stats {
    if stats[i].Position == position {
      return i
    }
  }
  return -1
}

func calculateExtraStats(stats PositionStats) PositionStats {
  stats.Hits = stats.Singles + stats.Doubles + stats.Triples + stats.HomeRuns
  stats.AtBats = stats.PlateAppearances - stats.Walks - stats.IntentionalWalks - stats.HitByPitch - stats.SacFlys - stats.SacHits
  stats.BattingAverage = float64(stats.Hits) / float64(stats.AtBats)
  stats.OnBasePercentage = float64(stats.Hits + stats.Walks + stats.IntentionalWalks + stats.HitByPitch) / float64(stats.AtBats + stats.Walks + stats.IntentionalWalks + stats.HitByPitch + stats.SacFlys + stats.SacHits)
  return stats
}

func printStats(stats PositionStats) {
  stats = calculateExtraStats(stats)
  log.Println()
  log.Println("Hitting while in position " + stats.PositionName)
  log.Println("Position " + stats.Position)
  log.Println("Plate appearances = " + strconv.Itoa(stats.PlateAppearances))
  log.Println("Walk Count = " + strconv.Itoa(stats.Walks))
  log.Println("Intentional Walk Count = " + strconv.Itoa(stats.IntentionalWalks))
  log.Println("Hit Pitch Count = " + strconv.Itoa(stats.HitByPitch))
  log.Println("Strikeout Count = " + strconv.Itoa(stats.Strikeouts))
  log.Println("Homerun Count = " + strconv.Itoa(stats.HomeRuns))
  log.Println("Single Count = " + strconv.Itoa(stats.Singles))
  log.Println("Double Count = " + strconv.Itoa(stats.Doubles))
  log.Println("Triple Count = " + strconv.Itoa(stats.Triples))
  log.Println("Sac Hit Count = " + strconv.Itoa(stats.SacHits))
  log.Println("Ground Double Play Count = " + strconv.Itoa(stats.GroundDoublePlay))
  log.Println("Hits = " + strconv.Itoa(stats.Hits))
  log.Println("At Bats = " + strconv.Itoa(stats.AtBats))
  log.Println("Batting Average = " + strconv.FormatFloat(stats.BattingAverage, 'G', 4, 64))
  log.Println("On Base Percentage = " + strconv.FormatFloat(stats.OnBasePercentage, 'G', 4, 64))
}

func sumStats(stats []PositionStats) PositionStats {
  combinedStats := getEmptyPositionStats("sum")
  combinedStats.Position = "sum"
  for num := range stats {
    combinedStats.PlateAppearances += stats[num].PlateAppearances
    combinedStats.Walks += stats[num].Walks
    combinedStats.IntentionalWalks += stats[num].IntentionalWalks
    combinedStats.HitByPitch += stats[num].HitByPitch
    combinedStats.Strikeouts += stats[num].Strikeouts
    combinedStats.HomeRuns += stats[num].HomeRuns
    combinedStats.Singles += stats[num].Singles
    combinedStats.Doubles += stats[num].Doubles
    combinedStats.Triples += stats[num].Triples
    combinedStats.SacFlys += stats[num].SacFlys
    combinedStats.SacHits += stats[num].SacHits
    combinedStats.GroundDoublePlay += stats[num].GroundDoublePlay
  }
  return combinedStats
}

func cleanString(object string) string {
    object = strings.Replace(object, " ", "", -1)
    object = strings.Replace(object, "%20%", "", -1)
    object = strings.Replace(object, "\"", "", -1)
    object = strings.Replace(object, "\n", "", -1)
    return object
}

// Input a player's name, this method will scan through the retrosheet storage for a player with that name and return their playerId
func getPlayerIdByPlayerName(playerName string) (playerId string) {
  directoryPath := "../data/rosters/"

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
        current = current + string(fileContent[i])
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

func getStatisticsByPostionForPlayer(playerId string) (positionStats []PositionStats) {
  if playerId != "" {
    directoryPath := "../data/stats/"
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
          current = current + string(fileContent[i])
        } else {
          // This is where we have a complete line - split on commas here
          list = strings.Split(current, ",")

          // Daubach start - set position
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

          if list[0] == sub  && list[1] == playerId {
            currentPos = sub
            currentPositionKey = checkIfPositionIncluded(positionStats, sub)
            if currentPositionKey == -1 {
              positionStats = append(positionStats, getEmptyPositionStats(sub))
              currentPositionKey = checkIfPositionIncluded(positionStats, sub)
            }
          }

          // Daubach play
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
  }

  positionStats = append(positionStats, sumStats(positionStats))

  for num := range positionStats {
    positionStats[num] = calculateExtraStats(positionStats[num])
  }

  return positionStats
}



type Response struct {
  PlayerStatsByPosition []PositionStats
  Message, PlayerId, PlayerName string
  Success bool
}

// Return JSON based on player data request
func getStatisticsByPostionForPlayerRequest(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusCreated)
  if r.URL.Path != "/favicon.ico" && r.Method != "OPTIONS" {
		r.ParseForm()
    var response Response
    var playerStatsByPosition []PositionStats
    playerId := ""
		for k, v := range r.Form {
			val := strings.Join(v, "")
			if k == "playerName" {
        playerId = getPlayerIdByPlayerName(val)
        playerStatsByPosition = getStatisticsByPostionForPlayer(playerId)
        response.Success = true
        response.Message = "Player stats by position returned"
        response.PlayerStatsByPosition = playerStatsByPosition
        response.PlayerName = val
        response.PlayerId = playerId

        json.NewEncoder(w).Encode(response)
        return
			}
		}
		if playerId == "" {
      response.Success = false
      response.Message = "Player could not be found"
      json.NewEncoder(w).Encode(response)
      return
		}
	} else {
    return
  }
}

func main() {
  // Setup log
  f, err := os.OpenFile("../log/calculating_stats_per_position_log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Fatal(err)
  }
  defer f.Close()
  log.SetOutput(f)
	http.HandleFunc("/getStatisticsByPostionForPlayer", getStatisticsByPostionForPlayerRequest)
	netErr := http.ListenAndServe(":9090", nil) // set listen port
	if netErr != nil {
		log.Fatal("ListenAndServe: ", netErr)
	}
}
