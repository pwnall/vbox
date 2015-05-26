#include "VBoxCAPIGlue.h"

// Wrapper declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);


HRESULT GoVboxCreateHardDisk(IVirtualBox* cbox, char* cformat, char* clocation,
    IMedium** cmedium) {
  BSTR wformat;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cformat, &wformat);
  if (FAILED(result))
    return result;

  BSTR wlocation;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(clocation, &wlocation);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnComUnallocString(wformat);
    return result;
  }

  result = IVirtualBox_CreateHardDisk(cbox, wformat, wlocation, cmedium);
  g_pVBoxFuncs->pfnComUnallocString(wlocation);
  g_pVBoxFuncs->pfnComUnallocString(wformat);

  return result;
}

HRESULT GoVboxMediumCreateBaseStorage(IMedium* cmedium, PRInt64 size,
    PRUint32 variantSize, PRUint32* cvariant, IProgress** cprogress) {
  return IMedium_CreateBaseStorage(cmedium, size, variantSize, cvariant,
      cprogress);
}
HRESULT GoVboxMediumDeleteStorage(IMedium* cmedium, IProgress** cprogress) {
  return IMedium_DeleteStorage(cmedium, cprogress);
}
HRESULT GoVboxGetMediumLocation(IMedium* cmedium, char** clocation) {
  BSTR wlocation = NULL;
  HRESULT result = IMedium_GetLocation(cmedium, &wlocation);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wlocation, clocation);
  g_pVBoxFuncs->pfnComUnallocString(wlocation);
  return result;
}
HRESULT GoVboxGetMediumState(IMedium* cmedium, PRUint32* cstate) {
  return IMedium_GetState(cmedium, cstate);
}
HRESULT GoVboxIMediumRelease(IMedium* cmedium) {
  return IMedium_Release(cmedium);
}
