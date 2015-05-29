package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/storage_controller.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
  "unsafe"
)

// The description of a VirtualBox storage controller
type StorageController struct {
  ccontroller *C.IStorageController
}

// GetName returns the controller's name.
// The controller's name identifies it in AttachDevice() calls.
// It returns a string and any error encountered.
func (controller *StorageController) GetName() (string, error) {
  var cname *C.char
  result := C.GoVboxGetStorageControllerName(controller.ccontroller, &cname)
  if C.GoVboxFAILED(result) != 0 || cname == nil {
    return "", errors.New(
        fmt.Sprintf("Failed to get IStorageController name: %x", result))
  }

  name := C.GoString(cname)
  C.GoVboxUtf8Free(cname)
  return name, nil
}

// GetBus returns the controller's bus type.
// It returns a number and any error encountered.
func (controller* StorageController) GetBus() (StorageBus, error) {
  var cbus C.PRUint32

  result := C.GoVboxGetStorageControllerBus(controller.ccontroller, &cbus)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IStorageController percent: %x", result))
  }
  return StorageBus(cbus), nil
}

// GetType returns the controller's type.
// It returns a number and any error encountered.
func (controller* StorageController) GetType() (StorageControllerType, error) {
  var ctype C.PRUint32

  result := C.GoVboxGetStorageControllerType(controller.ccontroller, &ctype)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IStorageController type: %x", result))
  }
  return StorageControllerType(ctype), nil
}

// SetType changes the controller's type.
// It returns a number and any error encountered.
func (controller* StorageController) SetType(
    controllerType StorageControllerType) error {
  result := C.GoVboxSetStorageControllerType(controller.ccontroller,
      C.PRUint32(controllerType))
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(
        fmt.Sprintf("Failed to set IStorageController type: %x", result))
  }
  return nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (controller *StorageController) Release() error {
  if controller.ccontroller != nil {
    result := C.GoVboxIStorageControllerRelease(controller.ccontroller)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IStorageController: %x", result))
    }
    controller.ccontroller = nil
  }
  return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (controller *StorageController) Initialized() bool {
  return controller.ccontroller != nil
}


// AddStorageController attaches a storage controller to a VirtualBox VM.
// It returns the created StorageController and any error encountered.
func (machine *Machine) AddStorageController(
    name string, connectionType StorageBus) (StorageController, error) {
  var controller StorageController
  if err := Init(); err != nil {
    return controller, err
  }

  cname := C.CString(name)
  result := C.GoVboxMachineAddStorageController(machine.cmachine, cname,
      C.PRUint32(connectionType), &controller.ccontroller)
  C.free(unsafe.Pointer(cname))

  if C.GoVboxFAILED(result) != 0 || controller.ccontroller == nil {
    return controller, errors.New(fmt.Sprintf(
        "Failed to add IStorageController: %x", result))
  }
  return controller, nil
}
