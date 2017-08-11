package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/usb_controller.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	"unsafe"
)

// The description of a VirtualBox USB controller
type UsbController struct {
	ccontroller *C.IUSBController
}

// GetName returns the controller's name.
// It returns a string and any error encountered.
func (controller *UsbController) GetName() (string, error) {
	var cname *C.char
	result := C.GoVboxGetUsbControllerName(controller.ccontroller, &cname)
	if C.GoVboxFAILED(result) != 0 || cname == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IUsbController name: %x", result))
	}

	name := C.GoString(cname)
	C.GoVboxUtf8Free(cname)
	return name, nil
}

// GetStandard returns the USB standard supported by the controller.
// It returns two numbers (the major and minor versions of the standard), and
// any error encountered.
func (controller *UsbController) GetStandard() (int, int, error) {
	var cstandard C.PRUint16

	result := C.GoVboxGetUsbControllerStandard(controller.ccontroller,
		&cstandard)
	if C.GoVboxFAILED(result) != 0 {
		return 0, 0, errors.New(
			fmt.Sprintf("Failed to get IUsbController percent: %x", result))
	}

	standard := int(cstandard)

	return standard >> 8, standard & 0xff, nil
}

// GetType returns the controller's type.
// It returns a number and any error encountered.
func (controller *UsbController) GetType() (UsbControllerType, error) {
	var ctype C.PRUint32

	result := C.GoVboxGetUsbControllerType(controller.ccontroller, &ctype)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IUsbController type: %x", result))
	}
	return UsbControllerType(ctype), nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (controller *UsbController) Release() error {
	if controller.ccontroller != nil {
		result := C.GoVboxIUSBControllerRelease(controller.ccontroller)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(
				fmt.Sprintf("Failed to release IUsbController: %x", result))
		}
		controller.ccontroller = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (controller *UsbController) Initialized() bool {
	return controller.ccontroller != nil
}

// AddUsbController attaches a storage controller to a VirtualBox VM.
// It returns the created UsbController and any error encountered.
func (machine *Machine) AddUsbController(
	name string, controllerType UsbControllerType) (UsbController, error) {
	var controller UsbController
	if err := Init(); err != nil {
		return controller, err
	}

	cname := C.CString(name)
	result := C.GoVboxMachineAddUsbController(machine.cmachine, cname,
		C.PRUint32(controllerType), &controller.ccontroller)
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 || controller.ccontroller == nil {
		return controller, errors.New(fmt.Sprintf(
			"Failed to add IUsbController: %x", result))
	}
	return controller, nil
}
