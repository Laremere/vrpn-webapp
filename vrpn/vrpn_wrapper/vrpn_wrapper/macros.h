#ifdef __cplusplus
#define EXTERN extern "C" __declspec( dllexport ) 
#else
#define EXTERN
#include <stdbool.h>
#endif