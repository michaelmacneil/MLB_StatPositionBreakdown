package main

import (
  "testing"
)

func Test_cleanString_spaces(t *testing.T) {
  spacesRemoved := cleanString("N O _ S P A C E S")
  if spacesRemoved != "NO_SPACES" {
    t.Errorf("cleanString('N O _ S P A C E S') is incorrect, got: %d, want: %d.", spacesRemoved, "NO_SPACES")
  }
}
