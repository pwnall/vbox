package vbox

import (
  "os"
  "path"
  "testing"
)

func TestCreateMachine(t *testing.T) {
  goldenPath, err := ComposeMachineFilename("pwnall_vbox_test", "", "")
  if err != nil {
    t.Fatal(err)
  }

  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
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
  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
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
  defer func() {
    progress, err := machine.DeleteConfig([]Medium{});
    if  err != nil {
      t.Error(err)
      return
    }
    defer progress.Release()
    if err = progress.WaitForCompletion(-1); err != nil {
      t.Error(err)
    }
  }()

  modified, err = machine.GetSettingsModified()
  if err != nil {
    t.Fatal(err)
  } else if modified != false {
    t.Fatal("SaveSettings machine still modified")
  }
}

func TestMachine_Register_Unregister(t *testing.T) {
  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  if err = machine.Register(); err != nil {
    t.Fatal(err)
  }

  configFile, err := machine.GetSettingsFilePath()
  if err != nil {
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

  if _, err := os.Stat(configFile); err != nil {
    t.Error("Could not stat VM config file: ", err)
  }

  progress, err := machine.DeleteConfig(mediaList)
  if err != nil {
    t.Fatal(err)
  }
  defer progress.Release()

  if err = progress.WaitForCompletion(10000); err != nil {
    t.Fatal(err)
  }

  if stat, err := os.Stat(configFile); err == nil {
    t.Error("Could still stat VM config file after DeleteConfig: ", stat)
  }
}

func TestFindMachine(t *testing.T) {
  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  if err = machine.Register(); err != nil {
    t.Fatal(err)
  }

  machine2, err := FindMachine("pwnall_vbox_test")
  if err != nil {
    t.Fatal(err)
  }
  defer machine2.Release()
  name2, err := machine2.GetName()
  if err != nil {
    t.Error(err)
  } else if name2 != "pwnall_vbox_test" {
    t.Error("Found VM with wrong name: ", name2)
  }

  mediaList, err := machine.Unregister(CleanupMode_Full)
  if err != nil {
    t.Fatal(err)
  }
  progress, err := machine.DeleteConfig(mediaList)
  if err != nil {
    t.Fatal(err)
  }
  defer progress.Release()
  if err = progress.WaitForCompletion(10000); err != nil {
    t.Fatal(err)
  }

  machine3, err := FindMachine("pwnall_vbox_test")
  if err == nil {
    t.Error("FindMachine retrieved unregistered VM")
    machine3.Release()
  }
}

func TestMachine_AttachDevice_GetMedium(t *testing.T) {
  cwd, err := os.Getwd()
  if err != nil {
    t.Fatal(err)
  }
  testDir := path.Join(cwd, "test_tmp")

  imageFile := path.Join(testDir, "TinyCore-6.2.iso")
  if _, err := os.Stat(imageFile); err != nil {
    t.Fatal(err)
  }

  medium, err := OpenMedium(imageFile, DeviceType_Dvd, AccessMode_ReadOnly,
      false)
  if err != nil {
    t.Fatal(err)
  }
  defer func() {
    // TODO: Figure out how to make this not error out and cause all following
    //       tests to fail.
    if err := medium.Close(); err != nil {
      t.Error(err)
    }
    medium.Release()
  }()

  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  controller, err := machine.AddStorageController("GoIDE", StorageBus_Ide)
  if err != nil {
    t.Fatal(err)
  }
  defer controller.Release()

  if err = controller.SetType(StorageControllerType_Piix4); err != nil {
    t.Fatal(err)
  }
  controller.Release()

  if err = machine.Register(); err != nil {
    t.Fatal(err)
  }

  session := Session{}
  if err := session.Init(); err != nil {
    t.Fatal(err)
  }
  defer session.Release()

  if err = session.LockMachine(machine, LockType_Write); err != nil {
    t.Fatal(err)
  }

  smachine, err := session.GetMachine()
  if err != nil {
    t.Fatal(err)
  }
  defer smachine.Release()

  err = smachine.AttachDevice("GoIDE", 1, 0, DeviceType_Dvd, medium)
  if err != nil {
    t.Fatal(err)
  }

  if err = smachine.SaveSettings(); err != nil {
    t.Fatal(err)
  }

  medium2, err := smachine.GetMedium("GoIDE", 1, 0)
  if err != nil {
    t.Fatal(err)
  }
  defer medium2.Release()

  location2, err := medium2.GetLocation()
  if err != nil {
    t.Error(err)
  } else if location2 != imageFile {
    t.Error("Incorrect medium location: ", location2, " expected: ", imageFile)
  }

  err = smachine.UnmountMedium("GoIDE", 1, 0, true)
  if err != nil {
    t.Fatal(err)
  }

  err = session.UnlockMachine()
  if err != nil {
    t.Fatal(err)
  }

  media, err := machine.Unregister(CleanupMode_DetachAllReturnHardDisksOnly)
  if err != nil {
    t.Fatal(err)
  }
  progress, err := machine.DeleteConfig(media)
  if err != nil {
    t.Fatal(err)
  }
  defer progress.Release()
  if err = progress.WaitForCompletion(50000); err != nil {
    t.Error(err)
  }
}
