package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/event_listener.c"
*/
import "C"
import (
	"errors"
	"fmt"
) // cgo's virtual package

// The description of a VirtualBox machine
type EventListener struct {
	ceventListener *C.IEventListener
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (eventListener *EventListener) Release() error {
	if eventListener.ceventListener != nil {
		result := C.GoVboxIEventListenerRelease(eventListener.ceventListener)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release IEventListener: %x", result))
		}
		eventListener.ceventListener = nil
	}
	return nil
}
