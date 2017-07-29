package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/bios_settings.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

// Controls a running VM.
type BiosSettings struct {
	csettings *C.IBIOSSettings
}

func (settings *BiosSettings) GetLogoImagePath() (string, error) {
	var clogoImagePath *C.char
	result := C.GoVboxGetBiosSettingsLogoImagePath(settings.csettings, &clogoImagePath)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get IBiosSettings logo image path: %x", result))
	}
	logoImagePath := C.GoString(clogoImagePath)
	C.GoVboxUtf8Free(clogoImagePath)
	return logoImagePath, nil
}

func (settings *BiosSettings) SetLogoImagePath(logoImagePath string) error {
	clogoImagePath := C.CString(logoImagePath)
	result := C.GoVboxSetBiosSettingsLogoImagePath(settings.csettings, clogoImagePath)
	C.free(unsafe.Pointer(clogoImagePath))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings logo image path: %x", result))
	}
	return nil
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

func (settings *BiosSettings) GetLogoDisplayTime() (time.Duration, error) {
	var cdisplayTime C.PRUint32
	result := C.GoVboxGetBiosSettingsLogoDisplayTime(settings.csettings,
		&cdisplayTime)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings log display time: %x", result))
	}
	return time.Duration(cdisplayTime) * time.Millisecond, nil
}

func (settings *BiosSettings) SetLogoDisplayTime(displayTime time.Duration) error {
	var cdisplayTime = C.PRUint32(displayTime.Nanoseconds() / 1000000)
	result := C.GoVboxSetBiosSettingsLogoDisplayTime(settings.csettings,
		cdisplayTime)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings log display time: %x", result))
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

func (settings *BiosSettings) GetACPIEnabled() (bool, error) {
	var cacpiEnabled C.PRBool
	result := C.GoVboxGetBiosSettingsACPIEnabled(settings.csettings,
		&cacpiEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings ACPI state: %x", result))
	}
	return cacpiEnabled != 0, nil
}

func (settings *BiosSettings) SetACPIEnabled(acpiEnabled bool) error {
	cacpiEnabled := C.PRBool(0)
	if acpiEnabled {
		cacpiEnabled = C.PRBool(1)
	}
	result := C.GoVboxSetBiosSettingsACPIEnabled(settings.csettings,
		cacpiEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings ACPI state: %x", result))
	}
	return nil
}

func (settings *BiosSettings) GetIOAPICEnabled() (bool, error) {
	var cioapicEnabled C.PRBool
	result := C.GoVboxGetBiosSettingsIOAPICEnabled(settings.csettings,
		&cioapicEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings IOAPIC state: %x", result))
	}
	return cioapicEnabled != 0, nil
}

func (settings *BiosSettings) SetIOAPICEnabled(ioapicEnabled bool) error {
	cioapicEnabled := C.PRBool(0)
	if ioapicEnabled {
		cioapicEnabled = C.PRBool(1)
	}
	result := C.GoVboxSetBiosSettingsIOAPICEnabled(settings.csettings,
		cioapicEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings IOAPIC state: %x", result))
	}
	return nil
}

func (settings *BiosSettings) GetAPICMode() (APICMode, error) {
	var capicMode C.PRUint32
	result := C.GoVboxGetBiosSettingsAPICMode(settings.csettings,
		&capicMode)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings APIC mode: %x", result))
	}
	return APICMode(capicMode), nil
}

func (settings *BiosSettings) SetAPICMode(apicMode APICMode) error {
	result := C.GoVboxSetBiosSettingsAPICMode(settings.csettings,
		C.PRUint32(apicMode))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings APIC mode: %x", result))
	}
	return nil
}

func (settings *BiosSettings) GetTimeOffset() (time.Duration, error) {
	var ctimeOffset C.PRInt64
	result := C.GoVboxGetBiosSettingsTimeOffset(settings.csettings,
		&ctimeOffset)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings time offset: %x", result))
	}
	return time.Duration(ctimeOffset) * time.Millisecond, nil
}

func (settings *BiosSettings) SetTimeOffset(timeOffset time.Duration) error {
	var ctimeOffset = C.PRInt64(timeOffset.Nanoseconds() / 1000000)
	result := C.GoVboxSetBiosSettingsTimeOffset(settings.csettings,
		ctimeOffset)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings time offset: %x", result))
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

func (settings *BiosSettings) GetPXEDebugEnabled() (bool, error) {
	var cPXEDebugEnabled C.PRBool
	result := C.GoVboxGetBiosSettingsPXEDebugEnabled(settings.csettings,
		&cPXEDebugEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IBiosSettings PXE debug state: %x", result))
	}
	return cPXEDebugEnabled != 0, nil
}

func (settings *BiosSettings) SetPXEDebugEnabled(PXEDebugEnabled bool) error {
	cPXEDebugEnabled := C.PRBool(0)
	if PXEDebugEnabled {
		cPXEDebugEnabled = C.PRBool(1)
	}
	result := C.GoVboxSetBiosSettingsPXEDebugEnabled(settings.csettings,
		cPXEDebugEnabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IBiosSettings PXE debug state: %x", result))
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
