#include "macros.h"

#ifdef __cplusplus
using namespace std;
#endif

EXTERN void* vrpn_Analog_Web_New(char* name, void *_c, int channels);

EXTERN void vrpn_Analog_Web_Delete(void* _self);

EXTERN void vrpn_Analog_Web_Update(void* _self, double value, int index );

EXTERN void vrpn_Analog_Web_Mainloop(void* _self);
