#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxKeyboardPutScancodes(IKeyboard* ckeyboard,
    PRUint32 scancodesCount, PRInt32* cscancodes, PRUint32* ccodesStored) {
  return IKeyboard_PutScancodes(ckeyboard, scancodesCount, cscancodes,
      ccodesStored);
}
HRESULT GoVboxIKeyboardRelease(IKeyboard* ckeyboard) {
  return IKeyboard_Release(ckeyboard);
}
