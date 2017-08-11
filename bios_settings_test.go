package vbox

import (
	"testing"
)

func TestBiosSettings(t *testing.T) {
	machine, err := CreateMachine("", "pwnall_vbox_test", "Linux", "")
	if err != nil {
		t.Fatal(err)
	}
	defer machine.Release()

	settings, err := machine.GetBiosSettings()
	if err != nil {
		t.Fatal(err)
	}
	defer settings.Release()

	if err := settings.SetLogoFadeIn(true); err != nil {
		t.Error(err)
	}
	logoFadeIn, err := settings.GetLogoFadeIn()
	if err != nil {
		t.Error(err)
	} else if logoFadeIn != true {
		t.Error("Setting logoFadeIn failed, got: ", logoFadeIn)
	}

	if err := settings.SetLogoFadeOut(true); err != nil {
		t.Error(err)
	}
	logoFadeOut, err := settings.GetLogoFadeOut()
	if err != nil {
		t.Error(err)
	} else if logoFadeOut != true {
		t.Error("Setting logoFadeOut failed, got: ", logoFadeOut)
	}

	if err := settings.SetBootMenuMode(BootMenuMode_Disabled); err != nil {
		t.Error(err)
	}
	bootMenuMode, err := settings.GetBootMenuMode()
	if err != nil {
		t.Error(err)
	} else if bootMenuMode != BootMenuMode_Disabled {
		t.Error("Setting bootMenuMode failed, got: ", bootMenuMode)
	}
}
