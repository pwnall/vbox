#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetConsoleMachine(IConsole* cconsole, IMachine** cmachine) {
  return IConsole_GetMachine(cconsole, cmachine);
}
HRESULT GoVboxConsolePowerDown(IConsole* cconsole, IProgress** cprogress) {
  return IConsole_PowerDown(cconsole, cprogress);
}
HRESULT GoVboxIConsoleRelease(IConsole* cconsole) {
  return IConsole_Release(cconsole);
}
