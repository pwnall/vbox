#include "VBoxCAPIGlue.h"

// NOTE: Including the C file is a sketchy but working method for getting it
//       compiled and linked with the Go source. The C must only be included in
//       one Go file. By convention, this is the file that wraps the
//       ClientInitialize() function.
#include "VBoxCAPIGlue.c"

HRESULT GoVboxFAILED(HRESULT result) {
  return FAILED(result);
}
HRESULT GoVboxArrayOutFree(void* carray) {
  return g_pVBoxFuncs->pfnArrayOutFree(carray);
}
void GoVboxUtf8Free(char* cstring) {
  g_pVBoxFuncs->pfnUtf8Free(cstring);
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

HRESULT GoVboxGetRevision(IVirtualBox* cbox, ULONG* revision) {
  return IVirtualBox_GetRevision(cbox, revision);
}
HRESULT GoVboxIVirtualBoxRelease(IVirtualBox* cbox) {
  return IVirtualBox_Release(cbox);
}
HRESULT GoVboxComposeMachineFilename(IVirtualBox* cbox, char* cname,
    char* cflags, char* cbaseFolder, char **cpath) {

  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  BSTR wflags = NULL;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cflags, &wflags);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wname);
  }

  BSTR wbaseFolder;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cbaseFolder, &wbaseFolder);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wflags);
    g_pVBoxFuncs->pfnUtf16Free(wname);
    return result;
  }

  BSTR wpath = NULL;
  result = IVirtualBox_ComposeMachineFilename(cbox, wname, NULL, wflags,
      wbaseFolder, &wpath);
  g_pVBoxFuncs->pfnUtf16Free(wbaseFolder);
  g_pVBoxFuncs->pfnUtf16Free(wflags);
  g_pVBoxFuncs->pfnUtf16Free(wname);
  if (FAILED(result))
    return result;


  g_pVBoxFuncs->pfnUtf16ToUtf8(wpath, cpath);
  g_pVBoxFuncs->pfnComUnallocString(wpath);
  return result;
}

HRESULT GoVboxGetSession(IVirtualBoxClient* client, ISession** csession) {
  return IVirtualBoxClient_GetSession(client, csession);
}
HRESULT GoVboxISessionRelease(ISession* csession) {
  return ISession_Release(csession);
}
