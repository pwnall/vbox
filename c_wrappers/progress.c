#include "VBoxCAPIGlue.h"

// Wrapper declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);


HRESULT GoVboxProgressWaitForCompletion(IProgress* cprogress, int timeout) {
  return IProgress_WaitForCompletion(cprogress, timeout);
}
HRESULT GoVboxIProgressRelease(IProgress* cprogress) {
  return IProgress_Release(cprogress);
}
