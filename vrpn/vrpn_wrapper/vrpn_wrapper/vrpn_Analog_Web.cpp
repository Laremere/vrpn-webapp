#include "vrpn_Analog_Web.h"
using namespace std;


vrpn_Analog_Web::vrpn_Analog_Web(const char* name, vrpn_Connection *c, int channels):
	vrpn_Analog(name, c){
	num_channel = channels;
}

vrpn_Analog_Web::~vrpn_Analog_Web(void){
}

void vrpn_Analog_Web::mainloop(){
	vrpn_gettimeofday(&(timestamp), NULL);
	report();
	server_mainloop();
}

void vrpn_Analog_Web::update(double value, int i){
	channel[i] = value;
}