#include "macros.h"

#ifdef __cplusplus
using namespace std;
#endif

EXTERN void* vrpn_Connection_New(int port);

EXTERN void vrpn_Connection_Mainloop(void* conn);
