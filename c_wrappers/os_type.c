#include "glue.h"

HRESULT GoVboxGetGuestOSTypes(IVirtualBox* cbox, IGuestOSType*** ctypes,
    ULONG* typeCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IVirtualBox_get_GuestOSTypes(cbox,
      ComSafeArrayAsOutIfaceParam(safeArray, IGuestOSType *));
  if (!FAILED(result)) {
    result = g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
        (IUnknown ***)ctypes, typeCount, safeArray);
  }
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}

HRESULT GoVboxGetGuestOSTypeId(IGuestOSType* ctype, char** cid) {
  BSTR wid = NULL;
  HRESULT result = IGuestOSType_get_Id(ctype, &wid);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wid, cid);
  g_pVBoxFuncs->pfnComUnallocString(wid);
  return result;
}
HRESULT GoVboxIGuestOSTypeRelease(IGuestOSType* ctype) {
  return IGuestOSType_Release(ctype);
}
