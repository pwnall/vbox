#include "glue.h"

HRESULT GoVboxEventSourceCreateListener(IEventSource* ceventSource,
    IEventListener** ceventListener) {
  return IEventSource_CreateListener(ceventSource, ceventListener);
}
HRESULT GoVboxEventSourceCreateAggregator(IEventSource* ceventSource,
    PRUint32 subordinatesCount, IEventSource** csubordinates,
    IEventSource** caggregator) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_UNKNOWN, 0, subordinatesCount);
  g_pVBoxFuncs->pfnSafeArrayCopyInParamHelper(pSafeArray, csubordinates,
      sizeof(IEventSource*) * subordinatesCount);
  HRESULT result = IEventSource_CreateAggregator(ceventSource,
    ComSafeArrayAsInParam(pSafeArray), caggregator);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxEventSourceRegisterListener(IEventSource* ceventSource,
    IEventListener* ceventListener, PRUint32 interestingCount,
    PRUint32* cinteresting, PRBool active) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_UI4, 0, interestingCount);
  g_pVBoxFuncs->pfnSafeArrayCopyInParamHelper(pSafeArray, cinteresting,
      sizeof(PRUint32) * interestingCount);
  HRESULT result = IEventSource_RegisterListener(ceventSource, ceventListener,
    ComSafeArrayAsInParam(pSafeArray), active);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxEventSourceUnregisterListener(IEventSource* ceventSource,
    IEventListener* ceventListener) {
  return IEventSource_UnregisterListener(ceventSource, ceventListener);
}
HRESULT GoVboxEventSourceFireEvent(IEventSource* ceventSource,
    IEvent* cevent, PRInt32 timeout, PRBool *fired) {
  return IEventSource_FireEvent(ceventSource, cevent, timeout, fired);
}
HRESULT GoVboxEventSourceGetEvent(IEventSource* ceventSource,
    IEventListener* ceventListener, PRInt32 timeout, IEvent **cevent) {
  return IEventSource_GetEvent(ceventSource, ceventListener, timeout, cevent);
}
HRESULT GoVboxEventSourceEventProcessed(IEventSource* ceventSource,
    IEventListener* ceventListener, IEvent* cevent) {
  return IEventSource_EventProcessed(ceventSource, ceventListener, cevent);
}
HRESULT GoVboxIEventSourceRelease(IEventSource* ceventSource) {
  return IEventSource_Release(ceventSource);
}
