#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxIEventRelease(IEvent* cevent) {
  return IEvent_Release(cevent);
}
HRESULT GoVboxEventGetType(IEvent* cevent, PRUint32 *ceventType) {
  return IEvent_GetType(cevent, ceventType);
}
HRESULT GoVboxEventGetSource(IEvent* cevent, IEventSource **ceventSource) {
  return IEvent_GetSource(cevent, ceventSource);
}
HRESULT GoVboxEventGetWaitable(IEvent* cevent, PRBool *cwaitable) {
  return IEvent_GetWaitable(cevent, cwaitable);
}
HRESULT GoVboxEventSetProcessed(IEvent* cevent) {
  return IEvent_SetProcessed(cevent);
}
HRESULT GoVboxEventWaitProcessed(IEvent* cevent, PRInt32 ctimeout, PRBool *cprocessed) {
  return IEvent_WaitProcessed(cevent, ctimeout, cprocessed);
}