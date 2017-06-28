package vbox

// This package contains common testing code.
//
// Normally, we'd consider exporting the helpers to package users. However,
// this package aims to export the VirtualBox XPCOM API, as it is. Furthermore,
// the helpers are rather ad-hoc, and we don't want to be bound to support
// their APIs in follow-up releases.

import (
	"os"
	"path"
	"testing"
	"time"
)

// WithDvdInVm runs the given func with a launched VM with an attached DVD.
// If isoName is the empty string, no DVD is inserted into the VM's DVD drive.
func WithDvdInVm(t *testing.T, isoName string, disableBootPrompt bool,
	testCase func(Machine, Session, Console)) {
	medium := Medium{}

	if isoName != "" {
		cwd, err := os.Getwd()
		if err != nil {
			t.Fatal(err)
		}
		testDir := path.Join(cwd, "test_tmp")

		imageFile := path.Join(testDir, isoName)
		if _, err := os.Stat(imageFile); err != nil {
			t.Fatal(err)
		}

		medium, err = OpenMedium(imageFile, DeviceType_Dvd, AccessMode_ReadOnly,
			false)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			// TODO: Figure out how to make this not error out and cause all
			//       following tests to fail.
			if err := medium.Close(); err != nil {
				t.Error(err)
			}
			medium.Release()
		}()
	}

	session := Session{}
	if err := session.Init(); err != nil {
		t.Fatal(err)
	}
	defer session.Release()

	machine, err := CreateMachine("pwnall_vbox_test", "Ubuntu", "")
	if err != nil {
		t.Fatal(err)
	}
	defer machine.Release()

	// NOTE: Using the USB 1.1 controller so tests can run on the open-sourced
	//       VirtualBox (no extension packs).
	controller, err := machine.AddUsbController("GoUSB", UsbControllerType_Ohci)
	if err != nil {
		t.Fatal(err)
	}
	defer controller.Release()

	// NOTE: LXDE wants at least 256MB for live CD, recommends 512MB.
	if err := machine.SetMemorySize(512); err != nil {
		t.Error(err)
	}
	if err := machine.SetPointingHidType(PointingHidType_UsbTablet); err != nil {
		t.Error(err)
	}

	if disableBootPrompt {
		settings, err := machine.GetBiosSettings()
		if err != nil {
			t.Fatal(err)
		}
		defer settings.Release()

		if err := settings.SetLogoFadeIn(false); err != nil {
			t.Error(err)
		}
		if err := settings.SetLogoFadeOut(false); err != nil {
			t.Error(err)
		}
		if err := settings.SetBootMenuMode(BootMenuMode_Disabled); err != nil {
			t.Error(err)
		}
	}

	if err := machine.Register(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		media, err := machine.Unregister(CleanupMode_DetachAllReturnHardDisksOnly)
		if err != nil {
			t.Error(err)
		}
		progress, err := machine.DeleteConfig(media)
		if err != nil {
			t.Error(err)
			return
		}
		defer progress.Release()
		if err = progress.WaitForCompletion(-1); err != nil {
			t.Error(err)
		}
	}()

	AddDvdToMachine(t, machine, medium, session)
	defer RemoveDvdFromMachine(t, machine, session)

	progress, err := machine.Launch(session, "gui", "")
	if err != nil {
		t.Fatal(err)
	}
	defer progress.Release()
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
		time.Sleep(1 * time.Second)
	}()

	if err = progress.WaitForCompletion(50000); err != nil {
		t.Fatal(err)
	}

	console, err := session.GetConsole()
	if err != nil {
		t.Fatal(err)
	}
	defer console.Release()

	defer func() {
		progress, err := console.PowerDown()
		if err != nil {
			t.Error(err)
			return
		}
		defer progress.Release()

		if err = progress.WaitForCompletion(50000); err != nil {
			t.Error(err)
		}

		percent, err := progress.GetPercent()
		if err != nil {
			t.Error(err)
		} else if percent != 100 {
			t.Error("VM power down died at percentage: ", percent)
		}
		code, err := progress.GetResultCode()
		if err != nil {
			t.Error(err)
		} else if code != 0 {
			t.Error("VM power down failed with error code: ", code)
		}
	}()

	testCase(machine, session, console)
}

func AddDvdToMachine(t *testing.T, machine Machine, medium Medium,
	session Session) {
	if err := session.LockMachine(machine, LockType_Write); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := session.UnlockMachine(); err != nil {
			t.Error(err)
		}
	}()

	// NOTE: Machine modifications require the mutable instance obtained from
	smachine, err := session.GetMachine()
	if err != nil {
		t.Fatal(err)
	}
	defer smachine.Release()

	controller, err := smachine.AddStorageController("GoIDE", StorageBus_Ide)
	if err != nil {
		t.Fatal(err)
	}
	defer controller.Release()

	if err = controller.SetType(StorageControllerType_Piix4); err != nil {
		t.Fatal(err)
	}

	err = smachine.AttachDevice("GoIDE", 1, 0, DeviceType_Dvd, medium)
	if err != nil {
		t.Fatal(err)
	}

	if err = smachine.SaveSettings(); err != nil {
		t.Fatal(err)
	}
}

func RemoveDvdFromMachine(t *testing.T, machine Machine, session Session) {
	if err := session.LockMachine(machine, LockType_Write); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := session.UnlockMachine(); err != nil {
			t.Error(err)
		}
	}()

	// NOTE: Machine modifications require the mutable instance obtained from
	smachine, err := session.GetMachine()
	if err != nil {
		t.Fatal(err)
	}
	defer smachine.Release()

	err = smachine.DetachDevice("GoIDE", 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	if err = smachine.SaveSettings(); err != nil {
		t.Fatal(err)
	}
}
