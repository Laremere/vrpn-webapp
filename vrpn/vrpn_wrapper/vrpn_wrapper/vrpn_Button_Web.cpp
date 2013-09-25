#include "vrpn_Button_Web.h"
using namespace std;

vrpn_Button_Web::vrpn_Button_Web(const char* name, vrpn_Connection *c, int channels) :
	vrpn_Button(name, c)
{
	vrpn_Button::num_buttons = channels;
}

vrpn_Button_Web::~vrpn_Button_Web(void)
{
}


void vrpn_Button_Web::update(bool state, int i){
	vrpn_Button::buttons[i] = state;
}


void vrpn_Button_Web::mainloop(){
    vrpn_gettimeofday(&(vrpn_Button::timestamp), NULL);
    vrpn_Button::report_changes();
	server_mainloop();
}
