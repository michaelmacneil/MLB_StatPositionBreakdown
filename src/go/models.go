package main

type Team struct {
  Abbrev string
  CityName string
  TeamName string
}

type Player struct {
  PlayerId string
  PlayerName string
  TeamAbbrev string
  ListedPosition string
}

type PlayerContainer struct {
  currentPositionKey int
  currentPosition string
  positionStats []PositionStats
}

func getEmptyPlayerContainer() (player PlayerContainer) {
  player.currentPositionKey = -1
  player.currentPosition = "not_set"
  return player
}
