#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


HRESULT GoVboxGetBiosSettingsLogoImagePath(IBIOSSettings* csettings,
    char** clogoImagePath) {
  BSTR wlogoImagePath = NULL;
  HRESULT result = IBIOSSettings_get_LogoImagePath(csettings, &wlogoImagePath);
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

  result = IBIOSSettings_put_LogoImagePath(csettings, wlogoImagePath);
  g_pVBoxFuncs->pfnUtf16Free(wlogoImagePath);

  return result;
}
HRESULT GoVboxGetBiosSettingsLogoFadeIn(IBIOSSettings* csettings,
    PRBool* clogoFadeIn) {
  return IBIOSSettings_get_LogoFadeIn(csettings, clogoFadeIn);
}
HRESULT GoVboxSetBiosSettingsLogoFadeIn(IBIOSSettings* csettings,
    PRBool clogoFadeIn) {
  return IBIOSSettings_put_LogoFadeIn(csettings, clogoFadeIn);
}
HRESULT GoVboxGetBiosSettingsLogoFadeOut(IBIOSSettings* csettings,
    PRBool* clogoFadeOut) {
  return IBIOSSettings_get_LogoFadeOut(csettings, clogoFadeOut);
}
HRESULT GoVboxSetBiosSettingsLogoFadeOut(IBIOSSettings* csettings,
    PRBool clogoFadeOut) {
  return IBIOSSettings_put_LogoFadeOut(csettings, clogoFadeOut);
}
HRESULT GoVboxGetBiosSettingsLogoDisplayTime(IBIOSSettings* csettings,
    PRUint32* cdisplayTime) {
  return IBIOSSettings_get_LogoDisplayTime(csettings, cdisplayTime);
}
HRESULT GoVboxSetBiosSettingsLogoDisplayTime(IBIOSSettings* csettings,
    PRUint32 cdisplayTime) {
  return IBIOSSettings_put_LogoDisplayTime(csettings, cdisplayTime);
}
HRESULT GoVboxGetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32* cmenuMode) {
  return IBIOSSettings_get_BootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxSetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32 cmenuMode) {
  return IBIOSSettings_put_BootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxGetBiosSettingsACPIEnabled(IBIOSSettings* csettings,
    PRBool* cacpiEnabled) {
  return IBIOSSettings_get_ACPIEnabled(csettings, cacpiEnabled);
}
HRESULT GoVboxSetBiosSettingsACPIEnabled(IBIOSSettings* csettings,
    PRBool cacpiEnabled) {
  return IBIOSSettings_put_ACPIEnabled(csettings, cacpiEnabled);
}
HRESULT GoVboxGetBiosSettingsIOAPICEnabled(IBIOSSettings* csettings,
    PRBool* cioapicEnabled) {
  return IBIOSSettings_get_IOAPICEnabled(csettings, cioapicEnabled);
}
HRESULT GoVboxSetBiosSettingsIOAPICEnabled(IBIOSSettings* csettings,
    PRBool cioapicEnabled) {
  return IBIOSSettings_put_IOAPICEnabled(csettings, cioapicEnabled);
}
HRESULT GoVboxGetBiosSettingsAPICMode(IBIOSSettings* csettings,
    PRUint32* capicMode) {
  return IBIOSSettings_get_APICMode(csettings, capicMode);
}
HRESULT GoVboxSetBiosSettingsAPICMode(IBIOSSettings* csettings,
    PRUint32 capicMode) {
  return IBIOSSettings_put_APICMode(csettings, capicMode);
}
HRESULT GoVboxGetBiosSettingsTimeOffset(IBIOSSettings* csettings,
    PRInt64* ctimeOffset) {
  return IBIOSSettings_get_TimeOffset(csettings, ctimeOffset);
}
HRESULT GoVboxSetBiosSettingsTimeOffset(IBIOSSettings* csettings,
    PRInt64 ctimeOffset) {
  return IBIOSSettings_put_TimeOffset(csettings, ctimeOffset);
}
HRESULT GoVboxSetBiosSettingsPXEDebugEnabled(IBIOSSettings* csettings,
    PRBool cPXEDebugEnabled) {
  return IBIOSSettings_put_PXEDebugEnabled(csettings, cPXEDebugEnabled);
}
HRESULT GoVboxGetBiosSettingsPXEDebugEnabled(IBIOSSettings* csettings,
    PRBool* cPXEDebugEnabled) {
  return IBIOSSettings_get_PXEDebugEnabled(csettings, cPXEDebugEnabled);
}
HRESULT GoVboxIBiosSettingsRelease(IBIOSSettings* csettings) {
  return IBIOSSettings_Release(csettings);
}

HRESULT GoVboxGetMachineBIOSSettings(IMachine* cmachine,
    IBIOSSettings** csettings) {
  return IMachine_get_BIOSSettings(cmachine, csettings);
}
