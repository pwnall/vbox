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
  return ISession_GetConsole(csession, cconsole);
}
HRESULT GoVboxGetSessionMachine(ISession* csession, IMachine** cmachine) {
  return ISession_GetMachine(csession, cmachine);
}
HRESULT GoVboxGetSessionState(ISession* csession, PRUint32* cstate) {
  return ISession_GetState(csession, cstate);
}
HRESULT GoVboxGetSessionType(ISession* csession, PRUint32* ctype) {
  return ISession_GetType(csession, ctype);
}

HRESULT GoVboxGetSession(IVirtualBoxClient* client, ISession** csession) {
  return IVirtualBoxClient_GetSession(client, csession);
}

HRESULT GoVboxLockMachine(IMachine* cmachine, ISession* csession,
    PRUint32 clock) {
  return IMachine_LockMachine(cmachine, csession, clock);
}


