#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxINetworkAdapterRelease(INetworkAdapter* cadapter) {
  return INetworkAdapter_Release(cadapter);
}
HRESULT GoVboxNetworkAdapterGetAdapterType(INetworkAdapter* cadapter, PRUint32 *cadapterType) {
  return INetworkAdapter_get_AdapterType(cadapter, cadapterType);
}
HRESULT GoVboxNetworkAdapterSetAdapterType(INetworkAdapter* cadapter, PRUint32 cadapterType) {
  return INetworkAdapter_put_AdapterType(cadapter, cadapterType);
}
HRESULT GoVboxNetworkAdapterGetSlot(INetworkAdapter* cadapter, PRUint32 *cslot) {
  return INetworkAdapter_get_Slot(cadapter, cslot);
}
HRESULT GoVboxNetworkAdapterGetEnabled(INetworkAdapter* cadapter, PRBool *cenabled) {
  return INetworkAdapter_get_Enabled(cadapter, cenabled);
}
HRESULT GoVboxNetworkAdapterSetEnabled(INetworkAdapter* cadapter, PRBool cenabled) {
  return INetworkAdapter_put_Enabled(cadapter, cenabled);
}
HRESULT GoVboxNetworkAdapterGetMACAddress(INetworkAdapter* cadapter,
    char** cmacAddress) {
  BSTR wmacAddress = NULL;
  HRESULT result = INetworkAdapter_get_MACAddress(cadapter, &wmacAddress);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wmacAddress, cmacAddress);
  g_pVBoxFuncs->pfnComUnallocString(wmacAddress);
  return result;
}
HRESULT GoVboxNetworkAdapterSetMACAddress(INetworkAdapter* cadapter,
    char* cmacAddress) {
  BSTR wmacAddress;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cmacAddress, &wmacAddress);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_MACAddress(cadapter, wmacAddress);
  g_pVBoxFuncs->pfnUtf16Free(wmacAddress);

  return result;
}
HRESULT GoVboxNetworkAdapterGetAttachmentType(INetworkAdapter* cadapter, PRUint32 *cattachmentType) {
  return INetworkAdapter_get_AttachmentType(cadapter, cattachmentType);
}
HRESULT GoVboxNetworkAdapterSetAttachmentType(INetworkAdapter* cadapter, PRUint32 cattachmentType) {
  return INetworkAdapter_put_AttachmentType(cadapter, cattachmentType);
}
HRESULT GoVboxNetworkAdapterGetBridgedInterface(INetworkAdapter* cadapter,
    char** cbridgedInterface) {
  BSTR wbridgedInterface = NULL;
  HRESULT result = INetworkAdapter_get_BridgedInterface(cadapter, &wbridgedInterface);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wbridgedInterface, cbridgedInterface);
  g_pVBoxFuncs->pfnComUnallocString(wbridgedInterface);
  return result;
}
HRESULT GoVboxNetworkAdapterSetBridgedInterface(INetworkAdapter* cadapter,
    char* cbridgedInterface) {
  BSTR wbridgedInterface;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cbridgedInterface, &wbridgedInterface);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_BridgedInterface(cadapter, wbridgedInterface);
  g_pVBoxFuncs->pfnUtf16Free(wbridgedInterface);

  return result;
}
HRESULT GoVboxNetworkAdapterGetHostOnlyInterface(INetworkAdapter* cadapter,
    char** chostOnlyInterface) {
  BSTR whostOnlyInterface = NULL;
  HRESULT result = INetworkAdapter_get_HostOnlyInterface(cadapter, &whostOnlyInterface);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(whostOnlyInterface, chostOnlyInterface);
  g_pVBoxFuncs->pfnComUnallocString(whostOnlyInterface);
  return result;
}
HRESULT GoVboxNetworkAdapterSetHostOnlyInterface(INetworkAdapter* cadapter,
    char* chostOnlyInterface) {
  BSTR whostOnlyInterface;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(chostOnlyInterface, &whostOnlyInterface);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_HostOnlyInterface(cadapter, whostOnlyInterface);
  g_pVBoxFuncs->pfnUtf16Free(whostOnlyInterface);

  return result;
}
HRESULT GoVboxNetworkAdapterGetInternalNetwork(INetworkAdapter* cadapter,
    char** cinternalNetwork) {
  BSTR winternalNetwork = NULL;
  HRESULT result = INetworkAdapter_get_InternalNetwork(cadapter, &winternalNetwork);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(winternalNetwork, cinternalNetwork);
  g_pVBoxFuncs->pfnComUnallocString(winternalNetwork);
  return result;
}
HRESULT GoVboxNetworkAdapterSetInternalNetwork(INetworkAdapter* cadapter,
    char* cinternalNetwork) {
  BSTR winternalNetwork;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cinternalNetwork, &winternalNetwork);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_InternalNetwork(cadapter, winternalNetwork);
  g_pVBoxFuncs->pfnUtf16Free(winternalNetwork);

  return result;
}
HRESULT GoVboxNetworkAdapterGetNATNetwork(INetworkAdapter* cadapter,
    char** cnatNetwork) {
  BSTR wnatNetwork = NULL;
  HRESULT result = INetworkAdapter_get_NATNetwork(cadapter, &wnatNetwork);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wnatNetwork, cnatNetwork);
  g_pVBoxFuncs->pfnComUnallocString(wnatNetwork);
  return result;
}
HRESULT GoVboxNetworkAdapterSetNATNetwork(INetworkAdapter* cadapter,
    char* cnatNetwork) {
  BSTR wnatNetwork;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cnatNetwork, &wnatNetwork);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_NATNetwork(cadapter, wnatNetwork);
  g_pVBoxFuncs->pfnUtf16Free(wnatNetwork);

  return result;
}
HRESULT GoVboxNetworkAdapterGetGenericDriver(INetworkAdapter* cadapter,
    char** cgenericDriver) {
  BSTR wgenericDriver = NULL;
  HRESULT result = INetworkAdapter_get_GenericDriver(cadapter, &wgenericDriver);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wgenericDriver, cgenericDriver);
  g_pVBoxFuncs->pfnComUnallocString(wgenericDriver);
  return result;
}
HRESULT GoVboxNetworkAdapterSetGenericDriver(INetworkAdapter* cadapter,
    char* cgenericDriver) {
  BSTR wgenericDriver;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cgenericDriver, &wgenericDriver);
  if (FAILED(result))
    return result;

  result = INetworkAdapter_put_GenericDriver(cadapter, wgenericDriver);
  g_pVBoxFuncs->pfnUtf16Free(wgenericDriver);

  return result;
}
HRESULT GoVboxNetworkAdapterGetCableConnected(INetworkAdapter* cadapter, PRBool *ccableConnected) {
  return INetworkAdapter_get_CableConnected(cadapter, ccableConnected);
}
HRESULT GoVboxNetworkAdapterSetCableConnected(INetworkAdapter* cadapter, PRBool ccableConnected) {
  return INetworkAdapter_put_CableConnected(cadapter, ccableConnected);
}
