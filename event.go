package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/event.c"

*/
import "C"
import (
	"errors"
	"fmt"
) // cgo's virtual package

// The description of a VirtualBox machine
type Event struct {
	cevent *C.IEvent
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (event *Event) Release() error {
	if event.cevent != nil {
		result := C.GoVboxIEventRelease(event.cevent)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release IEvent: %x", result))
		}
		event.cevent = nil
	}
	return nil
}

func (event *Event) GetType() (uint32, error) {
	var eventType C.PRUint32
	result := C.GoVboxEventGetType(event.cevent, &eventType)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(fmt.Sprintf("Failed to get event type: %x", result))
	}
	return uint32(eventType), nil
}

func (event *Event) GetSource() (EventSource, error) {
	var source EventSource
	result := C.GoVboxEventGetSource(event.cevent, &source.ceventSource)
	if C.GoVboxFAILED(result) != 0 {
		return source, errors.New(fmt.Sprintf("Failed to get event source: %x", result))
	}
	return source, nil
}

func (event *Event) GetWaitable() (bool, error) {
	var waitable C.PRBool
	result := C.GoVboxEventGetWaitable(event.cevent, &waitable)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(fmt.Sprintf("Failed to get event waitable: %x", result))
	}
	return waitable != 0, nil
}

func (event *Event) SetProcessed() error {
	result := C.GoVboxEventSetProcessed(event.cevent)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(fmt.Sprintf("Failed to set event as processed: %x", result))
	}
	return nil
}

func (event *Event) WaitProcessed(timeout int32) (bool, error) {
	var processed C.PRBool
	result := C.GoVboxEventWaitProcessed(event.cevent, C.PRInt32(timeout), &processed)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(fmt.Sprintf("Failed to wait for event to processed: %x", result))
	}
	return processed != 0, nil
}

type GuestPropertyChangedEvent struct {
	Event
	cevent *C.IGuestPropertyChangedEvent
}

func (event *GuestPropertyChangedEvent) GetName() (string, error) {
	var cname *C.char
	result := C.GoVboxGuestPropertyChangedEventGetName(event.cevent, &cname)
	if C.GoVboxFAILED(result) != 0 || cname == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IGuestPropertyChangedEvent property name: %x", result))
	}

	name := C.GoString(cname)
	C.GoVboxUtf8Free(cname)
	return name, nil
}

func (event *GuestPropertyChangedEvent) GetValue() (string, error) {
	var cvalue *C.char
	result := C.GoVboxGuestPropertyChangedEventGetValue(event.cevent, &cvalue)
	if C.GoVboxFAILED(result) != 0 || cvalue == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IGuestPropertyChangedEvent property value: %x", result))
	}

	value := C.GoString(cvalue)
	C.GoVboxUtf8Free(cvalue)
	return value, nil
}

func (event *GuestPropertyChangedEvent) GetFlags() (string, error) {
	var cflags *C.char
	result := C.GoVboxGuestPropertyChangedEventGetFlags(event.cevent, &cflags)
	if C.GoVboxFAILED(result) != 0 || cflags == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IGuestPropertyChangedEvent property flags: %x", result))
	}

	flags := C.GoString(cflags)
	C.GoVboxUtf8Free(cflags)
	return flags, nil
}

func NewGuestPropertyChangedEvent(event *Event) (GuestPropertyChangedEvent, error) {
	var guestPropEvent GuestPropertyChangedEvent

	result := C.GoVboxEventGetGuestPropertyChangedEvent(event.cevent, &guestPropEvent.cevent)
	if C.GoVboxFAILED(result) != 0 || guestPropEvent.cevent == nil {
		return guestPropEvent, errors.New(
			fmt.Sprintf("Failed to convert to IGuestPropertyChangedEvent event: %x", result))
	}

	return GuestPropertyChangedEvent{
		Event:  *event,
		cevent: guestPropEvent.cevent,
	}, nil
}
