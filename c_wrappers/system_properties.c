#include "VBoxCAPIGlue.h"

// Wrapper declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);


HRESULT GoVboxGetSystemProperties(IVirtualBox* cbox,
    ISystemProperties **cprops) {
  return IVirtualBox_GetSystemProperties(cbox, cprops);
}

HRESULT GoVboxGetSystemPropertiesMaxGuestRAM(ISystemProperties* cprops,
    ULONG *cmaxRam) {
  return ISystemProperties_GetMaxGuestRAM(cprops, cmaxRam);
}
HRESULT GoVboxGetSystemPropertiesMaxGuestVRAM(ISystemProperties* cprops,
    ULONG *cmaxVram) {
  return ISystemProperties_GetMaxGuestVRAM(cprops, cmaxVram);
}
HRESULT GoVboxGetSystemPropertiesMaxGuestCpuCount(ISystemProperties* cprops,
    ULONG *cmaxCpus) {
  return ISystemProperties_GetMaxGuestVRAM(cprops, cmaxCpus);
}
HRESULT GoVboxISystemPropertiesRelease(ISystemProperties* cprops) {
  return ISystemProperties_Release(cprops);
}

