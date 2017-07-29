#include "glue.h"

HRESULT GoVboxIEventListenerRelease(IEventListener* ceventListener) {
  return IEventListener_Release(ceventListener);
}
