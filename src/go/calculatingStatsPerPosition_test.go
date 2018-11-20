package main

import (
  "testing"
)

func Test_getPositionName_ValidString(t *testing.T) {
  positionName := getPositionName("1")
  if positionName != "DH" {
    t.Errorf("getPositionName('1') is incorrect, got: %d, want: %d.", positionName, "DH")
  }
}

func Test_getPositionName_InvalidString(t *testing.T) {
  positionName := getPositionName("N/A")
  if positionName != "Other" {
    t.Errorf("getPositionName('N/A') is incorrect, got: %d, want: %d.", positionName, "Other")
  }
}

func Test_cleanString_spaces(t *testing.T) {
  spacesRemoved := cleanString("N O _ S P A C E S")
  if spacesRemoved != "NO_SPACES" {
    t.Errorf("cleanString('N O _ S P A C E S') is incorrect, got: %d, want: %d.", spacesRemoved, "NO_SPACES")
  }
}

func Test_sumStats_twoPositions(t *testing.T) {
  dhStats := getEmptyPositionStats("DH")
  dhStats.Singles = 2
  firstStats := getEmptyPositionStats("3")
  firstStats.Singles = 1
  positionStats := []PositionStats{dhStats, firstStats}
  sumStats := sumStats(positionStats)
  if sumStats.Singles != 3 {
    t.Errorf("sumStats is incorrect, got: %d, want: %d.", sumStats.Singles, 3)
  }
}
