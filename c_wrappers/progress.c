#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxProgressWaitForCompletion(IProgress* cprogress, int timeout) {
  return IProgress_WaitForCompletion(cprogress, timeout);
}
HRESULT GoVboxGetProgressPercent(IProgress* cprogress, PRUint32* cpercent) {
  return IProgress_GetPercent(cprogress, cpercent);
}
HRESULT GoVboxGetProgressResultCode(IProgress* cprogress, PRInt32* code) {
  return IProgress_GetResultCode(cprogress, code);
}
HRESULT GoVboxIProgressRelease(IProgress* cprogress) {
  return IProgress_Release(cprogress);
}
