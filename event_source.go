package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/event_source.c"
*/
import "C"
import (
	"errors"
	"fmt"
) // cgo's virtual package

// The description of a VirtualBox machine
type EventSource struct {
	ceventSource *C.IEventSource
}

func (eventSource *EventSource) CreateListener() (EventListener, error) {
	var eventListener EventListener
	result := C.GoVboxEventSourceCreateListener(eventSource.ceventSource,
		&eventListener.ceventListener)
	if C.GoVboxFAILED(result) != 0 || eventListener.ceventListener == nil {
		return eventListener, errors.New(
			fmt.Sprintf("Failed to create IEventListener: %x", result))
	}
	return eventListener, nil
}

func (eventSource *EventSource) CreateAggregator(subordinates []EventSource) (EventSource, error) {
	var csubordinatesSlice []*C.IEventSource
	var csubordinates **C.IEventSource
	var aggregator EventSource
	if len(subordinates) > 0 {
		csubordinatesSlice = make([]*C.IEventSource, len(subordinates))
		for i, subordinate := range subordinates {
			csubordinatesSlice[i] = subordinate.ceventSource
		}
		csubordinates = &csubordinatesSlice[0]
	}
	result := C.GoVboxEventSourceCreateAggregator(eventSource.ceventSource,
		C.PRUint32(len(subordinates)), csubordinates, &aggregator.ceventSource)
	if C.GoVboxFAILED(result) != 0 || aggregator.ceventSource == nil {
		return aggregator, errors.New(
			fmt.Sprintf("Failed to create aggregator: %x", result))
	}
	return aggregator, nil
}

func (eventSource *EventSource) RegisterListener(eventListener EventListener,
	interesting []uint32, active bool) error {
	cactive := C.PRBool(0)
	if active {
		cactive = C.PRBool(1)
	}
	var cinteresting *C.PRUint32
	if len(interesting) > 0 {
		cinterestingSlice := make([]C.PRUint32, len(interesting))
		for i, interest := range interesting {
			cinterestingSlice[i] = C.PRUint32(interest)
		}
		cinteresting = &cinterestingSlice[0]
	}
	result := C.GoVboxEventSourceRegisterListener(eventSource.ceventSource,
		eventListener.ceventListener, C.PRUint32(len(interesting)), cinteresting,
		cactive)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to register IEventListener: %x", result))
	}
	return nil
}

func (eventSource *EventSource) UnregisterListener(eventListener EventListener) error {
	result := C.GoVboxEventSourceUnregisterListener(eventSource.ceventSource,
		eventListener.ceventListener)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to unregister IEventListener: %x", result))
	}
	return nil
}

func (eventSource *EventSource) FireEvent(eventListener EventListener,
	event Event, timeout int32) (bool, error) {
	var fired C.PRBool
	result := C.GoVboxEventSourceFireEvent(eventSource.ceventSource,
		event.cevent, C.PRInt32(timeout), &fired)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to fire event: %x", result))
	}
	return fired != 0, nil
}

func (eventSource *EventSource) GetEvent(eventListener EventListener,
	timeout int32) (*Event, error) {
	var event Event
	result := C.GoVboxEventSourceGetEvent(eventSource.ceventSource,
		eventListener.ceventListener, C.PRInt32(timeout), &event.cevent)
	if C.GoVboxFAILED(result) != 0 {
		return nil, errors.New(fmt.Sprintf("Failed to get event: %x", result))
	}
	if event.cevent == nil {
		return nil, nil
	}
	return &event, nil
}

func (eventSource *EventSource) EventProcessed(eventListener EventListener,
	event Event) error {
	result := C.GoVboxEventSourceEventProcessed(eventSource.ceventSource,
		eventListener.ceventListener, event.cevent)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(fmt.Sprintf("Failed to process event: %x", result))
	}
	return nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (eventSource *EventSource) Release() error {
	if eventSource.ceventSource != nil {
		result := C.GoVboxIEventSourceRelease(eventSource.ceventSource)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release IEventSource: %x", result))
		}
		eventSource.ceventSource = nil
	}
	return nil
}
