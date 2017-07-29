#include "glue.h"

HRESULT GoVboxIEventRelease(IEvent* cevent) {
  return IEvent_Release(cevent);
}
HRESULT GoVboxEventGetType(IEvent* cevent, PRUint32 *ceventType) {
  return IEvent_get_Type(cevent, ceventType);
}
HRESULT GoVboxEventGetSource(IEvent* cevent, IEventSource **ceventSource) {
  return IEvent_get_Source(cevent, ceventSource);
}
HRESULT GoVboxEventGetWaitable(IEvent* cevent, PRBool *cwaitable) {
  return IEvent_get_Waitable(cevent, cwaitable);
}
HRESULT GoVboxEventSetProcessed(IEvent* cevent) {
  return IEvent_SetProcessed(cevent);
}
HRESULT GoVboxEventWaitProcessed(IEvent* cevent, PRInt32 ctimeout, PRBool *cprocessed) {
  return IEvent_WaitProcessed(cevent, ctimeout, cprocessed);
}

#ifndef WIN32
const nsIID IID_IGuestPropertyChangedEvent = IGUESTPROPERTYCHANGEDEVENT_IID;
#endif
HRESULT GoVboxEventGetGuestPropertyChangedEvent(IEvent *cevent, IGuestPropertyChangedEvent **cguestPropEvent) {
  return IEvent_QueryInterface(cevent, &IID_IGuestPropertyChangedEvent, (void **)cguestPropEvent);
}
HRESULT GoVboxGuestPropertyChangedEventGetName(IGuestPropertyChangedEvent* cevent, char** cname) {
  BSTR wname = NULL;
  HRESULT result = IGuestPropertyChangedEvent_get_Name(cevent, &wname);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wname, cname);
  g_pVBoxFuncs->pfnComUnallocString(wname);
  return result;
}
HRESULT GoVboxGuestPropertyChangedEventGetValue(IGuestPropertyChangedEvent* cevent, char** cvalue) {
  BSTR wvalue = NULL;
  HRESULT result = IGuestPropertyChangedEvent_get_Value(cevent, &wvalue);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wvalue, cvalue);
  g_pVBoxFuncs->pfnComUnallocString(wvalue);
  return result;
}
HRESULT GoVboxGuestPropertyChangedEventGetFlags(IGuestPropertyChangedEvent* cevent, char** cflags) {
  BSTR wflags = NULL;
  HRESULT result = IGuestPropertyChangedEvent_get_Flags(cevent, &wflags);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wflags, cflags);
  g_pVBoxFuncs->pfnComUnallocString(wflags);
  return result;
}
