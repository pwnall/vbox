package vbox

import (
	"testing"
)

func TestMachine_AddUsbController(t *testing.T) {
	machine, err := CreateMachine("pwnall_vbox_test", "Linux", "")
	if err != nil {
		t.Fatal(err)
	}
	defer machine.Release()

	// NOTE: Using the USB 1.1 controller so tests can run on the open-sourced
	//       VirtualBox (no extension packs).
	controller, err := machine.AddUsbController("Controller: USB",
		UsbControllerType_Ohci)
	if err != nil {
		t.Fatal(err)
	}
	defer controller.Release()

	name, err := controller.GetName()
	if err != nil {
		t.Error(err)
	} else if name != "Controller: USB" {
		t.Error("Wrong controller name: ", name)
	}

	controllerType, err := controller.GetType()
	if err != nil {
		t.Error(err)
	} else if controllerType != UsbControllerType_Ohci {
		t.Error("Wrong controller type: ", controllerType)
	}

	major, minor, err := controller.GetStandard()
	if err != nil {
		t.Error(err)
	} else if major != 1 || minor != 1 {
		t.Errorf("Wrong controller USB standard: %d.%d\n", major, minor)
	}
}
