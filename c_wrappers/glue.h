#include "VBoxCAPIGlue.h"

#ifdef WIN32
typedef unsigned char PRUint8;
typedef signed char PRInt8;
typedef unsigned short PRUint16;
typedef short PRInt16;
typedef unsigned int PRUint32;
typedef int PRInt32;
typedef long PRInt64;
typedef unsigned long PRUint64;
typedef int PRIntn;
typedef unsigned int PRUintn;
typedef double PRFloat64;
typedef size_t PRSize;
typedef ptrdiff_t PRPtrdiff;
typedef unsigned long PRUptrdiff;
typedef PRIntn PRBool;
#endif

// Wrappers declared in vbox.c
HRESULT GoVboxFAILED(HRESULT result);
HRESULT GoVboxArrayOutFree(void* array);
void GoVboxUtf8Free(char* cstring);
