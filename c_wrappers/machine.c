#include "VBoxCAPIGlue.h"

// Wrapper declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);


HRESULT GoVboxGetMachineName(IMachine* cmachine, char** cname) {
  BSTR wname = NULL;
  HRESULT result = IMachine_GetName(cmachine, &wname);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wname, cname);
  g_pVBoxFuncs->pfnComUnallocString(wname);
  return result;
}
HRESULT GoVboxGetMachineOSTypeId(IMachine* cmachine, char** cosTypeId) {
  BSTR wosTypeId = NULL;
  HRESULT result = IMachine_GetOSTypeId(cmachine, &wosTypeId);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wosTypeId, cosTypeId);
  g_pVBoxFuncs->pfnComUnallocString(wosTypeId);
  return result;

}
HRESULT GoVboxGetMachineSettingsFilePath(IMachine* cmachine, char** cpath) {
  BSTR wpath = NULL;
  HRESULT result = IMachine_GetSettingsFilePath(cmachine, &wpath);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wpath, cpath);
  g_pVBoxFuncs->pfnComUnallocString(wpath);
  return result;
}
HRESULT GoVboxGetMachineSettingsModified(IMachine* cmachine,
    PRBool* cmodified) {
  return IMachine_GetSettingsModified(cmachine, cmodified);
}
HRESULT GoVboxMachineSaveSettings(IMachine* cmachine) {
  return IMachine_SaveSettings(cmachine);
}
HRESULT GoVboxIMachineRelease(IMachine* cmachine) {
  return IMachine_Release(cmachine);
}

HRESULT GoVboxCreateMachine(IVirtualBox* cbox, char* cname, char* cosTypeId,
    char* cflags, IMachine** cmachine) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  BSTR wosTypeId;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cosTypeId, &wosTypeId);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnComUnallocString(wname);
    return result;
  }

  BSTR wflags = NULL;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cflags, &wflags);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnComUnallocString(wosTypeId);
    g_pVBoxFuncs->pfnComUnallocString(wname);
  }

  result = IVirtualBox_CreateMachine(cbox, NULL, wname, 0, NULL, wosTypeId,
      wflags, cmachine);
  g_pVBoxFuncs->pfnComUnallocString(wflags);
  g_pVBoxFuncs->pfnComUnallocString(wosTypeId);
  g_pVBoxFuncs->pfnComUnallocString(wname);

  return result;
}
HRESULT GoVboxGetMachines(IVirtualBox* cbox, IMachine*** cmachines,
    ULONG* machineCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IVirtualBox_GetMachines(cbox,
      ComSafeArrayAsOutIfaceParam(safeArray, IMachine *));
  g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
      (IUnknown ***)cmachines, machineCount, safeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}
