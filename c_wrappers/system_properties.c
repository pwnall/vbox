#include "glue.h"

HRESULT GoVboxGetSystemProperties(IVirtualBox* cbox,
    ISystemProperties **cprops) {
  return IVirtualBox_get_SystemProperties(cbox, cprops);
}

HRESULT GoVboxGetSystemPropertiesMaxGuestRAM(ISystemProperties* cprops,
    ULONG *cmaxRam) {
  return ISystemProperties_get_MaxGuestRAM(cprops, cmaxRam);
}
HRESULT GoVboxGetSystemPropertiesMaxGuestVRAM(ISystemProperties* cprops,
    ULONG *cmaxVram) {
  return ISystemProperties_get_MaxGuestVRAM(cprops, cmaxVram);
}
HRESULT GoVboxGetSystemPropertiesMaxGuestCpuCount(ISystemProperties* cprops,
    ULONG *cmaxCpus) {
  return ISystemProperties_get_MaxGuestVRAM(cprops, cmaxCpus);
}
HRESULT GoVboxISystemPropertiesRelease(ISystemProperties* cprops) {
  return ISystemProperties_Release(cprops);
}
