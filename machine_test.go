package vbox

import (
  "testing"
)

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

func TestMachine_Register_Unregister(t *testing.T) {
  machine, err := CreateMachine("pwnall_vbox_test", "Linux",
      "forceOverwrite=1")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  if err = machine.Register(); err != nil {
    t.Fatal(err)
  }

  machineList, err := GetMachines()
  if err != nil {
    t.Fatal(err)
  }
  foundMachine := false
  for _, regMachine := range machineList {
    name, err := regMachine.GetName()
    if err != nil {
      t.Error(err)
    }
    t.Log("After Register, found machine with name: ", name)
    if name == "pwnall_vbox_test" {
      foundMachine = true
    }
    regMachine.Release()
  }
  if foundMachine == false {
    t.Error("Newly registered machine does not show up on GetMachines list")
  }

  mediaList, err := machine.Unregister(CleanupMode_Full)
  if err != nil {
    t.Fatal(err)
  }

  if len(mediaList) != 0 {
    t.Error("Machine un-registration returned unexpected attached media")
  }

  machineList, err = GetMachines()
  if err != nil {
    t.Fatal(err)
  }
  for _, regMachine := range machineList {
    name, err := regMachine.GetName()
    if err != nil {
      t.Error(err)
    }
    t.Log("After Unregister, found machine with name: ", name)
    if name == "pwnall_vbox_test" {
      t.Error("Unregistered machine still shows up on GetMachines list")
      break
    }
    regMachine.Release()
  }
}
