#pragma once
#include "vrpn_Button.h"
#include "vrpn_Connection.h"
using namespace std;

class vrpn_Button_Web: public vrpn_Button
{
public:
	vrpn_Button_Web(const char* name, vrpn_Connection *c, int channels);
	virtual ~vrpn_Button_Web(void);
	virtual void mainloop();

	void update(bool state, int i);
};

