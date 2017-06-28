#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetConsoleDisplay(IConsole* cconsole, IDisplay** cdisplay) {
  return IConsole_GetDisplay(cconsole, cdisplay);
}
HRESULT GoVboxGetConsoleKeyboard(IConsole* cconsole, IKeyboard** ckeyboard) {
  return IConsole_GetKeyboard(cconsole, ckeyboard);
}
HRESULT GoVboxGetConsoleMouse(IConsole* cconsole, IMouse** cmouse) {
  return IConsole_GetMouse(cconsole, cmouse);
}
HRESULT GoVboxGetConsoleMachine(IConsole* cconsole, IMachine** cmachine) {
  return IConsole_GetMachine(cconsole, cmachine);
}
HRESULT GoVboxGetConsoleEventSource(IConsole* cconsole,
    IEventSource** ceventSource) {
  return IConsole_GetEventSource(cconsole, ceventSource);
}
HRESULT GoVboxConsolePowerDown(IConsole* cconsole, IProgress** cprogress) {
  return IConsole_PowerDown(cconsole, cprogress);
}
HRESULT GoVboxIConsoleRelease(IConsole* cconsole) {
  return IConsole_Release(cconsole);
}
