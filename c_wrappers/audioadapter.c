#include "glue.h"

HRESULT GoVboxIAudioAdapterRelease(IAudioAdapter* caudioadapter) {
  return IAudioAdapter_Release(caudioadapter);
}
