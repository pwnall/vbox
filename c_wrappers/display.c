#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxDisplayGetScreenResolution(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32* cwidth, PRUint32* cheight,
    PRUint32* cbitsPerPixel, PRInt32* cxOrigin, PRInt32* cyOrigin) {
  return IDisplay_GetScreenResolution(cdisplay, cscreenId, cwidth, cheight,
      cbitsPerPixel, cxOrigin, cyOrigin, NULL);
}
HRESULT GoVboxDisplayTakeScreenShot(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint8* cdata, PRUint32 cwidth, PRUint32 cheight) {
  return IDisplay_TakeScreenShot(cdisplay, cscreenId, cdata, cwidth, cheight,
      BitmapFormat_RGBA);
}
HRESULT GoVboxDisplayTakeScreenShotToArray(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32 cwidth, PRUint32 cheight,
    PRUint32* cdataSize, PRUint8** cdata) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IDisplay_TakeScreenShotToArray(cdisplay, cscreenId, cwidth,
      cheight, BitmapFormat_RGBA,
      ComSafeArrayAsOutTypeParam(pSafeArray, PRUint8));
  if  (!FAILED(result)) {
    g_pVBoxFuncs->pfnSafeArrayCopyOutParamHelper((void **)cdata, cdataSize,
        VT_UI1, pSafeArray);
  }
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxDisplayTakeScreenShotPNGToArray(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32 cwidth, PRUint32 cheight,
    PRUint32* cdataSize, PRUint8** cdata) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IDisplay_TakeScreenShotToArray(cdisplay, cscreenId,
      cwidth, cheight, BitmapFormat_PNG,
      ComSafeArrayAsOutTypeParam(pSafeArray, PRUint8));
  if  (!FAILED(result)) {
    g_pVBoxFuncs->pfnSafeArrayCopyOutParamHelper((void **)cdata, cdataSize,
        VT_UI1, pSafeArray);
  }
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxIDisplayRelease(IDisplay* cdisplay) {
  return IDisplay_Release(cdisplay);
}
