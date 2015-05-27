#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxDisplayGetScreenResolution(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32* cwidth, PRUint32* cheight,
    PRUint32* cbitsPerPixel, PRInt32* cxOrigin, PRInt32* cyOrigin) {
  return IDisplay_GetScreenResolution(cdisplay, cscreenId, cwidth, cheight,
      cbitsPerPixel, cxOrigin, cyOrigin);
}
HRESULT GoVboxDisplayTakeScreenShot(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint8* cdata, PRUint32 cwidth, PRUint32 cheight) {
  return IDisplay_TakeScreenShot(cdisplay, cscreenId, cdata, cwidth, cheight);
}
HRESULT GoVboxDisplayTakeScreenShotToArray(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32 cwidth, PRUint32 cheight,
    PRUint32* cdataSize, PRUint8** cdata) {
  return IDisplay_TakeScreenShotToArray(cdisplay, cscreenId, cwidth,
      cheight, cdataSize, cdata);
}
HRESULT GoVboxDisplayTakeScreenShotPNGToArray(IDisplay* cdisplay,
    PRUint32 cscreenId, PRUint32 cwidth, PRUint32 cheight,
    PRUint32* cdataSize, PRUint8** cdata) {
  return IDisplay_TakeScreenShotPNGToArray(cdisplay, cscreenId, cwidth,
      cheight, cdataSize, cdata);
}
HRESULT GoVboxIDisplayRelease(IDisplay* cdisplay) {
  return IDisplay_Release(cdisplay);
}
