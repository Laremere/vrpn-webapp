#include "vrpn_Connection_Wrapper.h"
#include "vrpn_Connection.h"

using namespace std;


void* vrpn_Connection_New(){
    return vrpn_create_server_connection(3883);
}


void vrpn_Connection_Mainloop(void* conn)
{
    vrpn_Connection* _conn = (vrpn_Connection*) conn;
    _conn -> mainloop();
}
