package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/display.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	"unsafe"
)

// The display of a running VM.
type Display struct {
	cdisplay *C.IDisplay
}

// Represents the resolution of a running VM's display.
type Resolution struct {
	Width              uint
	Height             uint
	BitsPerPixel       uint
	XOrigin            int
	YOrigin            int
	GuestMonitorStatus uint
}

// GetScreenResolution reads the VM display's resolution.
// It returns any error encountered.
func (display *Display) GetScreenResolution(screenId uint,
	resolution *Resolution) error {
	var cwidth, cheight, cbitsPerPixel, cguestMonitorStatus C.PRUint32
	var cxOrigin, cyOrigin C.PRInt32

	result := C.GoVboxDisplayGetScreenResolution(display.cdisplay,
		C.PRUint32(screenId), &cwidth, &cheight, &cbitsPerPixel, &cxOrigin,
		&cyOrigin, &cguestMonitorStatus)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to get IDisplay resolution: %x", result))
	}

	resolution.Width = uint(cwidth)
	resolution.Height = uint(cheight)
	resolution.BitsPerPixel = uint(cbitsPerPixel)
	resolution.XOrigin = int(cxOrigin)
	resolution.YOrigin = int(cyOrigin)
	resolution.GuestMonitorStatus = uint(cguestMonitorStatus)
	return nil
}

// TakeScreenShot takes a screenshot of the VM's display.
// The slice passed in must be big enough to receive the image data, which is
// RGBA (4 bytes per pixel).
// It returns any error encountered.
func (display *Display) TakeScreenShot(screenId uint, imageData []byte,
	width uint, height uint, bitmapFormat uint) ([]byte, error) {
	dataSize := int(width * height * 4)
	if dataSize > cap(imageData) {
		return nil, errors.New(fmt.Sprintf(
			"Insufficient slice capacity %d, screenshot needs at least %d",
			cap(imageData), dataSize))
	}

	imageData = imageData[:dataSize]
	cdataPtr := (*C.PRUint8)(unsafe.Pointer(&imageData[0]))
	result := C.GoVboxDisplayTakeScreenShot(display.cdisplay,
		C.PRUint32(screenId), cdataPtr, C.PRUint32(width), C.PRUint32(height),
		C.PRUint32(bitmapFormat))
	if C.GoVboxFAILED(result) != 0 {
		return nil, errors.New(
			fmt.Sprintf("Failed to take IDisplay fast screenshot: %x", result))
	}
	return imageData, nil
}

// TakeScreenShotToArray takes a screenshot of the VM's display.
// It returns a byte slice encoding the image as RGBA, and any error
// encountered.
func (display *Display) TakeScreenShotToArray(screenId uint,
	width uint, height uint) ([]byte, error) {
	var cdataPtr *C.PRUint8
	var dataSize C.PRUint32

	result := C.GoVboxDisplayTakeScreenShotToArray(display.cdisplay,
		C.PRUint32(screenId), C.PRUint32(width), C.PRUint32(height),
		C.PRUint32(C.BitmapFormat_RGBA), &dataSize, &cdataPtr)
	if C.GoVboxFAILED(result) != 0 {
		return nil, errors.New(
			fmt.Sprintf("Failed to take IDisplay screenshot: %x", result))
	}

	data := C.GoBytes(unsafe.Pointer(cdataPtr), C.int(dataSize))
	C.GoVboxArrayOutFree(unsafe.Pointer(cdataPtr))
	return data, nil
}

// TakeScreenShotPNGToArray takes a screenshot of the VM's display.
// It returns a byte slice encoding the image as PNG, and any error
// encountered.
func (display *Display) TakeScreenShotPngToArray(screenId uint,
	width uint, height uint) ([]byte, error) {
	var cdataPtr *C.PRUint8
	var dataSize C.PRUint32

	result := C.GoVboxDisplayTakeScreenShotToArray(display.cdisplay,
		C.PRUint32(screenId), C.PRUint32(width), C.PRUint32(height),
		C.PRUint32(C.BitmapFormat_PNG), &dataSize, &cdataPtr)
	if C.GoVboxFAILED(result) != 0 {
		return nil, errors.New(
			fmt.Sprintf("Failed to take IDisplay PNG screenshot: %x", result))
	}

	data := C.GoBytes(unsafe.Pointer(cdataPtr), C.int(dataSize))
	C.GoVboxArrayOutFree(unsafe.Pointer(cdataPtr))
	return data, nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (display *Display) Release() error {
	if display.cdisplay != nil {
		result := C.GoVboxIDisplayRelease(display.cdisplay)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(
				fmt.Sprintf("Failed to release IDisplay: %x", result))
		}
		display.cdisplay = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (display *Display) Initialized() bool {
	return display.cdisplay != nil
}
