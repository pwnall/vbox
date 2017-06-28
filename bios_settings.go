package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/bios_settings.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
)

// Controls a running VM.
type BiosSettings struct {
	csettings *C.IBIOSSettings
}

// GetLogoFadeIn returns true if the BIOS logo fades in during boot.
// It returns a boolean and any error encountered.
func (settings *BiosSettings) GetLogoFadeIn() (bool, error) {
	var clogoFadeIn C.PRBool
	result := C.GoVboxGetBiosSettingsLogoFadeIn(settings.csettings, &clogoFadeIn)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings logoFadeIn: %x", result))
	}
	return clogoFadeIn != 0, nil
}

// SetLogoFadeIn sets whether the BIOS logo fades in during boot.
// It any error encountered.
func (settings *BiosSettings) SetLogoFadeIn(logoFadeIn bool) error {
	clogoFadeIn := C.PRBool(0)
	if logoFadeIn {
		clogoFadeIn = C.PRBool(1)
	}
	result := C.GoVboxSetBiosSettingsLogoFadeIn(settings.csettings, clogoFadeIn)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings logoFadeIn: %x", result))
	}
	return nil
}

// GetLogoFadeOut returns true if the BIOS logo fades out during boot.
// It returns a boolean and any error encountered.
func (settings *BiosSettings) GetLogoFadeOut() (bool, error) {
	var clogoFadeOut C.PRBool
	result := C.GoVboxGetBiosSettingsLogoFadeOut(settings.csettings,
		&clogoFadeOut)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings logoFadeOut: %x", result))
	}
	return clogoFadeOut != 0, nil
}

// SetLogoFadeOut sets whether the BIOS logo fades out during boot.
// It any error encountered.
func (settings *BiosSettings) SetLogoFadeOut(logoFadeOut bool) error {
	clogoFadeOut := C.PRBool(0)
	if logoFadeOut {
		clogoFadeOut = C.PRBool(1)
	}
	result := C.GoVboxSetBiosSettingsLogoFadeOut(settings.csettings,
		clogoFadeOut)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings logoFadeOut: %x", result))
	}
	return nil
}

// GetBootMenuMode returns true if the BIOS logo fades out during boot.
// It returns a boolean and any error encountered.
func (settings *BiosSettings) GetBootMenuMode() (BootMenuMode, error) {
	var cmenuMode C.PRUint32
	result := C.GoVboxGetBiosSettingsBootMenuMode(settings.csettings,
		&cmenuMode)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings boot menu mode: %x", result))
	}
	return BootMenuMode(cmenuMode), nil
}

// SetBootMenuMode sets whether the BIOS logo fades out during boot.
// It any error encountered.
func (settings *BiosSettings) SetBootMenuMode(menuMode BootMenuMode) error {
	result := C.GoVboxSetBiosSettingsBootMenuMode(settings.csettings,
		C.PRUint32(menuMode))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings boot menu mode: %x", result))
	}
	return nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (settings *BiosSettings) Release() error {
	if settings.csettings != nil {
		result := C.GoVboxIBiosSettingsRelease(settings.csettings)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(
				fmt.Sprintf("Failed to release IBiosSettings: %x", result))
		}
		settings.csettings = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (settings *BiosSettings) Initialized() bool {
	return settings.csettings != nil
}

// GetBiosSettings obtains the controls for the VM associated with this machine.
// The call fails unless the VM associated with this machine has started.
// It returns a new BiosSettings instance and any error encountered.
func (machine *Machine) GetBiosSettings() (BiosSettings, error) {
	var settings BiosSettings
	result := C.GoVboxGetMachineBIOSSettings(machine.cmachine,
		&settings.csettings)
	if C.GoVboxFAILED(result) != 0 {
		return settings, errors.New(
			fmt.Sprintf("Failed to get IMachine BIOS settings: %x", result))
	}
	return settings, nil
}
