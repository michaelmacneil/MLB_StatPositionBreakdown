package main

import(
  "os"
  "encoding/json"
  "log"
)

type Configuration struct {
  UpdateDatabase, ReadFromFile bool
  LogPath, RosterPath, TeamPath, StatsPath, MysqlConnection string
}

func getConfigurationFromFile(configurationFilePath string) (configuration Configuration) {
  // Setup configuration
  configFile, _ := os.Open(configurationFilePath)
  defer configFile.Close()
  configDecoder := json.NewDecoder(configFile)
  configErr := configDecoder.Decode(&configuration)
  if configErr != nil {
    log.Fatal("Error processing configuration file: " + configErr.Error())
  }
  return configuration
}
