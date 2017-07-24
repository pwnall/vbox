#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetConsoleDisplay(IConsole* cconsole, IDisplay** cdisplay) {
  return IConsole_get_Display(cconsole, cdisplay);
}
HRESULT GoVboxGetConsoleKeyboard(IConsole* cconsole, IKeyboard** ckeyboard) {
  return IConsole_get_Keyboard(cconsole, ckeyboard);
}
HRESULT GoVboxGetConsoleMouse(IConsole* cconsole, IMouse** cmouse) {
  return IConsole_get_Mouse(cconsole, cmouse);
}
HRESULT GoVboxGetConsoleMachine(IConsole* cconsole, IMachine** cmachine) {
  return IConsole_get_Machine(cconsole, cmachine);
}
HRESULT GoVboxGetConsoleEventSource(IConsole* cconsole,
    IEventSource** ceventSource) {
  return IConsole_get_EventSource(cconsole, ceventSource);
}
HRESULT GoVboxConsolePowerUp(IConsole* cconsole, IProgress** cprogress) {
  return IConsole_PowerUp(cconsole, cprogress);
}
HRESULT GoVboxConsolePowerDown(IConsole* cconsole, IProgress** cprogress) {
  return IConsole_PowerDown(cconsole, cprogress);
}
HRESULT GoVboxIConsoleRelease(IConsole* cconsole) {
  return IConsole_Release(cconsole);
}
