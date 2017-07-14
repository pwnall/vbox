#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetBiosSettingsLogoImagePath(IBIOSSettings* csettings,
    char** clogoImagePath) {
  BSTR wlogoImagePath = NULL;
  HRESULT result = IBIOSSettings_GetLogoImagePath(csettings, &wlogoImagePath);
  if (FAILED(result))
    return result;

  g_pVBoxFuncs->pfnUtf16ToUtf8(wlogoImagePath, clogoImagePath);
  g_pVBoxFuncs->pfnComUnallocString(wlogoImagePath);
  return result;
}
HRESULT GoVboxSetBiosSettingsLogoImagePath(IBIOSSettings* csettings,
    char* clogoImagePath) {
  BSTR wlogoImagePath;
  HRESULT result = g_pVBoxFuncs->pfnUtf8ToUtf16(clogoImagePath, &wlogoImagePath);
  if (FAILED(result))
    return result;

  result = IBIOSSettings_SetLogoImagePath(csettings, wlogoImagePath);
  g_pVBoxFuncs->pfnUtf16Free(wlogoImagePath);

  return result;
}
HRESULT GoVboxGetBiosSettingsLogoFadeIn(IBIOSSettings* csettings,
    PRBool* clogoFadeIn) {
  return IBIOSSettings_GetLogoFadeIn(csettings, clogoFadeIn);
}
HRESULT GoVboxSetBiosSettingsLogoFadeIn(IBIOSSettings* csettings,
    PRBool clogoFadeIn) {
  return IBIOSSettings_SetLogoFadeIn(csettings, clogoFadeIn);
}
HRESULT GoVboxGetBiosSettingsLogoFadeOut(IBIOSSettings* csettings,
    PRBool* clogoFadeOut) {
  return IBIOSSettings_GetLogoFadeOut(csettings, clogoFadeOut);
}
HRESULT GoVboxSetBiosSettingsLogoFadeOut(IBIOSSettings* csettings,
    PRBool clogoFadeOut) {
  return IBIOSSettings_SetLogoFadeOut(csettings, clogoFadeOut);
}
HRESULT GoVboxGetBiosSettingsLogoDisplayTime(IBIOSSettings* csettings,
    PRUint32* cdisplayTime) {
  return IBIOSSettings_GetLogoDisplayTime(csettings, cdisplayTime);
}
HRESULT GoVboxSetBiosSettingsLogoDisplayTime(IBIOSSettings* csettings,
    PRUint32 cdisplayTime) {
  return IBIOSSettings_SetLogoDisplayTime(csettings, cdisplayTime);
}
HRESULT GoVboxGetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32* cmenuMode) {
  return IBIOSSettings_GetBootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxSetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32 cmenuMode) {
  return IBIOSSettings_SetBootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxGetBiosSettingsACPIEnabled(IBIOSSettings* csettings,
    PRBool* cacpiEnabled) {
  return IBIOSSettings_GetACPIEnabled(csettings, cacpiEnabled);
}
HRESULT GoVboxSetBiosSettingsACPIEnabled(IBIOSSettings* csettings,
    PRBool cacpiEnabled) {
  return IBIOSSettings_SetACPIEnabled(csettings, cacpiEnabled);
}
HRESULT GoVboxGetBiosSettingsIOAPICEnabled(IBIOSSettings* csettings,
    PRBool* cioapicEnabled) {
  return IBIOSSettings_GetIOAPICEnabled(csettings, cioapicEnabled);
}
HRESULT GoVboxSetBiosSettingsIOAPICEnabled(IBIOSSettings* csettings,
    PRBool cioapicEnabled) {
  return IBIOSSettings_SetIOAPICEnabled(csettings, cioapicEnabled);
}
HRESULT GoVboxGetBiosSettingsAPICMode(IBIOSSettings* csettings,
    PRUint32* capicMode) {
  return IBIOSSettings_GetAPICMode(csettings, capicMode);
}
HRESULT GoVboxSetBiosSettingsAPICMode(IBIOSSettings* csettings,
    PRUint32 capicMode) {
  return IBIOSSettings_SetAPICMode(csettings, capicMode);
}
HRESULT GoVboxGetBiosSettingsTimeOffset(IBIOSSettings* csettings,
    PRInt64* ctimeOffset) {
  return IBIOSSettings_GetTimeOffset(csettings, ctimeOffset);
}
HRESULT GoVboxSetBiosSettingsTimeOffset(IBIOSSettings* csettings,
    PRInt64 ctimeOffset) {
  return IBIOSSettings_SetTimeOffset(csettings, ctimeOffset);
}
HRESULT GoVboxSetBiosSettingsPXEDebugEnabled(IBIOSSettings* csettings,
    PRBool cPXEDebugEnabled) {
  return IBIOSSettings_SetPXEDebugEnabled(csettings, cPXEDebugEnabled);
}
HRESULT GoVboxGetBiosSettingsPXEDebugEnabled(IBIOSSettings* csettings,
    PRBool* cPXEDebugEnabled) {
  return IBIOSSettings_GetPXEDebugEnabled(csettings, cPXEDebugEnabled);
}
HRESULT GoVboxIBiosSettingsRelease(IBIOSSettings* csettings) {
  return IBIOSSettings_Release(csettings);
}

HRESULT GoVboxGetMachineBIOSSettings(IMachine* cmachine,
    IBIOSSettings** csettings) {
  return IMachine_GetBIOSSettings(cmachine, csettings);
}
