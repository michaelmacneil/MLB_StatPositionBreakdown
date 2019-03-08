package main

import(
        "database/sql"
        _ "github.com/go-sql-driver/mysql"
        "log"
        "strconv"
)

func getPlayerPositionBaseStatisticsFromDB(playerId string, mysqlConnection string) (positionStats []PositionStats) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  log.Println("Getting position stats for playerId: " + playerId)

  // Get all team info
  selectResponse, err := db.Query("SELECT * FROM player_statistics WHERE PlayerId = '" + playerId + "';")
  if err != nil {
    log.Println(err.Error())
  }

  var tempPlayerId string

  // Loop each position and add info to response
  for selectResponse.Next() {
    var currentPosition PositionStats
    err = selectResponse.Scan(
      &tempPlayerId,
      &currentPosition.Position,
      &currentPosition.PositionName,
      &currentPosition.GroundDoublePlay,
      &currentPosition.SacHits,
      &currentPosition.SacFlys,
      &currentPosition.IntentionalWalks,
      &currentPosition.HitByPitch,
      &currentPosition.Strikeouts,
      &currentPosition.HomeRuns,
      &currentPosition.PlateAppearances,
      &currentPosition.Singles,
      &currentPosition.Doubles,
      &currentPosition.Triples)
    if err != nil {
      panic(err.Error())
    } else {
      positionStats = append(positionStats, currentPosition)
    }
  }
  return positionStats
}

func getPlayersFromTeamFromDB(teamAbbrev string, mysqlConnection string) (players []Player) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Get all team info
  selectResponse, err := db.Query("SELECT * FROM players WHERE team_abbrev = '" + teamAbbrev + "';")
  if err != nil {
    log.Println(err.Error())
  }

  // Go through each row return from select and add player to response
  for selectResponse.Next() {
    var player Player
    err = selectResponse.Scan(&player.PlayerId, &player.PlayerName, &player.TeamAbbrev)
    if err != nil {
      panic(err.Error())
    } else {
      players = append(players, player)
    }
  }
  return players
}

func getAllPlayersFromDB(mysqlConnection string) (players []Player) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Get all team info
  selectResponse, err := db.Query("SELECT * FROM players;")
  if err != nil {
    log.Println(err.Error())
  }

  // Go through each row return from select and add player to response
  for selectResponse.Next() {
    var player Player
    err = selectResponse.Scan(&player.PlayerId, &player.PlayerName, &player.TeamAbbrev, &player.ListedPosition)
    if err != nil {
      panic(err.Error())
    } else {
      players = append(players, player)
    }
  }
  return players
}

func addMLBPlayerStatisticsToDBFromFile(playerMap map[string]PlayerContainer, mysqlConnection string) {
  // Initialize database connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Drop table
  drop, err := db.Query("DROP TABLE player_statistics")
  if err != nil {
    log.Println(err.Error())
  }
  defer drop.Close()

  // Create table
  create, err := db.Query("CREATE TABLE player_statistics (PlayerId VARCHAR(20), Position VARCHAR(40), PositionName VARCHAR(40), GroundDoublePlay INT, SacHits INT, SacFlys INT, IntentionalWalks INT, HitByPitch INT, Strikeouts INT, HomeRuns INT, PlateAppearances INT, Singles INT, Doubles INT, Triples INT)")
  if err != nil {
    log.Println(err.Error())
  }
  defer create.Close()

  // Loop each player in the map
  for playerId, player := range playerMap {
    // Loop each position for the current player
    for _, currentPosition := range player.positionStats {
      insert, err := db.Query("INSERT INTO player_statistics VALUES ('" + playerId + "','" + currentPosition.Position + "','" + currentPosition.PositionName + "','" + strconv.Itoa(currentPosition.GroundDoublePlay) + "','" + strconv.Itoa(currentPosition.SacHits) + "','" + strconv.Itoa(currentPosition.SacFlys) + "','" + strconv.Itoa(currentPosition.IntentionalWalks) + "','" + strconv.Itoa(currentPosition.HitByPitch) + "','" + strconv.Itoa(currentPosition.Strikeouts) + "','" + strconv.Itoa(currentPosition.HomeRuns) + "','" + strconv.Itoa(currentPosition.PlateAppearances) + "','" + strconv.Itoa(currentPosition.Singles) + "','" + strconv.Itoa(currentPosition.Doubles) + "','" + strconv.Itoa(currentPosition.Triples) + "')")
      if err != nil {
        log.Println(err.Error())
      }
      insert.Close()
    }
  }
}

func getMLBTeamsFromDB(mysqlConnection string) (teams []Team) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Get all team info
  selectResponse, err := db.Query("SELECT * FROM teams;")
	if err != nil {
		log.Println(err.Error())
	}

  // Loop each team and add info to response
  for selectResponse.Next() {
    var team Team
    err = selectResponse.Scan(&team.Abbrev, &team.CityName, &team.TeamName)
    if err != nil {
      panic(err.Error())
    } else {
      teams = append(teams, team)
    }
  }
  return teams
}

func addMLBTeamsToDB(teams []Team, mysqlConnection string) {
  // Initialize database connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Drop table
  drop, err := db.Query("DROP TABLE teams")
  if err != nil {
    log.Println(err.Error())
  }
  defer drop.Close()

  // Create table
  create, err := db.Query("CREATE TABLE teams (abbrev VARCHAR(20), city_name VARCHAR(40), team_name VARCHAR(40))")
  if err != nil {
    log.Println(err.Error())
  }
  defer create.Close()

  // Loop over all teams
  for _, team := range teams {
    // Add single team to table
    insert, err := db.Query("INSERT INTO teams VALUES ('" + team.Abbrev + "','" + team.CityName + "','" + team.TeamName + "')")
  	if err != nil {
  		log.Println(err.Error())
  	}
    insert.Close()
  }
}

func addPlayersToDB(players []Player, mysqlConnection string) {
  // Initialize database connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Drop table
  drop, err := db.Query("DROP TABLE players")
  if err != nil {
    log.Println(err.Error())
  }
  defer drop.Close()

  // Create table
  create, err := db.Query("CREATE TABLE players (playerId VARCHAR(20), player_name VARCHAR(40), team_abbrev VARCHAR(40))")
  if err != nil {
    log.Println(err.Error())
  }
  defer create.Close()

  // Loop over all players
  for _, player := range players {
    // Add single player to table
    insert, err := db.Query("INSERT INTO players VALUES ('" + player.PlayerId + "','" + player.PlayerName + "','" + player.TeamAbbrev + "')")
    if err != nil {
      log.Println(err.Error())
    }
    insert.Close()
  }
}

func getPlayerIdByPlayerNameFromDB(playerName string, mysqlConnection string) (playerId string) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Get all team info
  selectResponse, err := db.Query("SELECT playerId FROM players WHERE player_name = '" + playerName + "';")
  if err != nil {
    log.Println(err.Error())
  }

  // Go through each row return from select and add player to response
  for selectResponse.Next() {
    err = selectResponse.Scan(&playerId)
    if err != nil {
      panic(err.Error())
    }
  }
  return playerId
}

func removePositionStatsWithZeroPlateAppearances(mysqlConnection string) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Delete position statistics that don't have any plate PlateAppearances
  _, err := db.Query("DELETE FROM player_statistics WHERE PlateAppearances = 0;")
  if err != nil {
    log.Println(err.Error())
  }
}

func removePlayersWithoutPositionStats(mysqlConnection string) {
  // Establish mysql connection
  db, openErr := sql.Open("mysql", mysqlConnection)
  if openErr != nil {
    log.Println(openErr.Error())
  }
  defer db.Close()

  // Delete all entries in players table that don't have a corresponding record in the statistics table
  _, err := db.Query("DELETE FROM players WHERE playerId NOT IN (SELECT stats.PlayerId FROM player_statistics stats);")
  if err != nil {
    log.Println(err.Error())
  }
}
