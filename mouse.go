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

// GetAbsoluteSupported checks if the guest handles absolute mouse positioning.
// If it returns false, PutMouseEventAbsolute() is a no-op.
// It returns a number and any error encountered.
func (mouse* Mouse) GetAbsoluteSupported() (bool, error) {
  var csupported C.PRBool

  result := C.GoVboxGetMouseAbsoluteSupported(mouse.cmouse, &csupported)
  if C.GoVboxFAILED(result) != 0 {
    return false, errors.New(
        fmt.Sprintf("Failed to get IMouse absoluteSupported: %x", result))
  }
  return csupported != 0, nil
}

// GetRelativeSupported checks if the guest handles relative mouse positioning.
// If it returns false, PutMouseEventRelative() is a no-op.
// It returns a number and any error encountered.
func (mouse* Mouse) GetRelativeSupported() (bool, error) {
  var csupported C.PRBool

  result := C.GoVboxGetMouseRelativeSupported(mouse.cmouse, &csupported)
  if C.GoVboxFAILED(result) != 0 {
    return false, errors.New(
        fmt.Sprintf("Failed to get IMouse relativeSupported: %x", result))
  }
  return csupported != 0, nil
}

// PutMouseEvent posts a mouse event to the guest OS event queue.
// dz represents vertical mouse wheel moves (rotations), with positive numbers
// for clockwise rotations. dw represents horizontal mouse wheel moves, with
// positive numbers for movements to the left.
// It returns any error encountered.
func (mouse* Mouse) PutMouseEvent(dx int, dy int, dz int, dw int,
    buttonState int) error {
  result := C.GoVboxPutMouseEvent(mouse.cmouse, C.PRInt32(dx), C.PRInt32(dy),
      C.PRInt32(dz), C.PRInt32(dw), C.PRInt32(buttonState))
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(fmt.Sprintf("Failed to post IMouse event: %x", result))
  }
  return nil
}

// PutMouseEventAbsolute posts a mouse event to the guest OS event queue.
// dz represents vertical mouse wheel moves (rotations), with positive numbers
// for clockwise rotations. dw represents horizontal mouse wheel moves, with
// positive numbers for movements to the left.
// It returns any error encountered.
func (mouse* Mouse) PutMouseEventAbsolute(x int, y int, dz int, dw int,
    buttonState int) error {
  result := C.GoVboxPutMouseEventAbsolute(mouse.cmouse, C.PRInt32(x),
      C.PRInt32(y), C.PRInt32(dz), C.PRInt32(dw), C.PRInt32(buttonState))
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(
        fmt.Sprintf("Failed to post IMouse absolute event: %x", result))
  }
  return nil
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
