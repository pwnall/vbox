package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo !windows LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/network_adapter.c"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
) // cgo's virtual package

// The description of a VirtualBox machine
type NetworkAdapter struct {
	cadapter *C.INetworkAdapter
}

/*
	TODO: add missing methods

	nsresult (*GetLineSpeed)(INetworkAdapter *pThis, PRUint32 *lineSpeed);
	nsresult (*SetLineSpeed)(INetworkAdapter *pThis, PRUint32 lineSpeed);

	nsresult (*GetPromiscModePolicy)(INetworkAdapter *pThis, PRUint32 *promiscModePolicy);
	nsresult (*SetPromiscModePolicy)(INetworkAdapter *pThis, PRUint32 promiscModePolicy);

	nsresult (*GetTraceEnabled)(INetworkAdapter *pThis, PRBool *traceEnabled);
	nsresult (*SetTraceEnabled)(INetworkAdapter *pThis, PRBool traceEnabled);

	nsresult (*GetTraceFile)(INetworkAdapter *pThis, PRUnichar * *traceFile);
	nsresult (*SetTraceFile)(INetworkAdapter *pThis, PRUnichar * traceFile);

	nsresult (*GetNATEngine)(INetworkAdapter *pThis, INATEngine * *NATEngine);

	nsresult (*GetBootPriority)(INetworkAdapter *pThis, PRUint32 *bootPriority);
	nsresult (*SetBootPriority)(INetworkAdapter *pThis, PRUint32 bootPriority);

	nsresult (*GetBandwidthGroup)(INetworkAdapter *pThis, IBandwidthGroup * *bandwidthGroup);
	nsresult (*SetBandwidthGroup)(INetworkAdapter *pThis, IBandwidthGroup * bandwidthGroup);

	nsresult (*GetProperty)(
			INetworkAdapter *pThis,
			PRUnichar * key,
			PRUnichar * * value
	);

	nsresult (*SetProperty)(
			INetworkAdapter *pThis,
			PRUnichar * key,
			PRUnichar * value
	);

	nsresult (*GetProperties)(
			INetworkAdapter *pThis,
			PRUnichar * names,
			PRUint32 *returnNamesSize,
			PRUnichar *** returnNames,
			PRUint32 *returnValuesSize,
			PRUnichar *** returnValues
	);
*/

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (adapter *NetworkAdapter) Release() error {
	if adapter.cadapter != nil {
		result := C.GoVboxINetworkAdapterRelease(adapter.cadapter)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release INetworkAdapter: %x", result))
		}
		adapter.cadapter = nil
	}
	return nil
}

func (adapter *NetworkAdapter) GetAdapterType() (NetworkAdapterType, error) {
	var ctype C.PRUint32

	result := C.GoVboxNetworkAdapterGetAdapterType(adapter.cadapter, &ctype)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter type: %x", result))
	}
	return NetworkAdapterType(ctype), nil
}

func (adapter *NetworkAdapter) SetAdapterType(adapterType NetworkAdapterType) error {
	result := C.GoVboxNetworkAdapterSetAdapterType(adapter.cadapter,
		C.PRUint32(adapterType))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter type: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetSlot() (uint, error) {
	var cslot C.PRUint32

	result := C.GoVboxNetworkAdapterGetSlot(adapter.cadapter, &cslot)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter slot: %x", result))
	}
	return uint(cslot), nil
}

func (adapter *NetworkAdapter) GetEnabled() (bool, error) {
	var cenabled C.PRBool

	result := C.GoVboxNetworkAdapterGetEnabled(adapter.cadapter, &cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter state: %x", result))
	}
	return cenabled != 0, nil
}

func (adapter *NetworkAdapter) SetEnabled(enabled bool) error {
	cenabled := C.PRBool(0)
	if enabled {
		cenabled = C.PRBool(1)
	}

	result := C.GoVboxNetworkAdapterSetEnabled(adapter.cadapter, cenabled)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter state: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetMACAddress() (string, error) {
	var cmacAddress *C.char
	result := C.GoVboxNetworkAdapterGetMACAddress(adapter.cadapter, &cmacAddress)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter MAC address: %x", result))
	}
	macAddress := C.GoString(cmacAddress)
	C.GoVboxUtf8Free(cmacAddress)
	return macAddress, nil
}

func (adapter *NetworkAdapter) SetMACAddress(macAddress string) error {
	cmacAddress := C.CString(macAddress)
	result := C.GoVboxNetworkAdapterSetMACAddress(adapter.cadapter, cmacAddress)
	C.free(unsafe.Pointer(cmacAddress))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter MAC address: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetAttachmentType() (NetworkAttachmentType, error) {
	var ctype C.PRUint32

	result := C.GoVboxNetworkAdapterGetAttachmentType(adapter.cadapter, &ctype)
	if C.GoVboxFAILED(result) != 0 {
		return 0, errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter attacment type: %x", result))
	}
	return NetworkAttachmentType(ctype), nil
}

func (adapter *NetworkAdapter) SetAttachmentType(attachmentType NetworkAttachmentType) error {
	result := C.GoVboxNetworkAdapterSetAttachmentType(adapter.cadapter,
		C.PRUint32(attachmentType))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter attachment type: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetBridgeInterface() (string, error) {
	var cbridgedInterface *C.char
	result := C.GoVboxNetworkAdapterGetBridgedInterface(adapter.cadapter, &cbridgedInterface)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter bridged interface: %x", result))
	}
	bridgedInterface := C.GoString(cbridgedInterface)
	C.GoVboxUtf8Free(cbridgedInterface)
	return bridgedInterface, nil
}

func (adapter *NetworkAdapter) SetBridgedInterface(bridgedInterface string) error {
	cbridgedInterface := C.CString(bridgedInterface)
	result := C.GoVboxNetworkAdapterSetBridgedInterface(adapter.cadapter, cbridgedInterface)
	C.free(unsafe.Pointer(cbridgedInterface))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter bridged interface: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetHostOnlyInterface() (string, error) {
	var chostOnlyInterface *C.char
	result := C.GoVboxNetworkAdapterGetHostOnlyInterface(adapter.cadapter, &chostOnlyInterface)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter host only interface: %x", result))
	}
	hostOnlyInterface := C.GoString(chostOnlyInterface)
	C.GoVboxUtf8Free(chostOnlyInterface)
	return hostOnlyInterface, nil
}

func (adapter *NetworkAdapter) SetHostOnlyInterface(hostOnlyInterface string) error {
	chostOnlyInterface := C.CString(hostOnlyInterface)
	result := C.GoVboxNetworkAdapterSetHostOnlyInterface(adapter.cadapter, chostOnlyInterface)
	C.free(unsafe.Pointer(chostOnlyInterface))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter host only interface: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetInternalNetwork() (string, error) {
	var cinternalNetwork *C.char
	result := C.GoVboxNetworkAdapterGetInternalNetwork(adapter.cadapter, &cinternalNetwork)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter internal network: %x", result))
	}
	internalNetwork := C.GoString(cinternalNetwork)
	C.GoVboxUtf8Free(cinternalNetwork)
	return internalNetwork, nil
}

func (adapter *NetworkAdapter) SetInternalNetwork(internalNetwork string) error {
	cinternalNetwork := C.CString(internalNetwork)
	result := C.GoVboxNetworkAdapterSetInternalNetwork(adapter.cadapter, cinternalNetwork)
	C.free(unsafe.Pointer(cinternalNetwork))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter internal network: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetNATNetwork() (string, error) {
	var cnatNetwork *C.char
	result := C.GoVboxNetworkAdapterGetNATNetwork(adapter.cadapter, &cnatNetwork)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter NAT network: %x", result))
	}
	natNetwork := C.GoString(cnatNetwork)
	C.GoVboxUtf8Free(cnatNetwork)
	return natNetwork, nil
}

func (adapter *NetworkAdapter) SetNATNetwork(natNetwork string) error {
	cnatNetwork := C.CString(natNetwork)
	result := C.GoVboxNetworkAdapterSetNATNetwork(adapter.cadapter, cnatNetwork)
	C.free(unsafe.Pointer(cnatNetwork))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter NAT network: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetGenericDriver() (string, error) {
	var cgenericDriver *C.char
	result := C.GoVboxNetworkAdapterGetGenericDriver(adapter.cadapter, &cgenericDriver)
	if C.GoVboxFAILED(result) != 0 {
		return "", errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter MAC address: %x", result))
	}
	genericDriver := C.GoString(cgenericDriver)
	C.GoVboxUtf8Free(cgenericDriver)
	return genericDriver, nil
}

func (adapter *NetworkAdapter) SetGenericDriver(genericDriver string) error {
	cgenericDriver := C.CString(genericDriver)
	result := C.GoVboxNetworkAdapterSetGenericDriver(adapter.cadapter, cgenericDriver)
	C.free(unsafe.Pointer(cgenericDriver))
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter MAC address: %x", result))
	}
	return nil
}

func (adapter *NetworkAdapter) GetCableConnected() (bool, error) {
	var ccableConnected C.PRBool

	result := C.GoVboxNetworkAdapterGetCableConnected(adapter.cadapter, &ccableConnected)
	if C.GoVboxFAILED(result) != 0 {
		return false, errors.New(
			fmt.Sprintf("Failed to get NetworkAdapter cable state: %x", result))
	}
	return ccableConnected != 0, nil
}

func (adapter *NetworkAdapter) SetCableConnected(cableConnected bool) error {
	ccableConnected := C.PRBool(0)
	if cableConnected {
		ccableConnected = C.PRBool(1)
	}

	result := C.GoVboxNetworkAdapterSetCableConnected(adapter.cadapter, ccableConnected)
	if C.GoVboxFAILED(result) != 0 {
		return errors.New(
			fmt.Sprintf("Failed to set NetworkAdapter cable state: %x", result))
	}
	return nil
}
