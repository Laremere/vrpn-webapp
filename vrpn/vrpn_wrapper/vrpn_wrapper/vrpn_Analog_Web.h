#pragma once
#include "vrpn_Analog.h"
#include "vrpn_Connection.h"
using namespace std;

class vrpn_Analog_Web: public vrpn_Analog
{
public:
	vrpn_Analog_Web(const char* name, vrpn_Connection *c, int channels);
	virtual ~vrpn_Analog_Web(void);
	virtual void mainloop();

	void update(double value, int i);
};

