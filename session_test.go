package vbox

import (
	"testing"
)

func TestSession_LockMachine_UnlockMachine(t *testing.T) {
	session := Session{}
	if err := session.Init(); err != nil {
		t.Fatal(err)
	}
	defer session.Release()

	machine, err := CreateMachine("", "pwnall_vbox_test", "Linux", "")
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

	if err = session.LockMachine(machine, LockType_Write); err != nil {
		t.Fatal(err)
	}

	machine2, err := session.GetMachine()
	if err != nil {
		t.Error(err)
	} else {
		name2, err := machine2.GetName()
		if err != nil {
			t.Error(err)
		} else if name2 != "pwnall_vbox_test" {
			t.Error("Locked wrong VM: ", name2)
		}
		machine2.Release()
	}

	stype, err := session.GetType()
	if err != nil {
		t.Error(err)
	} else if stype != SessionType_WriteLock {
		t.Error("Wrong locked session type: ", stype)
	}

	state, err := session.GetState()
	if err != nil {
		t.Error(err)
	} else if state != SessionState_Locked {
		t.Error("Wrong locked session state: ", state)
	}

	if err = session.UnlockMachine(); err != nil {
		t.Fatal(err)
	}
}
