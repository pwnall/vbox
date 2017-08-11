package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/system_properties.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	//"reflect"
	//"unsafe"
)

// The description of a VirtualBox storage medium
type SystemProperties struct {
	cprops *C.ISystemProperties
}

// GetMaxGuestRAM reads the maximum allowed amount of RAM on a guest VM.
// It returns a megabyte quantity and any error encountered.
func (props *SystemProperties) GetMaxGuestRam() (uint, error) {
	var cmaxRam C.ULONG
	result := C.GoVboxGetSystemPropertiesMaxGuestRAM(props.cprops, &cmaxRam)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get ISystemProperties max RAM: %x", result))
	}
	return uint(cmaxRam), nil
}

// GetMaxGuestVRAM reads the maximum allowed amount of video RAM on a guest VM.
// It returns a megabyte quantity and any error encountered.
func (props *SystemProperties) GetMaxGuestVram() (uint, error) {
	var cmaxVram C.ULONG
	result := C.GoVboxGetSystemPropertiesMaxGuestVRAM(props.cprops, &cmaxVram)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get ISystemProperties max VRAM: %x", result))
	}
	return uint(cmaxVram), nil
}

// GetMaxGuestCpuCount reads the maximum number of CPUs on a guest VM.
// It returns a number and any error encountered.
func (props *SystemProperties) GetMaxGuestCpuCount() (uint, error) {
	var cmaxCpus C.ULONG
	result := C.GoVboxGetSystemPropertiesMaxGuestCpuCount(props.cprops,
		&cmaxCpus)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get ISystemProperties max CPUs: %x", result))
	}
	return uint(cmaxCpus), nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (props *SystemProperties) Release() error {
	if props.cprops != nil {
		result := C.GoVboxISystemPropertiesRelease(props.cprops)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(
				fmt.Sprintf("Failed to release ISystemProperties: %x", result))
		}
		props.cprops = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (props *SystemProperties) Initialized() bool {
	return props.cprops != nil
}

// GetSystemProperties fetches the VirtualBox system properties.
// It returns the a new SystemProperties instance and any error encountered.
func GetSystemProperties() (SystemProperties, error) {
	var props SystemProperties
	if err := Init(); err != nil {
		return props, err
	}

	result := C.GoVboxGetSystemProperties(cbox, &props.cprops)
	if C.GoVboxFAILED(result) != 0 || props.cprops == nil {
		return props, errors.New(
			fmt.Sprintf("Failed to create IMachine: %x", result))
	}
	return props, nil
}
