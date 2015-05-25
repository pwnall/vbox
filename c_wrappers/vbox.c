#include "VBoxCAPIGlue.h"

// NOTE: Including the C file is a sketchy but working method for getting it
//       compiled and linked with the Go source. The C must only be included in
//       one Go file. By convention, this is the file that wraps the
//       ClientInitialize() function.
#include "VBoxCAPIGlue.c"

HRESULT GoVboxFAILED(HRESULT result) {
  return FAILED(result);
}

HRESULT GoVboxCGlueInit() {
  return VBoxCGlueInit();
}
unsigned int GoVboxGetAppVersion() {
  return g_pVBoxFuncs->pfnGetVersion();
}
unsigned int GoVboxGetApiVersion() {
  return g_pVBoxFuncs->pfnGetAPIVersion();
}
HRESULT GoVboxClientInitialize(IVirtualBoxClient** client) {
  return g_pVBoxFuncs->pfnClientInitialize(NULL, client);
}
HRESULT GoVboxClientThreadInitialize() {
  return g_pVBoxFuncs->pfnClientThreadInitialize();
}
HRESULT GoVboxClientThreadUninitialize() {
  return g_pVBoxFuncs->pfnClientThreadUninitialize();
}
void GoVboxClientUninitialize() {
  g_pVBoxFuncs->pfnClientUninitialize();
}
HRESULT GoVboxGetVirtualBox(IVirtualBoxClient* client, IVirtualBox** cbox) {
  return IVirtualBoxClient_GetVirtualBox(client, cbox);
}
HRESULT GoVboxIVirtualBoxRelease(IVirtualBox* cbox) {
  return IVirtualBox_Release(cbox);
}

HRESULT GoVboxGetRevision(IVirtualBox* cbox, ULONG* revision) {
  return IVirtualBox_get_Revision(cbox, revision);
}
HRESULT GoVboxGetGuestOSTypes(IVirtualBox* cbox, IGuestOSType*** ctypes,
    ULONG* typeCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IVirtualBox_get_GuestOSTypes(cbox,
      ComSafeArrayAsOutIfaceParam(safeArray, IGuestOSType *));
  g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
      (IUnknown ***)ctypes, typeCount, safeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}

HRESULT GoVboxGetGuestOSTypeId(IGuestOSType* ctype, char** cid) {
  BSTR wid = NULL;
  HRESULT result = IGuestOSType_GetId(ctype, &wid);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wid, cid);
  g_pVBoxFuncs->pfnComUnallocString(wid);
  return result;
}
HRESULT GoVboxIGuestOSTypeRelease(IGuestOSType* ctype) {
  return IGuestOSType_Release(ctype);
}


HRESULT GoVboxGetSession(IVirtualBoxClient* client, ISession** csession) {
  return IVirtualBoxClient_GetSession(client, csession);
}
HRESULT GoVboxISessionRelease(ISession* csession) {
  return ISession_Release(csession);
}
