package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/keyboard.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
)

// The keyboard of a running VM.
type Keyboard struct {
	ckeyboard *C.IKeyboard
}

// PutScancodes posts keyboard scancodes to the guest OS event queue.
// It returns any error encountered.
func (keyboard *Keyboard) PutScancodes(scancodes []int) (uint, error) {
	scancodesCount := len(scancodes)
	cscancodes := make([]C.PRInt32, scancodesCount)
	for i, scancode := range scancodes {
		cscancodes[i] = C.PRInt32(scancode)
	}

	var ccodesStored C.PRUint32
	result := C.GoVboxKeyboardPutScancodes(keyboard.ckeyboard,
		C.PRUint32(scancodesCount), &cscancodes[0], &ccodesStored)
	if C.GoVboxFAILED(result) != 0 {
		return uint(ccodesStored), errors.New(
			fmt.Sprintf("Failed to post IKeyboard scancodes: %x", result))
	}
	return uint(ccodesStored), nil
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

// Initialized returns true if there is VirtualBox data associated with this.
func (keyboard *Keyboard) Initialized() bool {
	return keyboard.ckeyboard != nil
}
