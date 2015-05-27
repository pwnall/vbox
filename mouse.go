package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/mouse.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
)

// The mouse of a running VM.
type Mouse struct {
  cmouse *C.IMouse
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (mouse *Mouse) Release() error {
  if mouse.cmouse != nil {
    result := C.GoVboxIMouseRelease(mouse.cmouse)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IMouse: %x", result))
    }
    mouse.cmouse = nil
  }
  return nil
}
