package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/machine.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// The description of a VirtualBox machine
type Machine struct {
	cmachine *C.IMachine
}

// GetName returns the machine's name.
// It returns a string and any error encountered.
func (machine *Machine) GetName() (string, error) {
	var cname *C.char
	result := C.GoVboxGetMachineName(machine.cmachine, &cname)
	if C.GoVboxFAILED(result) != 0 || cname == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IMachine name: %x", result))
	}

	name := C.GoString(cname)
	C.GoVboxUtf8Free(cname)
	return name, nil
}

// GetOsTypeId returns a string used to identify the guest OS type.
// It returns a string and any error encountered.
func (machine *Machine) GetOsTypeId() (string, error) {
	var cosTypeId *C.char
	result := C.GoVboxGetMachineOSTypeId(machine.cmachine, &cosTypeId)
	if C.GoVboxFAILED(result) != 0 || cosTypeId == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IMachine OS type ID: %x", result))
	}

	osTypeId := C.GoString(cosTypeId)
	C.GoVboxUtf8Free(cosTypeId)
	return osTypeId, nil
}

// GetSettingsFilePath returns the path of the machine's settings file.
// It returns a string and any error encountered.
func (machine *Machine) GetSettingsFilePath() (string, error) {
	var cpath *C.char
	result := C.GoVboxGetMachineSettingsFilePath(machine.cmachine, &cpath)
	if C.GoVboxFAILED(result) != 0 || cpath == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IMachine settings file path: %x", result))
	}

	path := C.GoString(cpath)
	C.GoVboxUtf8Free(cpath)
	return path, nil
}

// GetSettingsModified asks VirtualBox if this machine has unsaved settings.
// It returns a boolean and any error encountered.
func (machine *Machine) GetSettingsModified() (bool, error) {
	var cmodified C.PRBool
	result := C.GoVboxGetMachineSettingsModified(machine.cmachine, &cmodified)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IMachine modified flag: %x", result))
	}
	return cmodified != 0, nil
}

// SaveSettings saves a machine's modified settings.
// A new machine must have its settings saved before it can be registered.
// It returns a boolean and any error encountered.
func (machine *Machine) SaveSettings() error {
	result := C.GoVboxMachineSaveSettings(machine.cmachine)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to save IMachine settings: %x", result))
	}
	return nil
}

// GetPointingHidType returns the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) GetPointingHidType() (PointingHidType, error) {
	var ctype C.PRUint32

	result := C.GoVboxGetMachinePointingHIDType(machine.cmachine, &ctype)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine pointing HID type: %x", result))
	}
	return PointingHidType(ctype), nil
}

// SetPointingHidType changes the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) SetPointingHidType(
	pointingHidType PointingHidType) error {
	result := C.GoVboxSetMachinePointingHIDType(machine.cmachine,
		C.PRUint32(pointingHidType))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine pointing HID type: %x", result))
	}
	return nil
}

// GetMemorySize returns the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) GetMemorySize() (uint, error) {
	var cram C.PRUint32

	result := C.GoVboxGetMachineMemorySize(machine.cmachine, &cram)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine memory size: %x", result))
	}
	return uint(cram), nil
}

// SetMemorySize changes the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) SetMemorySize(ram uint) error {
	result := C.GoVboxSetMachineMemorySize(machine.cmachine, C.PRUint32(ram))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine memory size: %x", result))
	}
	return nil
}

// GetVramSize returns the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) GetVramSize() (uint, error) {
	var cvram C.PRUint32

	result := C.GoVboxGetMachineVRAMSize(machine.cmachine, &cvram)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine VRAM size: %x", result))
	}
	return uint(cvram), nil
}

// SetVramSize changes the machine's emulated mouse type.
// It returns a number and any error encountered.
func (machine *Machine) SetVramSize(vram uint) error {
	result := C.GoVboxSetMachineVRAMSize(machine.cmachine, C.PRUint32(vram))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine VRAM size: %x", result))
	}
	return nil
}

func (machine *Machine) GetCPUCount() (uint, error) {
	var ccpus C.PRUint32

	result := C.GoVboxGetMachineCPUCount(machine.cmachine, &ccpus)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine CPU count: %x", result))
	}
	return uint(ccpus), nil
}

func (machine *Machine) SetCPUCount(cpus uint) error {
	result := C.GoVboxSetMachineCPUCount(machine.cmachine, C.PRUint32(cpus))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine CPU count: %x", result))
	}
	return nil
}

func (machine *Machine) GetState() (uint, error) {
	var cstate C.PRUint32

	result := C.GoVboxGetMachineState(machine.cmachine, &cstate)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine state: %x", result))
	}
	return uint(cstate), nil
}

// Register adds this to VirtualBox's list of registered machines.
// Once a VM is registered, it becomes immutable. Its configuration can only be
// changed by creating a Session, LockMachine-ing the machine to the session,
// and obtaining the Session's version of the machine via GetMachine.
// It returns any error encountered.
func (machine *Machine) Register() error {
	// NOTE: This is a rare case where the underlying VirtualBox API call doesn't
	//       match the Go object model precisely. Register() really feels like it
	//       should belong to Machine and not to VirtualBox, because it takes a
	//       Machine argument, and VirtualBox is a singleton.
	result := C.GoVboxRegisterMachine(cbox, machine.cmachine)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(fmt.Sprintf("Failed to register IMachine: %x", result))
	}
	return nil
}

// Unregister removes this from VirtualBox's list of registered machines.
// The returned slice of Medium instances is intended to be passed to
// DeleteConfig to get all the VM's files cleaned.
// It returns an array of detached Medium instances and any error encountered.
func (machine *Machine) Unregister(cleanupMode CleanupMode) ([]Medium, error) {
	var cmediaPtr **C.IMedium
	var mediaCount C.ULONG

	result := C.GoVboxMachineUnregister(machine.cmachine,
		C.PRUint32(cleanupMode), &cmediaPtr, &mediaCount)
	if C.GoVboxFAILED(result) != 0 || (cmediaPtr == nil && mediaCount != 0) {
		return nil, errors.New(
			fmt.Sprintf("Failed to unregister IMachine: %x", result))
	}

	sliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cmediaPtr)),
		Len:  int(mediaCount),
		Cap:  int(mediaCount),
	}
	cmediaSlice := *(*[]*C.IMedium)(unsafe.Pointer(&sliceHeader))

	var media = make([]Medium, mediaCount)
	for i := range cmediaSlice {
		media[i] = Medium{cmediaSlice[i]}
	}

	C.GoVboxArrayOutFree(unsafe.Pointer(cmediaPtr))
	return media, nil
}

// DeleteConfig removes a VM's config file, and can remove its disk images.
// The Medium array is intended to be obtained from a previous Unregister call.
// It returns a Progress instance and any error encountered.
func (machine *Machine) DeleteConfig(media []Medium) (Progress, error) {
	var cmediaSlice []*C.IMedium
	var cmedia **C.IMedium
	if len(media) > 0 {
		cmediaSlice = make([]*C.IMedium, len(media))
		for i, medium := range media {
			cmediaSlice[i] = medium.cmedium
		}
		cmedia = &cmediaSlice[0]
	}

	var progress Progress
	result := C.GoVboxMachineDeleteConfig(machine.cmachine,
		C.PRUint32(len(media)), cmedia, &progress.cprogress)
	if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
		return progress, errors.New(
			fmt.Sprintf("Failed to delete IMachine config: %x", result))
	}
	return progress, nil
}

// AttachDevice connects a Medium to this VM.
// deviceSlot is 0 for IDE master and 1 for IDE slave. All other bus types use
// deviceSlot 0.
// It returns any error encountered.
func (machine *Machine) AttachDevice(controllerName string, controllerPort int,
	deviceSlot int, deviceType DeviceType, medium Medium) error {
	cname := C.CString(controllerName)
	result := C.GoVboxMachineAttachDevice(machine.cmachine, cname,
		C.PRInt32(controllerPort), C.PRInt32(deviceSlot), C.PRUint32(deviceType),
		medium.cmedium)
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to attach IMedium to IMachine: %x", result))
	}
	return nil
}

// DetachDevice disconnects a Medium from this VM.
// deviceSlot is 0 for IDE master and 1 for IDE slave. All other bus types use
// deviceSlot 0.
// It returns any error encountered.
func (machine *Machine) DetachDevice(controllerName string, controllerPort int,
	deviceSlot int) error {
	cname := C.CString(controllerName)
	result := C.GoVboxMachineDetachDevice(machine.cmachine, cname,
		C.PRInt32(controllerPort), C.PRInt32(deviceSlot))
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to attach IMedium to IMachine: %x", result))
	}
	return nil
}

// UnmountMedium ejects a removable Medium from this VM.
// It returns any error encountered.
func (machine *Machine) UnmountMedium(controllerName string,
	controllerPort int, deviceSlot int, force bool) error {
	cforce := C.PRBool(0)
	if force {
		cforce = C.PRBool(1)
	}

	cname := C.CString(controllerName)
	result := C.GoVboxMachineUnmountMedium(machine.cmachine, cname,
		C.PRInt32(controllerPort), C.PRInt32(deviceSlot), cforce)
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to unmount medium from IMachine: %x", result))
	}
	return nil
}

// GetMedium returns a Medium connected to this VM.
// It returns the requested Medium and any error encountered.
func (machine *Machine) GetMedium(controllerName string, controllerPort int,
	deviceSlot int) (Medium, error) {
	var medium Medium
	cname := C.CString(controllerName)
	result := C.GoVboxMachineGetMedium(machine.cmachine, cname,
		C.PRInt32(controllerPort), C.PRInt32(deviceSlot), &medium.cmedium)
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 || (medium.cmedium == nil) {
		return medium, errors.New(
			fmt.Sprintf("Failed to get IMedium from IMachine: %x", result))
	}
	return medium, nil
}

func (machine *Machine) GetNetworkAdapter(deviceSlot int) (NetworkAdapter, error) {
	var adapter NetworkAdapter
	result := C.GoVboxMachineGetNetworkAdapter(machine.cmachine,
		C.PRInt32(deviceSlot), &adapter.cadapter)

	if C.GoVboxFAILED(result) != 0 || (adapter.cadapter == nil) {
		return adapter, errors.New(
			fmt.Sprintf("Failed to get INetworkAdapter from IMachine: %x", result))
	}
	return adapter, nil
}

func (machine *Machine) GetAudioAdapter() (AudioAdapter, error) {
	var adapter AudioAdapter
	result := C.GoVboxMachineGetAudioAdapter(machine.cmachine,
		&adapter.caudioadapter)

	if C.GoVboxFAILED(result) != 0 || (adapter.caudioadapter == nil) {
		return adapter, errors.New(
			fmt.Sprintf("Failed to get IAudioAdapter from IMachine: %x", result))
	}
	return adapter, nil
}

// Launch swapns a process that executes this VM.
// The given session will receive a shared lock on the VM.
// It returns a Progress and any error encountered.
func (machine *Machine) Launch(session Session, uiType string,
	environment string) (Progress, error) {
	var progress Progress
	cuiType := C.CString(uiType)
	cenvironment := C.CString(environment)
	result := C.GoVboxMachineLaunchVMProcess(machine.cmachine, session.csession,
		cuiType, cenvironment, &progress.cprogress)
	C.free(unsafe.Pointer(cuiType))
	C.free(unsafe.Pointer(cenvironment))

	if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
		return progress, errors.New(
			fmt.Sprintf("Failed to launch IMachine VM: %x", result))
	}
	return progress, nil
}

func (machine *Machine) SetExtraData(key string, value string) error {
	ckey := C.CString(key)
	cvalue := C.CString(value)
	result := C.GoVboxMachineSetExtraData(machine.cmachine, ckey, cvalue)
	C.free(unsafe.Pointer(ckey))
	C.free(unsafe.Pointer(cvalue))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set extra data on IMachine: %x", result))
	}
	return nil
}

func (machine *Machine) GetAccelerate2DVideoEnabled() (bool, error) {
	var cenabled C.PRBool
	result := C.GoVboxGetMachineAccelerate2DVideoEnabled(machine.cmachine, &cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IMachine 2D video acceleration: %x", result))
	}
	return cenabled != 0, nil
}

func (machine *Machine) SetAccelerate2DVideoEnabled(enabled bool) error {
	cenabled := C.PRBool(0)
	if enabled {
		cenabled = C.PRBool(1)
	}
	result := C.GoVboxSetMachineAccelerate2DVideoEnabled(machine.cmachine, cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine 2D video acceleration: %x", result))
	}
	return nil
}

func (machine *Machine) GetAccelerate3DEnabled() (bool, error) {
	var cenabled C.PRBool
	result := C.GoVboxGetMachineAccelerate3DEnabled(machine.cmachine, &cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get IMachine 3D acceleration: %x", result))
	}
	return cenabled != 0, nil
}

func (machine *Machine) GetClipboardMode() (ClipboardMode, error) {
	var cmode C.PRUint32

	result := C.GoVboxSetClipboardMode(machine.cmachine, cmode)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine clipboard mode: %x", result))
	}
	return ClipboardMode(cmode), nil
}

func (machine *Machine) SetClipboardMode(mode ClipboardMode) error {
	result := C.GoVboxSetClipboardMode(machine.cmachine, C.PRUint32(mode))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine clipboard mode: %x", result))
	}
	return nil
}

func (machine *Machine) GetDnDMode() (DnDMode, error) {
	var cmode C.PRUint32

	result := C.GoVboxSetDnDMode(machine.cmachine, cmode)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get IMachine DnD mode: %x", result))
	}
	return DnDMode(cmode), nil
}

func (machine *Machine) SetDnDMode(mode DnDMode) error {
	result := C.GoVboxSetDnDMode(machine.cmachine, C.PRUint32(mode))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine DnD mode: %x", result))
	}
	return nil
}

func (machine *Machine) SetAccelerate3DEnabled(enabled bool) error {
	cenabled := C.PRBool(0)
	if enabled {
		cenabled = C.PRBool(1)
	}
	result := C.GoVboxSetMachineAccelerate3DEnabled(machine.cmachine, cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set IMachine 3D acceleration: %x", result))
	}
	return nil
}

func (machine *Machine) CreateSharedFolder(name string, hostPath string, writable bool, automount bool) error {
	cname := C.CString(name)
	chostPath := C.CString(hostPath)
	cwritable := C.PRBool(0)
	if writable {
		cwritable = C.PRBool(1)
	}
	cautomount := C.PRBool(0)
	if automount {
		cautomount = C.PRBool(1)
	}
	result := C.GoVboxMachineCreateSharedFolder(machine.cmachine, cname, chostPath, cwritable, cautomount)
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(chostPath))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to create shared folder on IMachine: %x", result))
	}
	return nil
}

func (machine *Machine) RemoveSharedFolder(name string) error {
	cname := C.CString(name)
	result := C.GoVboxMachineRemoveSharedFolder(machine.cmachine, cname)
	C.free(unsafe.Pointer(cname))

	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to remove shared folder on IMachine: %x", result))
	}
	return nil
}

func (machine *Machine) SetSettingsFilePath(path string) (Progress, error) {
	var progress Progress
	cpath := C.CString(path)
	result := C.GoVboxMachineSetSettingsFilePath(machine.cmachine, cpath, &progress.cprogress)
	C.free(unsafe.Pointer(cpath))

	if C.GoVboxFAILED(result) != 0 {
		return progress, errors.New(
			fmt.Sprintf("Failed to set settings file path for IMachine: %x", result))
	}
	return progress, nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (machine *Machine) Release() error {
	if machine.cmachine != nil {
		result := C.GoVboxIMachineRelease(machine.cmachine)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release IMachine: %x", result))
		}
		machine.cmachine = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (machine *Machine) Initialized() bool {
	return machine.cmachine != nil
}

// CreateMachine creates a VirtualBox machine.
// The machine must be registered by calling Register before it shows up in the
// GetMachines list.
// Flags is comma-separated. The most interesting flag is forceOverwrite=1.
// It returns the created machine and any error encountered.
func CreateMachine(
	settings string, name string, osTypeId string, flags string) (Machine, error) {
	var machine Machine
	if err := Init(); err != nil {
		return machine, err
	}

	csettings := C.CString(settings)
	cname := C.CString(name)
	cosTypeId := C.CString(osTypeId)
	cflags := C.CString(flags)
	result := C.GoVboxCreateMachine(cbox, csettings, cname, cosTypeId, cflags,
		&machine.cmachine)
	C.free(unsafe.Pointer(csettings))
	C.free(unsafe.Pointer(cname))
	C.free(unsafe.Pointer(cosTypeId))
	C.free(unsafe.Pointer(cflags))

	if C.GoVboxFAILED(result) != 0 || machine.cmachine == nil {
		return machine, errors.New(
			fmt.Sprintf("Failed to create IMachine: %x", result))
	}
	return machine, nil
}

// FindMachine returns the VirtualBox machine with the given name.
// It returns a new Machine instance and any error encountered.
func FindMachine(nameOrId string) (Machine, error) {
	var machine Machine
	if err := Init(); err != nil {
		return machine, err
	}

	cnameOrId := C.CString(nameOrId)
	result := C.GoVboxFindMachine(cbox, cnameOrId, &machine.cmachine)
	C.free(unsafe.Pointer(cnameOrId))

	if C.GoVboxFAILED(result) != 0 || machine.cmachine == nil {
		return machine, errors.New(
			fmt.Sprintf("Failed to find IMachine: %x", result))
	}
	return machine, nil
}

// GetMachines returns the machines known to VirtualBox.
// It returns a slice of Machine instances and any error encountered.
func GetMachines() ([]Machine, error) {
	if err := Init(); err != nil {
		return nil, err
	}

	var cmachinesPtr **C.IMachine
	var machineCount C.ULONG

	result := C.GoVboxGetMachines(cbox, &cmachinesPtr, &machineCount)
	if C.GoVboxFAILED(result) != 0 ||
		(cmachinesPtr == nil && machineCount != 0) {
		return nil, errors.New(
			fmt.Sprintf("Failed to get IMachine array: %x", result))
	}

	sliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cmachinesPtr)),
		Len:  int(machineCount),
		Cap:  int(machineCount),
	}
	cmachinesSlice := *(*[]*C.IMachine)(unsafe.Pointer(&sliceHeader))

	var machines = make([]Machine, machineCount)
	for i := range cmachinesSlice {
		machines[i] = Machine{cmachinesSlice[i]}
	}

	C.GoVboxArrayOutFree(unsafe.Pointer(cmachinesPtr))
	return machines, nil
}
