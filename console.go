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

// GetDisplay obtains the display of the VM controlled by this.
// It returns a new Display instance and any error encountered.
func (console *Console) GetDisplay() (Display, error) {
  var display Display
  result := C.GoVboxGetConsoleDisplay(console.cconsole, &display.cdisplay)
  if C.GoVboxFAILED(result) != 0 || display.cdisplay == nil {
    return display, errors.New(
        fmt.Sprintf("Failed to get IConsole display: %x", result))
  }
  return display, nil
}

// GetKeyboard obtains the keyboard of the VM controlled by this.
// It returns a new Keyboard instance and any error encountered.
func (console *Console) GetKeyboard() (Keyboard, error) {
  var keyboard Keyboard
  result := C.GoVboxGetConsoleKeyboard(console.cconsole, &keyboard.ckeyboard)
  if C.GoVboxFAILED(result) != 0 || keyboard.ckeyboard == nil {
    return keyboard, errors.New(
        fmt.Sprintf("Failed to get IConsole keyboard: %x", result))
  }
  return keyboard, nil
}

// GetMouse obtains the mouse of the VM controlled by this.
// It returns a new Mouse instance and any error encountered.
func (console *Console) GetMouse() (Mouse, error) {
  var mouse Mouse
  result := C.GoVboxGetConsoleMouse(console.cconsole, &mouse.cmouse)
  if C.GoVboxFAILED(result) != 0 || mouse.cmouse == nil {
    return mouse, errors.New(
        fmt.Sprintf("Failed to get IConsole mouse: %x", result))
  }
  return mouse, nil
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

// Initialized returns true if there is VirtualBox data associated with this.
func (console *Console) Initialized() bool {
  return console.cconsole != nil
}
