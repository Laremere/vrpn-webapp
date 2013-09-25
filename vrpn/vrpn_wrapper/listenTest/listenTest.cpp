// listenTest.cpp : Defines the entry point for the console application.
//

#include "stdafx.h"
#include "vrpn_Button.h"
#include <stdlib.h>

void VRPN_CALLBACK callback(void *userdata, vrpn_BUTTONCB b){
	printf("B%ld is %ld\n", b.button, b.state);
}

int _tmain(int argc, _TCHAR* argv[])
{
	vrpn_Button_Remote myButton ("Button0@localhost");
	myButton.register_change_handler(NULL, callback);

	while (true){
		myButton.mainloop();
	}
	return 0;
}


