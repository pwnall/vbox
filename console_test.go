package vbox

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestConsole_GetMachine_PowerDown(t *testing.T) {
	WithDvdInVm(t, "", false, /* disableBootMenu */
		func(machine Machine, session Session, console Console) {
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

func TestConsole_GetEventSource(t *testing.T) {
	var listeningEvents atomic.Value
	listeningEvents.Store(true)

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

	// NOTE: LXDE wants at least 256MB for live CD, recommends 512MB.
	if err := machine.SetMemorySize(512); err != nil {
		t.Error(err)
	}

	if err := machine.Register(); err != nil {
		t.Fatal(err)
	}
	defer func() {
		media, err := machine.Unregister(CleanupMode_Full)
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

	eventSource, err := console.GetEventSource()
	if err != nil {
		t.Fatal(err)
	}
	defer eventSource.Release()

	listener, err := eventSource.CreateListener()
	if err != nil {
		t.Fatal(err)
	}
	defer listener.Release()

	interestingEvents := []uint32{EventType_OnStateChanged}
	if err := eventSource.RegisterListener(listener, interestingEvents, false); err != nil {
		t.Fatal(err)
	}
	defer func() {
		err := eventSource.UnregisterListener(listener)
		if err != nil {
			t.Fatal(err)
		}
	}()

	listeningEvents.Store(true)
	go func() {
		for listeningEvents.Load() == true {
			t.Log("Waiting for event")
			event, err := eventSource.GetEvent(listener, 250)
			if err != nil {
				t.Fatal(err)
			}

			if event == nil {
				continue
			}

			eventType, err := event.GetType()
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("Received event %d", eventType)
			if eventType != EventType_OnStateChanged {
				t.Errorf("Wrong type for event. Expected %d, got %d", EventType_OnStateChanged, eventType)
			}

			err = eventSource.EventProcessed(listener, *event)
			if err != nil {
				t.Fatal(err)
			}

			event.Release()
		}
	}()

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
		t.Log("VM is powered down")
	}()

	time.Sleep(3 * time.Second)
	listeningEvents.Store(false)
}
