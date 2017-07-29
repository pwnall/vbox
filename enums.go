package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include "VBoxCAPIGlue.h"
*/
import "C" // cgo's virtual package

// Enumeration of AccessMode values
type AccessMode uint

const (
	// Open the image file in read-only mode.
	AccessMode_ReadOnly = C.AccessMode_ReadOnly
	// Open the image file in read-write mode.
	AccessMode_ReadWrite = C.AccessMode_ReadWrite
)

// Enumeration of APICMode values
type APICMode uint

const (
	APICMode_Disabled = C.APICMode_Disabled
	APICMode_APIC     = C.APICMode_APIC
	APICMode_X2APIC   = C.APICMode_X2APIC
)

// Enumeration of BIOSBootMenuMode values
type BootMenuMode uint

const (
	BootMenuMode_Disabled       = C.BIOSBootMenuMode_Disabled
	BootMenuMode_MenuOnly       = C.BIOSBootMenuMode_MenuOnly
	BootMenuMode_MessageAndMenu = C.BIOSBootMenuMode_MessageAndMenu
)

// Enumeration of CleanupMode values
type CleanupMode uint

const (
	// Unregister the machine, do not detach media or delete snapshots.
	CleanupMode_UnregisterOnly = C.CleanupMode_UnregisterOnly
	// Delete snapshots, detach media, do not return any media for closing.
	CleanupMode_DetachAllReturnNone = C.CleanupMode_DetachAllReturnNone
	// Delete snapshots, detach media, return hard disks for closing.
	// This does not return removable media attached to the VM.
	CleanupMode_DetachAllReturnHardDisksOnly = C.CleanupMode_DetachAllReturnHardDisksOnly
	// Delete snapshots, detach media, return all media for closing.
	// This is not recommended by the API documentation, because users generally
	// want to keep their ISOs around.
	CleanupMode_Full = C.CleanupMode_Full
)

// Enumeration of ClipboardMode values

type ClipboardMode uint

const (
	ClipboardMode_Disabled      = C.ClipboardMode_Disabled
	ClipboardMode_HostToGuest   = C.ClipboardMode_HostToGuest
	ClipboardMode_GuestToHost   = C.ClipboardMode_GuestToHost
	ClipboardMode_Bidirectional = C.ClipboardMode_Bidirectional
)

// Enumeration of DeviceType values
type DeviceType uint

const (
	// No device
	DeviceType_Null = C.DeviceType_Null
	// Floppy device
	DeviceType_Floppy = C.DeviceType_Floppy
	// DVD/CD-ROM device
	DeviceType_Dvd = C.DeviceType_DVD
	// Hard disk device
	DeviceType_HardDisk = C.DeviceType_HardDisk
	// Hard disk device
	DeviceType_Network = C.DeviceType_Network
	// Hard disk device
	DeviceType_Usb = C.DeviceType_USB
	// Hard disk device
	DeviceType_SharedFolder = C.DeviceType_SharedFolder
)

// Enumeration of DnDMode values

type DnDMode uint

const (
	DnDMode_Disabled      = C.DnDMode_Disabled
	DnDMode_HostToGuest   = C.DnDMode_HostToGuest
	DnDMode_GuestToHost   = C.DnDMode_GuestToHost
	DnDMode_Bidirectional = C.DnDMode_Bidirectional
)

// Enumeration of VBoxEventType values
type EventType uint

const (
	EventType_Invalid                                 = C.VBoxEventType_Invalid
	EventType_Any                                     = C.VBoxEventType_Any
	EventType_Vetoable                                = C.VBoxEventType_Vetoable
	EventType_MachineEvent                            = C.VBoxEventType_MachineEvent
	EventType_SnapshotEvent                           = C.VBoxEventType_SnapshotEvent
	EventType_InputEvent                              = C.VBoxEventType_InputEvent
	EventType_LastWildcard                            = C.VBoxEventType_LastWildcard
	EventType_OnMachineStateChanged                   = C.VBoxEventType_OnMachineStateChanged
	EventType_OnMachineDataChanged                    = C.VBoxEventType_OnMachineDataChanged
	EventType_OnExtraDataChanged                      = C.VBoxEventType_OnExtraDataChanged
	EventType_OnExtraDataCanChange                    = C.VBoxEventType_OnExtraDataCanChange
	EventType_OnMediumRegistered                      = C.VBoxEventType_OnMediumRegistered
	EventType_OnMachineRegistered                     = C.VBoxEventType_OnMachineRegistered
	EventType_OnSessionStateChanged                   = C.VBoxEventType_OnSessionStateChanged
	EventType_OnSnapshotTaken                         = C.VBoxEventType_OnSnapshotTaken
	EventType_OnSnapshotDeleted                       = C.VBoxEventType_OnSnapshotDeleted
	EventType_OnSnapshotChanged                       = C.VBoxEventType_OnSnapshotChanged
	EventType_OnGuestPropertyChanged                  = C.VBoxEventType_OnGuestPropertyChanged
	EventType_OnMousePointerShapeChanged              = C.VBoxEventType_OnMousePointerShapeChanged
	EventType_OnMouseCapabilityChanged                = C.VBoxEventType_OnMouseCapabilityChanged
	EventType_OnKeyboardLedsChanged                   = C.VBoxEventType_OnKeyboardLedsChanged
	EventType_OnStateChanged                          = C.VBoxEventType_OnStateChanged
	EventType_OnAdditionsStateChanged                 = C.VBoxEventType_OnAdditionsStateChanged
	EventType_OnNetworkAdapterChanged                 = C.VBoxEventType_OnNetworkAdapterChanged
	EventType_OnSerialPortChanged                     = C.VBoxEventType_OnSerialPortChanged
	EventType_OnParallelPortChanged                   = C.VBoxEventType_OnParallelPortChanged
	EventType_OnStorageControllerChanged              = C.VBoxEventType_OnStorageControllerChanged
	EventType_OnMediumChanged                         = C.VBoxEventType_OnMediumChanged
	EventType_OnVRDEServerChanged                     = C.VBoxEventType_OnVRDEServerChanged
	EventType_OnUSBControllerChanged                  = C.VBoxEventType_OnUSBControllerChanged
	EventType_OnUSBDeviceStateChanged                 = C.VBoxEventType_OnUSBDeviceStateChanged
	EventType_OnSharedFolderChanged                   = C.VBoxEventType_OnSharedFolderChanged
	EventType_OnRuntimeError                          = C.VBoxEventType_OnRuntimeError
	EventType_OnCanShowWindow                         = C.VBoxEventType_OnCanShowWindow
	EventType_OnShowWindow                            = C.VBoxEventType_OnShowWindow
	EventType_OnCPUChanged                            = C.VBoxEventType_OnCPUChanged
	EventType_OnVRDEServerInfoChanged                 = C.VBoxEventType_OnVRDEServerInfoChanged
	EventType_OnEventSourceChanged                    = C.VBoxEventType_OnEventSourceChanged
	EventType_OnCPUExecutionCapChanged                = C.VBoxEventType_OnCPUExecutionCapChanged
	EventType_OnGuestKeyboard                         = C.VBoxEventType_OnGuestKeyboard
	EventType_OnGuestMouse                            = C.VBoxEventType_OnGuestMouse
	EventType_OnNATRedirect                           = C.VBoxEventType_OnNATRedirect
	EventType_OnHostPCIDevicePlug                     = C.VBoxEventType_OnHostPCIDevicePlug
	EventType_OnVBoxSVCAvailabilityChanged            = C.VBoxEventType_OnVBoxSVCAvailabilityChanged
	EventType_OnBandwidthGroupChanged                 = C.VBoxEventType_OnBandwidthGroupChanged
	EventType_OnGuestMonitorChanged                   = C.VBoxEventType_OnGuestMonitorChanged
	EventType_OnStorageDeviceChanged                  = C.VBoxEventType_OnStorageDeviceChanged
	EventType_OnClipboardModeChanged                  = C.VBoxEventType_OnClipboardModeChanged
	EventType_OnDnDModeChanged                        = C.VBoxEventType_OnDnDModeChanged
	EventType_OnNATNetworkChanged                     = C.VBoxEventType_OnNATNetworkChanged
	EventType_OnNATNetworkStartStop                   = C.VBoxEventType_OnNATNetworkStartStop
	EventType_OnNATNetworkAlter                       = C.VBoxEventType_OnNATNetworkAlter
	EventType_OnNATNetworkCreationDeletion            = C.VBoxEventType_OnNATNetworkCreationDeletion
	EventType_OnNATNetworkSetting                     = C.VBoxEventType_OnNATNetworkSetting
	EventType_OnNATNetworkPortForward                 = C.VBoxEventType_OnNATNetworkPortForward
	EventType_OnGuestSessionStateChanged              = C.VBoxEventType_OnGuestSessionStateChanged
	EventType_OnGuestSessionRegistered                = C.VBoxEventType_OnGuestSessionRegistered
	EventType_OnGuestProcessRegistered                = C.VBoxEventType_OnGuestProcessRegistered
	EventType_OnGuestProcessStateChanged              = C.VBoxEventType_OnGuestProcessStateChanged
	EventType_OnGuestProcessInputNotify               = C.VBoxEventType_OnGuestProcessInputNotify
	EventType_OnGuestProcessOutput                    = C.VBoxEventType_OnGuestProcessOutput
	EventType_OnGuestFileRegistered                   = C.VBoxEventType_OnGuestFileRegistered
	EventType_OnGuestFileStateChanged                 = C.VBoxEventType_OnGuestFileStateChanged
	EventType_OnGuestFileOffsetChanged                = C.VBoxEventType_OnGuestFileOffsetChanged
	EventType_OnGuestFileRead                         = C.VBoxEventType_OnGuestFileRead
	EventType_OnGuestFileWrite                        = C.VBoxEventType_OnGuestFileWrite
	EventType_OnVideoCaptureChanged                   = C.VBoxEventType_OnVideoCaptureChanged
	EventType_OnGuestUserStateChanged                 = C.VBoxEventType_OnGuestUserStateChanged
	EventType_OnGuestMultiTouch                       = C.VBoxEventType_OnGuestMultiTouch
	EventType_OnHostNameResolutionConfigurationChange = C.VBoxEventType_OnHostNameResolutionConfigurationChange
	EventType_OnSnapshotRestored                      = C.VBoxEventType_OnSnapshotRestored
	EventType_OnMediumConfigChanged                   = C.VBoxEventType_OnMediumConfigChanged
	EventType_Last                                    = C.VBoxEventType_Last
)

type LockType uint

const (
	// Shared lock that can be used to read the VM settings.
	LockType_Shared = C.LockType_Shared
	// Exclusive lock needed to change VM settings or start it.
	LockType_Write = C.LockType_Write
)

// Enumeration of MediumState values
type MachineState uint

const (
	MachineState_Null                   = C.MachineState_Null
	MachineState_PoweredOff             = C.MachineState_PoweredOff
	MachineState_Saved                  = C.MachineState_Saved
	MachineState_Teleported             = C.MachineState_Teleported
	MachineState_Aborted                = C.MachineState_Aborted
	MachineState_Running                = C.MachineState_Running
	MachineState_Paused                 = C.MachineState_Paused
	MachineState_Stuck                  = C.MachineState_Stuck
	MachineState_Teleporting            = C.MachineState_Teleporting
	MachineState_LiveSnapshotting       = C.MachineState_LiveSnapshotting
	MachineState_Starting               = C.MachineState_Starting
	MachineState_Stopping               = C.MachineState_Stopping
	MachineState_Saving                 = C.MachineState_Saving
	MachineState_Restoring              = C.MachineState_Restoring
	MachineState_TeleportingPausedVM    = C.MachineState_TeleportingPausedVM
	MachineState_TeleportingIn          = C.MachineState_TeleportingIn
	MachineState_FaultTolerantSyncing   = C.MachineState_FaultTolerantSyncing
	MachineState_DeletingSnapshotOnline = C.MachineState_DeletingSnapshotOnline
	MachineState_DeletingSnapshotPaused = C.MachineState_DeletingSnapshotPaused
	MachineState_OnlineSnapshotting     = C.MachineState_OnlineSnapshotting
	MachineState_RestoringSnapshot      = C.MachineState_RestoringSnapshot
	MachineState_DeletingSnapshot       = C.MachineState_DeletingSnapshot
	MachineState_SettingUp              = C.MachineState_SettingUp
	MachineState_Snapshotting           = C.MachineState_Snapshotting
	MachineState_FirstOnline            = C.MachineState_FirstOnline
	MachineState_LastOnline             = C.MachineState_LastOnline
	MachineState_FirstTransient         = C.MachineState_FirstTransient
	MachineState_LastTransient          = C.MachineState_LastTransient
)

// Enumeration of MediumState values
type MediumState uint

const (
	// The medium's backing image was not created or was deleted.
	MediumState_NotCreated = C.MediumState_NotCreated
	// The medium's backing image was not created or was deleted.
	MediumState_Created = C.MediumState_Created
	// The medium's backing image is locked with a shared reader lock.
	MediumState_LockedRead = C.MediumState_LockedRead
	// The medium's backing image is locked with an exclusive writer lock.
	MediumState_LockedWrite = C.MediumState_LockedWrite
	// The medium's backing image cannot / was not accessed.
	MediumState_Inaccessible = C.MediumState_Inaccessible
	// The medium's backing image is being built.
	MediumState_Creating = C.MediumState_Creating
	// The medium's backing image is being deleted.
	MediumState_Deleting = C.MediumState_Deleting
)

// Enumeration of MediumVariant values
type MediumVariant uint

const (
	// Default image options.
	MediumVariant_Standard MediumVariant = C.MediumVariant_Standard
	// Entire image is allocated at creation time.
	MediumVariant_Fixed MediumVariant = C.MediumVariant_Fixed
	// The image's directory is not created.
	MediumVariant_NoCreateDir MediumVariant = C.MediumVariant_NoCreateDir
)

// Enumeration of MouseButtonState values
type MouseButtonState uint

const (
	// No button is pressed.
	MouseButtonState_None         = 0
	MouseButtonState_LeftButton   = C.MouseButtonState_LeftButton
	MouseButtonState_RightButton  = C.MouseButtonState_RightButton
	MouseButtonState_MiddleButton = C.MouseButtonState_MiddleButton
	MouseButtonState_WheelUp      = C.MouseButtonState_WheelUp
	MouseButtonState_WheelDown    = C.MouseButtonState_WheelDown
	MouseButtonState_XButton1     = C.MouseButtonState_XButton1
	MouseButtonState_XButton2     = C.MouseButtonState_XButton2
)

// Enumeration of NetworkAdapterType values
type NetworkAdapterType uint

const (
	NetworkAdapterType_Null      = C.NetworkAdapterType_Null
	NetworkAdapterType_Am79C970A = C.NetworkAdapterType_Am79C970A
	NetworkAdapterType_Am79C973  = C.NetworkAdapterType_Am79C973
	NetworkAdapterType_I82540EM  = C.NetworkAdapterType_I82540EM
	NetworkAdapterType_I82543GC  = C.NetworkAdapterType_I82543GC
	NetworkAdapterType_I82545EM  = C.NetworkAdapterType_I82545EM
	NetworkAdapterType_Virtio    = C.NetworkAdapterType_Virtio
)

// Enumeration of NetworkAttachmentType values
type NetworkAttachmentType uint

const (
	NetworkAttachmentType_Null       = C.NetworkAttachmentType_Null
	NetworkAttachmentType_NAT        = C.NetworkAttachmentType_NAT
	NetworkAttachmentType_Bridged    = C.NetworkAttachmentType_Bridged
	NetworkAttachmentType_Internal   = C.NetworkAttachmentType_Internal
	NetworkAttachmentType_HostOnly   = C.NetworkAttachmentType_HostOnly
	NetworkAttachmentType_Generic    = C.NetworkAttachmentType_Generic
	NetworkAttachmentType_NATNetwork = C.NetworkAttachmentType_NATNetwork
)

// Enumeration of PointingHIDType values
type PointingHidType uint

const (
	// No mouse
	PointingHidType_None = C.PointingHIDType_None
	// PS/2 mouse
	PointingHidType_Ps2Mouse = C.PointingHIDType_PS2Mouse
	// USB mouse (relative pointer)
	PointingHidType_UsbMouse = C.PointingHIDType_USBMouse
	// USB tablet (absolute pointer)
	PointingHidType_UsbTablet = C.PointingHIDType_USBTablet
	// Combo PS2/2 or USB mouse, depending on guest (negative perf implications)
	PointingHidType_ComboMouse = C.PointingHIDType_ComboMouse
	// USB multi-touch device
	// This also adds USB tablet and mouse devices.
	PointingHidType_UsbMultiTouch = C.PointingHIDType_USBMultiTouch
)

// Enumeration of StorageBus values
type StorageBus uint

const (
	// Null value that is never used by the API
	StorageBus_Null   = C.StorageBus_Null
	StorageBus_Ide    = C.StorageBus_IDE
	StorageBus_Sata   = C.StorageBus_SATA
	StorageBus_Scsi   = C.StorageBus_SCSI
	StorageBus_Floppy = C.StorageBus_Floppy
	StorageBus_Sas    = C.StorageBus_SAS
)

// Enumeration of SessionState values
type SessionState uint

const (
	// Null value that is never used by the API
	SessionState_Null = C.SessionState_Null
	// The session / machine is not locked.
	SessionState_Unlocked = C.SessionState_Unlocked
	// The session / machine is locked.
	SessionState_Locked = C.SessionState_Locked
	// Transient state while a VM is locked and started.
	SessionState_Spawning = C.SessionState_Spawning
	// The session is getting unlocked.
	SessionState_Unlocking = C.SessionState_Unlocking
)

// Enumeration of SessionType values
type SessionType uint

const (
	// Null value that is never used by the API
	SessionType_Null = C.SessionType_Null
	// The session has an exclusive lock on a VM.
	SessionType_WriteLock = C.SessionType_WriteLock
	// The session has launched a VM process.
	SessionType_Remote = C.SessionType_Remote
	// The session has a shared lock on a VM.
	SessionType_Shared = C.SessionType_Shared
)

// Enumeration of StorageControllerType values
type StorageControllerType uint

const (
	// Null value that is never used by the API
	StorageControllerType_Null = C.StorageControllerType_Null
	// SCSI controller of the LsiLogic variant
	StorageControllerType_LsiLogic = C.StorageControllerType_LsiLogic
	// SCSI controller of the BusLogic variant
	StorageControllerType_BusLogic = C.StorageControllerType_BusLogic
	// The only SATA controller available
	StorageControllerType_IntelAhci = C.StorageControllerType_IntelAhci
	// IDE controller of the PIIX3 variant
	StorageControllerType_Piix3 = C.StorageControllerType_PIIX3
	// IDE controller of the PIIX4 variant
	StorageControllerType_Piix4 = C.StorageControllerType_PIIX4
	// IDE controller of the ICH6 variant
	StorageControllerType_Ich6 = C.StorageControllerType_ICH6
	// The only floppy drive controller available
	StorageControllerType_I82078 = C.StorageControllerType_I82078
	// LsiLogic SCSI controller that uses SAS
	StorageControllerType_LsiLogicSas = C.StorageControllerType_LsiLogicSas
)

// Enumeration of UsbControllerType values
type UsbControllerType uint

const (
	// Null value that is never used by the API
	UsbControllerType_Null = C.USBControllerType_Null
	// USB 1.1 controller available in the free VirtualBox edition
	UsbControllerType_Ohci = C.USBControllerType_OHCI
	// USB 2.0 controller that requires the extension pack
	UsbControllerType_Ehci = C.USBControllerType_EHCI
)
