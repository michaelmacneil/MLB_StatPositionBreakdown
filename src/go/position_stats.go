package main

import(
  "strings"
)

type PositionStats struct {
  Position, PositionName string
  Hits, AtBats, GroundDoublePlay, SacHits, SacFlys, IntentionalWalks, Walks, HitByPitch, Strikeouts, HomeRuns, PlateAppearances, Singles,  Doubles, Triples int
  BattingAverage, OnBasePercentage float64
}

func calculateExtraStats(stats PositionStats) PositionStats {
  stats.Hits = stats.Singles + stats.Doubles + stats.Triples + stats.HomeRuns
  stats.AtBats = stats.PlateAppearances - stats.Walks - stats.IntentionalWalks - stats.HitByPitch - stats.SacFlys - stats.SacHits
  stats.BattingAverage = float64(stats.Hits) / float64(stats.AtBats)
  stats.OnBasePercentage = float64(stats.Hits + stats.Walks + stats.IntentionalWalks + stats.HitByPitch) / float64(stats.AtBats + stats.Walks + stats.IntentionalWalks + stats.HitByPitch + stats.SacFlys + stats.SacHits)
  return stats
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

func getPositionName(position string) (positionName string) {
  if strings.Contains(position, "10") {
    return "DH"
  } else if strings.Contains(position, "11") {
    return "Pinch Hitter"
  } else if strings.Contains(position, "12") {
    return "Pinch Runner"
  } else if strings.Contains(position, "1") {
    return "Pitcher"
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

func checkIfPositionIncluded(stats []PositionStats, position string) (key int) {
  for i := range stats {
    if stats[i].Position == position {
      return i
    }
  }
  return -1
}
