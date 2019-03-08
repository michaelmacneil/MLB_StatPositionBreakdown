package main

import (
        "log"
        "os"
        "net/http"
)

func main() {
  srv := server{}
  srv.Init()

  // Setup log
  f, err := os.OpenFile(srv.configuration.LogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
    log.Fatal(srv.configuration.LogPath)
  }
  defer f.Close()
  log.SetOutput(f)

  // Only repopulate database if configured to do so
  if (srv.configuration.UpdateDatabase) {
    addMLBTeamsToDB(getMLBTeamsFromFile(srv.configuration.TeamPath), srv.configuration.MysqlConnection)
    addPlayersToDB(getAllPlayersFromFile(srv.configuration.RosterPath), srv.configuration.MysqlConnection)
    addMLBPlayerStatisticsToDBFromFile(getPlayerPositionBaseStatisticsFromFileForAllPlayers(srv.configuration.StatsPath), srv.configuration.MysqlConnection)
    removePositionStatsWithZeroPlateAppearances(srv.configuration.MysqlConnection)
    removePlayersWithoutPositionStats(srv.configuration.MysqlConnection)
    log.Print("Finished mysql update")
  }

  srv.routes()

  // Start web server
	netErr := http.ListenAndServe(":9090", srv.router)
	if netErr != nil {
		log.Fatal("ListenAndServe: ", netErr)
	}
}
