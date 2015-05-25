#include "VBoxCAPIGlue.h"

// Wrapper declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);

HRESULT GoVboxGetMediumFormats(ISystemProperties* cprops,
    IMediumFormat*** cformats, ULONG* formatCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = ISystemProperties_GetMediumFormats(cprops,
      ComSafeArrayAsOutIfaceParam(safeArray, IMediumFormat *));
  g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
      (IUnknown ***)cformats, formatCount, safeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}

HRESULT GoVboxGetMediumFormatId(IMediumFormat* cformat, char** cid) {
  BSTR wid = NULL;
  HRESULT result = IMediumFormat_GetId(cformat, &wid);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wid, cid);
  g_pVBoxFuncs->pfnComUnallocString(wid);
  return result;
}
HRESULT GoVboxIMediumFormatRelease(IMediumFormat* cformat) {
  return IMediumFormat_Release(cformat);
}
