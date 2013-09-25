// listenTest.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "vrpn_Button.h"
#include "vrpn_Analog.h"
#include <stdlib.h>

void VRPN_CALLBACK callback(void *userdata, vrpn_BUTTONCB b){
	printf("%s: B%ld is %ld\n", userdata, b.button, b.state);
}

void	VRPN_CALLBACK handle_analog(void *userdata, vrpn_ANALOGCB b) 
{ 
  for (int i=0; i< b.num_channel; i++)
   printf("Chan[%d] = %lf\n", i, b.channel[i]);
} 

int _tmain(int argc, _TCHAR* argv[])
{
	char* buttonName = "Button0@localhost";
	vrpn_Button_Remote myButton (buttonName);
	myButton.register_change_handler((void*)buttonName, callback);

	char* analogName = "Analog0@localhost";
	vrpn_Analog_Remote myAnalog(analogName);
	myAnalog.register_change_handler((void*)analogName, handle_analog);

	while (true){
		myButton.mainloop();
		myAnalog.mainloop();
	}
	return 0;
}


