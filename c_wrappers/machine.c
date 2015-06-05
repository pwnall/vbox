#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


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
HRESULT GoVboxGetMachineMemorySize(IMachine* cmachine, PRUint32* cram) {
  return IMachine_GetMemorySize(cmachine, cram);
}
HRESULT GoVboxSetMachineMemorySize(IMachine* cmachine, PRUint32 cram) {
  return IMachine_SetMemorySize(cmachine, cram);
}
HRESULT GoVboxGetMachineVRAMSize(IMachine* cmachine, PRUint32* cvram) {
  return IMachine_GetVRAMSize(cmachine, cvram);
}
HRESULT GoVboxSetMachineVRAMSize(IMachine* cmachine, PRUint32 cvram) {
  return IMachine_SetVRAMSize(cmachine, cvram);
}
HRESULT GoVboxGetMachinePointingHIDType(IMachine* cmachine, PRUint32* ctype) {
  return IMachine_GetPointingHIDType(cmachine, ctype);
}
HRESULT GoVboxSetMachinePointingHIDType(IMachine* cmachine, PRUint32 ctype) {
  return IMachine_SetPointingHIDType(cmachine, ctype);
}
HRESULT GoVboxGetMachineSettingsModified(IMachine* cmachine,
    PRBool* cmodified) {
  return IMachine_GetSettingsModified(cmachine, cmodified);
}
HRESULT GoVboxMachineSaveSettings(IMachine* cmachine) {
  return IMachine_SaveSettings(cmachine);
}
HRESULT GoVboxMachineUnregister(IMachine* cmachine, PRUint32 cleanupMode,
    IMedium*** cmedia, ULONG* mediaCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IMachine_Unregister(cmachine, cleanupMode,
      ComSafeArrayAsOutIfaceParam(safeArray, IMedium *));
  if (!FAILED(result)) {
    result = g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
        (IUnknown ***)cmedia, mediaCount, safeArray);
  }
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}
HRESULT GoVboxMachineDeleteConfig(IMachine* cmachine, PRUint32 mediaCount,
    IMedium** cmedia, IProgress** cprogress) {
  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_UNKNOWN, 0, mediaCount);
  g_pVBoxFuncs->pfnSafeArrayCopyInParamHelper(pSafeArray, cmedia,
      sizeof(IMedium*) * mediaCount);
  HRESULT result = IMachine_DeleteConfig(cmachine,
      ComSafeArrayAsInParam(pSafeArray), cprogress);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  return result;
}
HRESULT GoVboxMachineAttachDevice(IMachine* cmachine, char* cname, PRInt32
    cport, PRInt32 cdevice, PRUint32 ctype, IMedium* cmedium) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_AttachDevice(cmachine, wname, cport, cdevice, ctype,
      cmedium);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
HRESULT GoVboxMachineUnmountMedium(IMachine* cmachine, char* cname, PRInt32
    cport, PRInt32 cdevice, PRBool cforce) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_UnmountMedium(cmachine, wname, cport, cdevice, cforce);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
HRESULT GoVboxMachineGetMedium(IMachine* cmachine, char* cname, PRInt32
    cport, PRInt32 cdevice, IMedium** cmedium) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_GetMedium(cmachine, wname, cport, cdevice, cmedium);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
HRESULT GoVboxMachineLaunchVMProcess(IMachine* cmachine, ISession* csession,
    char* cuiType, char* cenvironment, IProgress** cprogress) {
  BSTR wuiType;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cuiType, &wuiType);
  if (FAILED(result))
    return result;

  BSTR wenvironment;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cenvironment, &wenvironment);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wuiType);
    return result;
  }

  result = IMachine_LaunchVMProcess(cmachine, csession, wuiType, wenvironment,
      cprogress);
  g_pVBoxFuncs->pfnUtf16Free(wenvironment);
  g_pVBoxFuncs->pfnUtf16Free(wuiType);

  return result;
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
    g_pVBoxFuncs->pfnUtf16Free(wname);
    return result;
  }

  BSTR wflags = NULL;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cflags, &wflags);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wosTypeId);
    g_pVBoxFuncs->pfnUtf16Free(wname);
  }

  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_BSTR, 0, 0);
  result = IVirtualBox_CreateMachine(cbox, NULL, wname,
      ComSafeArrayAsInParam(pSafeArray), wosTypeId, wflags, cmachine);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  g_pVBoxFuncs->pfnUtf16Free(wflags);
  g_pVBoxFuncs->pfnUtf16Free(wosTypeId);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
HRESULT GoVboxFindMachine(IVirtualBox* cbox, char* cnameOrId,
    IMachine** cmachine) {
  BSTR wnameOrId;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cnameOrId, &wnameOrId);
  if (FAILED(result))
    return result;

  result = IVirtualBox_FindMachine(cbox, wnameOrId, cmachine);
  g_pVBoxFuncs->pfnUtf16Free(wnameOrId);

  return result;
}
HRESULT GoVboxGetMachines(IVirtualBox* cbox, IMachine*** cmachines,
    ULONG* machineCount) {
  SAFEARRAY *safeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  HRESULT result = IVirtualBox_GetMachines(cbox,
      ComSafeArrayAsOutIfaceParam(safeArray, IMachine *));
  if (!FAILED(result)) {
    result = g_pVBoxFuncs->pfnSafeArrayCopyOutIfaceParamHelper(
        (IUnknown ***)cmachines, machineCount, safeArray);
  }
  g_pVBoxFuncs->pfnSafeArrayDestroy(safeArray);
  return result;
}
HRESULT GoVboxRegisterMachine(IVirtualBox* cbox, IMachine* cmachine) {
  return IVirtualBox_RegisterMachine(cbox, cmachine);
}
