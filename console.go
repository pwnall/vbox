package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/console.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
)

// Controls a running VM.
type Console struct {
  cconsole *C.IConsole
}

// GetMachine obtains the VM associated with this set of VM controls.
// It returns a new Machine instance and any error encountered.
func (console *Console) GetMachine() (Machine, error) {
  var machine Machine
  result := C.GoVboxGetConsoleMachine(console.cconsole, &machine.cmachine)
  if C.GoVboxFAILED(result) != 0 {
    return machine, errors.New(
        fmt.Sprintf("Failed to get IConsole machine: %x", result))
  }
  return machine, nil
}

// PowerDown starts forcibly powering off the controlled VM.
// It returns a Progress and any error encountered.
func (console *Console) PowerDown() (Progress, error) {
  var progress Progress
  result := C.GoVboxConsolePowerDown(console.cconsole, &progress.cprogress)
  if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
    return progress, errors.New(
        fmt.Sprintf("Failed to power down VM via IConsole: %x", result))
  }
  return progress, nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (console *Console) Release() error {
  if console.cconsole != nil {
    result := C.GoVboxIConsoleRelease(console.cconsole)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IConsole: %x", result))
    }
    console.cconsole = nil
  }
  return nil
}
