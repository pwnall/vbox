package vbox

import (
  "testing"
)

func TestSystemProperties_GetMediumFormats(t *testing.T) {
  props, err := GetSystemProperties()
  if err != nil {
    t.Fatal(err)
  }
  defer props.Release()

  formats, err := props.GetMediumFormats()
  if err != nil {
    t.Fatal(err)
  }
  if len(formats) == 0 {
    t.Error("GetMediumFormats returned empty array")
  }

  hasVdi := false
  hasVmdk := false
  for _, format := range formats {
    id, err := format.GetId()
    if err != nil {
      t.Error(err)
    }
    t.Log("Found format: ", id)
    switch id {
    case "VDI":
      hasVdi = true
    case "VMDK":
      hasVmdk = true
    }
    if err := format.Release(); err != nil {
      t.Error(err)
    }

    format.Release()
  }

  if hasVdi == false {
    t.Error("VDI format not found")
  }
  if hasVmdk == false {
    t.Error("VMDK format not found")
  }
}
