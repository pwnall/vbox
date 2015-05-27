package vbox

import (
  "testing"
  "time"
)

func TestConsole_PowerDown(t *testing.T) {
  session := Session{}
  if err := session.Init(); err != nil {
    t.Fatal(err)
  }
  defer session.Release()

  machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
  if err != nil {
    t.Fatal(err)
  }
  defer machine.Release()

  if err := machine.Register(); err != nil {
    t.Fatal(err)
  }
  defer func() {
    media, err := machine.Unregister(CleanupMode_DetachAllReturnHardDisksOnly)
    if err != nil {
      t.Error(err)
    }
    if progress, err := machine.DeleteConfig(media); err != nil {
      t.Error(err)
    } else {
      if err = progress.WaitForCompletion(-1); err != nil {
        t.Error(err)
      }
    }
  }()

  progress, err := machine.Launch(session, "gui", "");
  if err != nil {
    t.Fatal(err)
  }
  defer func() {
    if err = session.UnlockMachine(); err != nil {
      t.Error(err)
      return
    }
    for {
      state, err := session.GetState()
      if err != nil {
        t.Error(err)
        return
      }
      t.Log("Session state: ", state)
      if state == SessionState_Unlocked {
        break
      }
    }

    // TODO(pwnall): Figure out how to get rid of this timeout. The VM should
    //     be unlocked, according to the check above, but unregistering the VM
    //     fails if we don't wait.
    time.Sleep(300 * time.Millisecond)
  }()

  if err = progress.WaitForCompletion(50000); err != nil {
    t.Fatal(err)
  }
  percent, err := progress.GetPercent()
  if err != nil {
    t.Error(err)
  } else if percent != 100 {
    t.Error("VM launch died at percentage: ", percent)
  }
  code, err := progress.GetResultCode()
  if err != nil {
    t.Error(err)
  } else if code != 0 {
    t.Error("VM launch failed with error code: ", code)
  }

  console, err := session.GetConsole()
  if err != nil {
    t.Fatal(err)
  }
  defer console.Release()

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

  progress, err = console.PowerDown()
  if err != nil {
    t.Error(err)
  }

  if err = progress.WaitForCompletion(50000); err != nil {
    t.Fatal(err)
  }
  percent, err = progress.GetPercent()
  if err != nil {
    t.Error(err)
  } else if percent != 100 {
    t.Error("VM power down died at percentage: ", percent)
  }
  code, err = progress.GetResultCode()
  if err != nil {
    t.Error(err)
  } else if code != 0 {
    t.Error("VM power down failed with error code: ", code)
  }
}
