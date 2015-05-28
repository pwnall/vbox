#include "VBoxCAPIGlue.h"

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);


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
HRESULT GoVboxGetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32* cmenuMode) {
  return IBIOSSettings_GetBootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxSetBiosSettingsBootMenuMode(IBIOSSettings* csettings,
    PRUint32 cmenuMode) {
  return IBIOSSettings_SetBootMenuMode(csettings, cmenuMode);
}
HRESULT GoVboxIBiosSettingsRelease(IBIOSSettings* csettings) {
  return IBIOSSettings_Release(csettings);
}

HRESULT GoVboxGetMachineBIOSSettings(IMachine* cmachine,
    IBIOSSettings** csettings) {
  return IMachine_GetBIOSSettings(cmachine, csettings);
}
