package vbox

import (
  "testing"
)

func TestSystemProperties_GetMaxGuestRam(t *testing.T) {
  props, err := GetSystemProperties()
  if err != nil {
    t.Fatal(err)
  }
  defer props.Release()

  maxRam, err := props.GetMaxGuestRam()
  if err != nil {
    t.Fatal(err)
  }
  if maxRam < (2 << 10) {
    t.Error("Maximum supported guest RAM less than 1024 MB: ", maxRam)
  }
}

func TestSystemProperties_GetMaxGuestVram(t *testing.T) {
  props, err := GetSystemProperties()
  if err != nil {
    t.Fatal(err)
  }
  defer props.Release()

  maxVram, err := props.GetMaxGuestVram()
  if err != nil {
    t.Fatal(err)
  }
  if maxVram < 8 {
    t.Error("Maximum supported guest VRAM less than 8 MB: ", maxVram)
  }
}

func TestSystemProperties_GetMaxGuestCpuCount(t *testing.T) {
  props, err := GetSystemProperties()
  if err != nil {
    t.Fatal(err)
  }
  defer props.Release()

  maxCpus, err := props.GetMaxGuestCpuCount()
  if err != nil {
    t.Fatal(err)
  }
  if maxCpus < 1 {
    t.Error("Maximum supported guest CPU count than 1: ", maxCpus)
  }
}
