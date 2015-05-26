package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include "VBoxCAPIGlue.h"
*/
import "C"  // cgo's virtual package

// Enumeration of AccessMode values
type AccessMode uint
const (
  // Open the image file in read-only mode.
  AccessMode_ReadOnly = C.AccessMode_ReadOnly
  // Open the image file in read-write mode.
  AccessMode_ReadWrite = C.AccessMode_ReadWrite
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
  CleanupMode_DetachAllReturnHardDisksOnly =
      C.CleanupMode_DetachAllReturnHardDisksOnly
  // Delete snapshots, detach media, return all media for closing.
  // This is not recommended by the API documentation, because users generally
  // want to keep their ISOs around.
  CleanupMode_Full = C.CleanupMode_Full
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

type LockType uint
const (
  // Shared lock that can be used to read the VM settings.
  LockType_Shared = C.LockType_Shared
  // Exclusive lock needed to change VM settings or start it.
  LockType_Write = C.LockType_Write
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

// Enumeration of StorageBus values
type StorageBus uint

const (
  // Null value that is never used by the API
  StorageBus_Null = C.StorageBus_Null
  StorageBus_Ide = C.StorageBus_IDE
  StorageBus_Sata = C.StorageBus_SATA
  StorageBus_Scsi = C.StorageBus_SCSI
  StorageBus_Floppy = C.StorageBus_Floppy
  StorageBus_Sas = C.StorageBus_SAS
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
