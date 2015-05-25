package vbox

import (
  "testing"
)

func TestGetMachines(t *testing.T) {
  machines, err := GetMachines()
  if err != nil {
    t.Fatal(err)
  }
  if len(machines) == 0 {
    t.Error("GetMachines returned empty array")
  }

  for _, machine := range machines {
    if err := machine.Release(); err != nil {
      t.Error(err)
    }
  }
}

func TestCreateMachine(t *testing.T) {
  goldenPath, err := ComposeMachineFilename("pwnall_vbox_test", "", "")
  if err != nil {
    t.Fatal(err)
  }

  machine, err := CreateMachine("pwnall_vbox_test", "Linux",
      "forceOverwrite=1")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  name, err := machine.GetName()
  if err != nil {
    t.Error(err)
  } else if name != "pwnall_vbox_test" {
    t.Error("Wrong machine name: ", name)
  }

  osTypeId, err := machine.GetOsTypeId()
  if err != nil {
    t.Error(err)
  } else if osTypeId != "Linux" {
    t.Error("Wrong OS type ID: ", osTypeId)
  }

  path, err := machine.GetSettingsFilePath()
  if err != nil {
    t.Error(err)
  } else if path != goldenPath {
    t.Error("Wrong settings path: ", path, " expected: ", goldenPath)
  }

  modified, err := machine.GetSettingsModified()
  if err != nil {
    t.Error(err)
  } else if modified != true {
    t.Error("Wrong modified flag", modified)
  }
}

func TestMachine_SaveSettings(t *testing.T) {
  machine, err := CreateMachine("pwnall_vbox_test", "Linux",
      "forceOverwrite=1")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  modified, err := machine.GetSettingsModified()
  if err != nil {
    t.Fatal(err)
  } else if modified != true {
    t.Fatal("New machine appears not to be modified")
  }

  if err = machine.SaveSettings(); err != nil {
    t.Fatal(err)
  }

  modified, err = machine.GetSettingsModified()
  if err != nil {
    t.Fatal(err)
  } else if modified != false {
    t.Fatal("SaveSettings machine still modified")
  }
}
