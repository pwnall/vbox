#include "glue.h"

HRESULT GoVboxCreateHardDisk(IVirtualBox* cbox, char* cformat, char* clocation,
    IMedium** cmedium) {
  BSTR wformat;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cformat, &wformat);
  if (FAILED(result))
    return result;

  BSTR wlocation;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(clocation, &wlocation);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wformat);
    return result;
  }

  result = IVirtualBox_CreateMedium(cbox, wformat, wlocation,
      AccessMode_ReadOnly, DeviceType_HardDisk, cmedium);
  g_pVBoxFuncs->pfnUtf16Free(wlocation);
  g_pVBoxFuncs->pfnUtf16Free(wformat);

  return result;
}
HRESULT GoVboxOpenMedium(IVirtualBox* cbox, char* clocation,
    PRUint32 cdeviceType, PRUint32 caccessType, PRBool cforceNewUuid,
    IMedium** cmedium) {
  BSTR wlocation;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(clocation, &wlocation);
  if (FAILED(result))
    return result;

  result = IVirtualBox_OpenMedium(cbox, wlocation, cdeviceType, caccessType,
      cforceNewUuid, cmedium);
  g_pVBoxFuncs->pfnUtf16Free(wlocation);

  return result;
}

HRESULT GoVboxMediumCreateBaseStorage(IMedium* cmedium, PRInt64 size,
    PRUint32 variantCount, PRUint32* cvariant, IProgress** cprogress) {

  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_UI4, 0, variantCount);
  g_pVBoxFuncs->pfnSafeArrayCopyInParamHelper(pSafeArray, cvariant,
      sizeof(PRUint32) * variantCount);
  HRESULT result = IMedium_CreateBaseStorage(cmedium, size,
      ComSafeArrayAsInParam(pSafeArray), cprogress);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxMediumDeleteStorage(IMedium* cmedium, IProgress** cprogress) {
  return IMedium_DeleteStorage(cmedium, cprogress);
}
HRESULT GoVboxMediumClose(IMedium* cmedium) {
  return IMedium_Close(cmedium);
}
HRESULT GoVboxGetMediumLocation(IMedium* cmedium, char** clocation) {
  BSTR wlocation = NULL;
  HRESULT result = IMedium_get_Location(cmedium, &wlocation);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wlocation, clocation);
  g_pVBoxFuncs->pfnComUnallocString(wlocation);
  return result;
}
HRESULT GoVboxGetMediumState(IMedium* cmedium, PRUint32* cstate) {
  return IMedium_get_State(cmedium, cstate);
}
HRESULT GoVboxGetMediumSize(IMedium* cmedium, PRInt64* csize) {
  return IMedium_get_Size(cmedium, csize);
}
HRESULT GoVboxIMediumRelease(IMedium* cmedium) {
  return IMedium_Release(cmedium);
}
