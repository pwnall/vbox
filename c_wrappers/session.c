#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxUnlockMachine(ISession* csession) {
  return ISession_UnlockMachine(csession);
}
HRESULT GoVboxISessionRelease(ISession* csession) {
  return ISession_Release(csession);
}
HRESULT GoVboxGetSessionConsole(ISession* csession, IConsole** cconsole) {
  return ISession_get_Console(csession, cconsole);
}
HRESULT GoVboxGetSessionMachine(ISession* csession, IMachine** cmachine) {
  return ISession_get_Machine(csession, cmachine);
}
HRESULT GoVboxGetSessionState(ISession* csession, PRUint32* cstate) {
  return ISession_get_State(csession, cstate);
}
HRESULT GoVboxGetSessionType(ISession* csession, PRUint32* ctype) {
  return ISession_get_Type(csession, ctype);
}

HRESULT GoVboxGetSession(IVirtualBoxClient* client, ISession** csession) {
  return IVirtualBoxClient_get_Session(client, csession);
}

HRESULT GoVboxLockMachine(IMachine* cmachine, ISession* csession,
    PRUint32 clock) {
  return IMachine_LockMachine(cmachine, csession, clock);
}
