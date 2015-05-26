package vbox

import (
  "testing"
)


func TestMachine_AddStorageController(t *testing.T) {
  machine, err := CreateMachine("pwnall_vbox_test", "Linux",
      "forceOverwrite=1")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  controller, err := machine.AddStorageController(
      "Controller: IDE", StorageBus_Ide)
  if err != nil {
    t.Error(err)
  }
  defer controller.Release()

  name, err := controller.GetName()
  if err != nil {
    t.Error(err)
  } else if name != "Controller: IDE" {
    t.Error("Wrong controller name: ", name)
  }

  bus, err := controller.GetBus()
  if err != nil {
    t.Error(err)
  } else if bus != StorageBus_Ide {
    t.Error("Wrong controller bus: ", bus)
  }
}
