#include "glue.h"

HRESULT GoVboxProgressWaitForCompletion(IProgress* cprogress, int timeout) {
  return IProgress_WaitForCompletion(cprogress, timeout);
}
HRESULT GoVboxGetProgressPercent(IProgress* cprogress, PRUint32* cpercent) {
  return IProgress_get_Percent(cprogress, cpercent);
}
HRESULT GoVboxGetProgressResultCode(IProgress* cprogress, PRInt32* code) {
  return IProgress_get_ResultCode(cprogress, code);
}
HRESULT GoVboxIProgressRelease(IProgress* cprogress) {
  return IProgress_Release(cprogress);
}
