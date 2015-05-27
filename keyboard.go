package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/keyboard.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
)

// The keyboard of a running VM.
type Keyboard struct {
  ckeyboard *C.IKeyboard
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (keyboard *Keyboard) Release() error {
  if keyboard.ckeyboard != nil {
    result := C.GoVboxIKeyboardRelease(keyboard.ckeyboard)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IKeyboard: %x", result))
    }
    keyboard.ckeyboard = nil
  }
  return nil
}
