#include "glue.h"
#include <stdio.h>

HRESULT GoVboxGetMachineName(IMachine* cmachine, char** cname) {
  BSTR wname = NULL;
  HRESULT result = IMachine_get_Name(cmachine, &wname);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wname, cname);
  g_pVBoxFuncs->pfnComUnallocString(wname);
  return result;
}
HRESULT GoVboxGetMachineOSTypeId(IMachine* cmachine, char** cosTypeId) {
  BSTR wosTypeId = NULL;
  HRESULT result = IMachine_get_OSTypeId(cmachine, &wosTypeId);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wosTypeId, cosTypeId);
  g_pVBoxFuncs->pfnComUnallocString(wosTypeId);
  return result;

}
HRESULT GoVboxGetMachineSettingsFilePath(IMachine* cmachine, char** cpath) {
  BSTR wpath = NULL;
  HRESULT result = IMachine_get_SettingsFilePath(cmachine, &wpath);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wpath, cpath);
  g_pVBoxFuncs->pfnComUnallocString(wpath);
  return result;
}
HRESULT GoVboxGetMachineMemorySize(IMachine* cmachine, PRUint32* cram) {
  return IMachine_get_MemorySize(cmachine, cram);
}
HRESULT GoVboxSetMachineMemorySize(IMachine* cmachine, PRUint32 cram) {
  return IMachine_put_MemorySize(cmachine, cram);
}
HRESULT GoVboxGetMachineVRAMSize(IMachine* cmachine, PRUint32* cvram) {
  return IMachine_get_VRAMSize(cmachine, cvram);
}
HRESULT GoVboxSetMachineVRAMSize(IMachine* cmachine, PRUint32 cvram) {
  return IMachine_put_VRAMSize(cmachine, cvram);
}
HRESULT GoVboxGetMachineCPUCount(IMachine* cmachine, PRUint32* ccpus) {
  return IMachine_get_CPUCount(cmachine, ccpus);
}
HRESULT GoVboxSetMachineCPUCount(IMachine* cmachine, PRUint32 ccpus) {
  return IMachine_put_CPUCount(cmachine, ccpus);
}
HRESULT GoVboxGetClipboardMode(IMachine* cmachine, PRUint32* cmode) {
  return IMachine_get_ClipboardMode(cmachine, cmode);
}
HRESULT GoVboxSetClipboardMode(IMachine* cmachine, PRUint32 cmode) {
  return IMachine_put_ClipboardMode(cmachine, cmode);
}
HRESULT GoVboxGetDnDMode(IMachine* cmachine, PRUint32* cmode) {
  return IMachine_get_DnDMode(cmachine, cmode);
}
HRESULT GoVboxSetDnDMode(IMachine* cmachine, PRUint32 cmode) {
  return IMachine_put_DnDMode(cmachine, cmode);
}
HRESULT GoVboxGetMachineState(IMachine* cmachine, PRUint32* cstate) {
  return IMachine_get_State(cmachine, cstate);
}
HRESULT GoVboxGetMachinePointingHIDType(IMachine* cmachine, PRUint32* ctype) {
  return IMachine_get_PointingHIDType(cmachine, ctype);
}
HRESULT GoVboxSetMachinePointingHIDType(IMachine* cmachine, PRUint32 ctype) {
  return IMachine_put_PointingHIDType(cmachine, ctype);
}
HRESULT GoVboxGetMachineSettingsModified(IMachine* cmachine,
    PRBool* cmodified) {
  return IMachine_get_SettingsModified(cmachine, cmodified);
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
HRESULT GoVboxMachineDetachDevice(IMachine* cmachine, char* cname, PRInt32
    cport, PRInt32 cdevice) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_DetachDevice(cmachine, wname, cport, cdevice);
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
HRESULT GoVboxMachineGetNetworkAdapter(IMachine* cmachine, PRInt32 cdevice,
    INetworkAdapter** cadapter) {
  return IMachine_GetNetworkAdapter(cmachine, cdevice, cadapter);
}
HRESULT GoVboxMachineGetAudioAdapter(IMachine* cmachine,
    IAudioAdapter** cadapter) {
  return IMachine_get_AudioAdapter(cmachine, cadapter);
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
HRESULT GoVboxMachineSetExtraData(IMachine* cmachine, char* ckey, char *cvalue) {
  BSTR wkey;
  BSTR wvalue;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cvalue, &wvalue);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wkey);
    return result;
  }

  result = IMachine_SetExtraData(cmachine, wkey, wvalue);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  g_pVBoxFuncs->pfnUtf16Free(wvalue);

  return result;
}
HRESULT GoVboxMachineGetGuestProperty(IMachine* cmachine, char* ckey, char** cvalue, PRInt64* ctimestamp, char** cflags) {
  BSTR wkey;
  BSTR wvalue = NULL;
  BSTR wflags = NULL;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = IMachine_GetGuestProperty(cmachine, wkey, &wvalue, ctimestamp, &wflags);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  if (FAILED(result)) {
    return result;
  }

  g_pVBoxFuncs->pfnUtf16ToUtf8(wvalue, cvalue);
  g_pVBoxFuncs->pfnComUnallocString(wvalue);

  g_pVBoxFuncs->pfnUtf16ToUtf8(wflags, cflags);
  g_pVBoxFuncs->pfnComUnallocString(wflags);

  return result;
}
HRESULT GoVboxMachineGetGuestPropertyValue(IMachine* cmachine, char* ckey, char** cvalue) {
  BSTR wkey;
  BSTR wvalue = NULL;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = IMachine_GetGuestPropertyValue(cmachine, wkey, &wvalue);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  if (FAILED(result)) {
    return result;
  }

  g_pVBoxFuncs->pfnUtf16ToUtf8(wvalue, cvalue);
  g_pVBoxFuncs->pfnComUnallocString(wvalue);

  return result;
}
HRESULT GoVboxMachineSetGuestProperty(IMachine* cmachine, char* ckey, char *cvalue, char *cflags) {
  BSTR wkey;
  BSTR wvalue;
  BSTR wflags;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cvalue, &wvalue);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wkey);
    return result;
  }

  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cflags, &wflags);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wkey);
    g_pVBoxFuncs->pfnUtf16Free(wvalue);
    return result;
  }

  result = IMachine_SetGuestProperty(cmachine, wkey, wvalue, wflags);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  g_pVBoxFuncs->pfnUtf16Free(wvalue);
  g_pVBoxFuncs->pfnUtf16Free(wflags);

  return result;
}
HRESULT GoVboxMachineSetGuestPropertyValue(IMachine* cmachine, char* ckey, char *cvalue) {
  BSTR wkey;
  BSTR wvalue;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cvalue, &wvalue);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wkey);
    return result;
  }

  result = IMachine_SetGuestPropertyValue(cmachine, wkey, wvalue);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  g_pVBoxFuncs->pfnUtf16Free(wvalue);

  return result;
}
HRESULT GoVboxMachineDeleteGuestProperty(IMachine* cmachine, char* ckey) {
  BSTR wkey;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(ckey, &wkey);
  if (FAILED(result))
    return result;

  result = IMachine_DeleteGuestProperty(cmachine, wkey);
  g_pVBoxFuncs->pfnUtf16Free(wkey);
  return result;
}
char** convertUtf16StringArray(SAFEARRAY* sa, ULONG* count) {
  BSTR *values = NULL;
  HRESULT result = g_pVBoxFuncs->pfnSafeArrayCopyOutParamHelper(
      (void **)&values, count, VT_BSTR, sa);
  *count /= sizeof(values[0]);
  char **ar = malloc(*count * sizeof(char*));
  for (ULONG i = 0; i < *count; ++i) {
    g_pVBoxFuncs->pfnUtf16ToUtf8(values[i], &ar[i]);
    g_pVBoxFuncs->pfnComUnallocString(values[i]);
  }

  g_pVBoxFuncs->pfnArrayOutFree(values);
  return ar;
}
HRESULT GoVboxMachineEnumerateGuestProperties(IMachine* cmachine, char* cpattern,
    char*** cnames, ULONG* nameCount,
    char*** cvalues, ULONG* valueCount,
    LONG64** ctimestamps, ULONG* timestampCount,
    char*** cflags, ULONG* flagCount) {
  SAFEARRAY *namesSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  SAFEARRAY *timestampsSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  SAFEARRAY *valuesSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();
  SAFEARRAY *flagsSafeArray = g_pVBoxFuncs->pfnSafeArrayOutParamAlloc();

  HRESULT result = IMachine_EnumerateGuestProperties(cmachine,
    NULL,
    ComSafeArrayAsOutTypeParam(namesSafeArray, BSTR),
    ComSafeArrayAsOutTypeParam(valuesSafeArray, BSTR),
    ComSafeArrayAsOutTypeParam(timestampsSafeArray, LONG64),
    ComSafeArrayAsOutTypeParam(flagsSafeArray, BSTR));

  if (SUCCEEDED(result)) {
    *cnames = convertUtf16StringArray(namesSafeArray, nameCount);
    *cvalues = convertUtf16StringArray(valuesSafeArray, valueCount);
    *cflags = convertUtf16StringArray(flagsSafeArray, flagCount);
    result = g_pVBoxFuncs->pfnSafeArrayCopyOutParamHelper(
        (void **)ctimestamps, timestampCount, VT_I8, timestampsSafeArray);
  }

  g_pVBoxFuncs->pfnSafeArrayDestroy(namesSafeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(timestampsSafeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(valuesSafeArray);
  g_pVBoxFuncs->pfnSafeArrayDestroy(flagsSafeArray);

  return result;
}

HRESULT GoVboxGetMachineAccelerate2DVideoEnabled(IMachine* cmachine,
    PRBool* cenabled) {
  return IMachine_get_Accelerate2DVideoEnabled(cmachine, cenabled);
}
HRESULT GoVboxSetMachineAccelerate2DVideoEnabled(IMachine* cmachine,
    PRBool cenabled) {
  return IMachine_put_Accelerate2DVideoEnabled(cmachine, cenabled);
}
HRESULT GoVboxGetMachineAccelerate3DEnabled(IMachine* cmachine,
    PRBool* cenabled) {
  return IMachine_get_Accelerate3DEnabled(cmachine, cenabled);
}
HRESULT GoVboxSetMachineAccelerate3DEnabled(IMachine* cmachine,
    PRBool cenabled) {
  return IMachine_put_Accelerate3DEnabled(cmachine, cenabled);
}
HRESULT GoVboxMachineCreateSharedFolder(IMachine* cmachine, char* cname,
    char *chostPath, PRBool writable, PRBool automount) {
  BSTR wname;
  BSTR whostPath;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = g_pVBoxFuncs->pfnUtf8ToUtf16(chostPath, &whostPath);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wname);
    return result;
  }

  result = IMachine_CreateSharedFolder(cmachine, wname, whostPath, writable,
      automount);
  g_pVBoxFuncs->pfnUtf16Free(wname);
  g_pVBoxFuncs->pfnUtf16Free(whostPath);

  return result;
}
HRESULT GoVboxMachineRemoveSharedFolder(IMachine* cmachine, char* cname) {
  BSTR wname;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  result = IMachine_RemoveSharedFolder(cmachine, wname);
  g_pVBoxFuncs->pfnUtf16Free(wname);

  return result;
}
HRESULT GoVboxMachineSetSettingsFilePath(IMachine* cmachine, char* cpath,
    IProgress** cprogress) {
  BSTR wpath;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cpath, &wpath);
  if (FAILED(result))
    return result;

  result = IMachine_SetSettingsFilePath(cmachine, wpath, cprogress);
  g_pVBoxFuncs->pfnUtf16Free(wpath);

  return result;
}
HRESULT GoVboxIMachineRelease(IMachine* cmachine) {
  return IMachine_Release(cmachine);
}

HRESULT GoVboxCreateMachine(IVirtualBox* cbox, char *settingsFile, char* cname,
    char* cosTypeId, char* cflags, IMachine** cmachine) {
  BSTR wsettings = NULL;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wsettings);
  if (FAILED(result))
    return result;

  BSTR wname;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cname, &wname);
  if (FAILED(result))
    return result;

  BSTR wosTypeId;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cosTypeId, &wosTypeId);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wsettings);
    g_pVBoxFuncs->pfnUtf16Free(wname);
    return result;
  }

  BSTR wflags = NULL;
  result = g_pVBoxFuncs->pfnUtf8ToUtf16(cflags, &wflags);
  if (FAILED(result)) {
    g_pVBoxFuncs->pfnUtf16Free(wsettings);
    g_pVBoxFuncs->pfnUtf16Free(wosTypeId);
    g_pVBoxFuncs->pfnUtf16Free(wname);
  }

  SAFEARRAY *pSafeArray = g_pVBoxFuncs->pfnSafeArrayCreateVector(
      VT_BSTR, 0, 0);
  result = IVirtualBox_CreateMachine(cbox, wsettings, wname,
      ComSafeArrayAsInParam(pSafeArray), wosTypeId, wflags, cmachine);
  g_pVBoxFuncs->pfnSafeArrayDestroy(pSafeArray);
  g_pVBoxFuncs->pfnUtf16Free(wflags);
  g_pVBoxFuncs->pfnUtf16Free(wosTypeId);
  g_pVBoxFuncs->pfnUtf16Free(wname);
  g_pVBoxFuncs->pfnUtf16Free(wsettings);

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
  HRESULT result = IVirtualBox_get_Machines(cbox,
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
