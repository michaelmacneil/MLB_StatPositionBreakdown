package main

import(
  "strings"
  "log"
  "strconv"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func cleanString(object string) string {
    object = strings.Replace(object, " ", "", -1)
    object = strings.Replace(object, "'", "", -1)
    object = strings.Replace(object, "%20%", "", -1)
    object = strings.Replace(object, "\"", "", -1)
    object = strings.Replace(object, "\n", "", -1)
    object = strings.Replace(object, "\r", "", -1)
    return object
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
