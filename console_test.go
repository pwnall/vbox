package vbox

import (
  "testing"
)

func TestConsole_GetMachine_PowerDown(t *testing.T) {
  WithDvdInVm(t, "", func (machine Machine, session Session, console Console) {
    machine2, err := console.GetMachine()
    if err != nil {
      t.Error(err)
    } else {
      name2, err := machine2.GetName()
      if err != nil {
        t.Error(err)
      } else if name2 != "pwnall_vbox_test" {
        t.Error("Got wrong VM from console: ", name2)
      }
      machine2.Release()
    }

    // NOTE: PowerDown is tested by the WithDvdInVm helper.
  })
}
