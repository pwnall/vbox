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

