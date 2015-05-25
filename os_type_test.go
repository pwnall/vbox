package vbox

import (
  "testing"
)

func TestGetGuestOsTypes(t *testing.T) {
  types, err := GetGuestOsTypes()
  if err != nil {
    t.Fatal(err)
  }
  if len(types) == 0 {
    t.Error("GetTypes returned empty array")
  }

  hasLinux := false
  hasWindows764 := false
  for _, osType := range types {
    id, err := osType.GetId()
    if err != nil {
      t.Error(err)
    }
    switch id {
    case "Linux":
      hasLinux = true
    case "Windows7_64":
      hasWindows764 = true
    }
    if err := osType.Release(); err != nil {
      t.Error(err)
    }
  }

  if hasLinux == false {
    t.Error("Linux type not found")
  }
  if hasWindows764 == false {
    t.Error("Windows7_64 type not found")
  }
}
