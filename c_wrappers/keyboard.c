#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxKeyboardPutScancodes(IKeyboard* ckeyboard,
    PRUint32 scancodesCount, PRInt32* cscancodes, PRUint32* ccodesStored) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_I4, 0, scancodesCount);
  g_pVBoxFuncs->pfnSafeArrayCopyInParamHelper(pSafeArray, cscancodes,
      sizeof(PRInt32) * scancodesCount);
  HRESULT result = IKeyboard_PutScancodes(ckeyboard,
      ComSafeArrayAsInParam(pSafeArray), ccodesStored);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxIKeyboardRelease(IKeyboard* ckeyboard) {
  return IKeyboard_Release(ckeyboard);
}
