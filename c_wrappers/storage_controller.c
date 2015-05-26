#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetStorageControllerName(IStorageController* ccontroller,
    char** cname) {
  BSTR wname = NULL;
  HRESULT result = IStorageController_GetName(ccontroller, &wname);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wname, cname);
  g_pVBoxFuncs->pfnComUnallocString(wname);
  return result;
}
HRESULT GoVboxGetStorageControllerBus(IStorageController* ccontroller,
    PRUint32* cbus) {
  return IStorageController_GetBus(ccontroller, cbus);
}
HRESULT GoVboxGetStorageControllerType(IStorageController* ccontroller,
    PRUint32* ctype) {
  return IStorageController_GetControllerType(ccontroller, ctype);
}
HRESULT GoVboxSetStorageControllerType(IStorageController* ccontroller,
    PRUint32 ctype) {
  return IStorageController_SetControllerType(ccontroller, ctype);
}
HRESULT GoVboxIStorageControllerRelease(IStorageController* ccontroller) {
  return IStorageController_Release(ccontroller);
}

HRESULT GoVboxMachineAddStorageController(IMachine* cmachine, char* cname,
    PRUint32 connectionType, IStorageController** ccontroller) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_AddStorageController(cmachine, wname, connectionType,
      ccontroller);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
