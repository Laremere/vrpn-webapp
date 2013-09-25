#include "vrpn_Button_Web_Wrapper.h"
#include "vrpn_Button_Web.h"

using namespace std;

void* vrpn_Button_Web_New(char* name, void *_c, int channels){
	vrpn_Connection *c = (vrpn_Connection*) _c;
	vrpn_Button_Web* self = new vrpn_Button_Web(name, c, channels);
	return (void*)self;
}

void vrpn_Button_Web_Delete(void* _self){
	vrpn_Button_Web* self = (vrpn_Button_Web*) _self;
	delete self;
}

void vrpn_Button_Web_Update(void* _self, char active, int index ){
	vrpn_Button_Web* self = (vrpn_Button_Web*) _self;
	self->update((bool) active, index);
}

void vrpn_Button_Web_Mainloop(void* _self){
	vrpn_Button_Web* self = (vrpn_Button_Web*) _self;
	self->mainloop();
}