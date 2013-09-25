#include "vrpn_Analog_Web_Wrapper.h"
#include "vrpn_Analog_Web.h"

using namespace std;

void* vrpn_Analog_Web_New(char* name, void *_c, int channels){
	vrpn_Connection *c = (vrpn_Connection*) _c;
	vrpn_Analog_Web* self = new vrpn_Analog_Web(name, c, channels);
	return (void*)self;
}

void vrpn_Analog_Web_Delete(void* _self){
	vrpn_Analog_Web* self = (vrpn_Analog_Web*) _self;
	delete self;
}

void vrpn_Analog_Web_Update(void* _self, double value, int index ){
	vrpn_Analog_Web* self = (vrpn_Analog_Web*) _self;
	self->update( value, index);
}

void vrpn_Analog_Web_Mainloop(void* _self){
	vrpn_Analog_Web* self = (vrpn_Analog_Web*) _self;
	self->mainloop();
}