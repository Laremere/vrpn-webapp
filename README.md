vrpn-webapp
===========

A vrpn server and http server for webapp controlers

Building
===========
To build the shared library (DLL) you'll need vrpn.  The main vrpn static library needs to be linked along with adding the top level vrpn folder to the include path.  I also linked/included the quat directory in vrpn, but I'm unsure if that's required.

The go build command is used to build the final executable.

The only files/folders needed in realese should be the executable, the dll, the static folder and the config folder.

Config
===========
Two files in the config folder are used for configuration of the devices you want to use.  Config.json contains the configuration for the devices and the ports to use for the vrpn and html servers.  Style.css is available to use if you want to change the look of your interface.  Each element has it's set ID based on the vrpn name of the device.  I recomend using Chrome's inspect capability to view the final html given by config.json.

Config.json must be a well formated json file.  The vrpnPort parameter defines which port vrpn should use to create it's server, likewise with the httpPort.
The parameter "devices" is an array of devices to display on screen.  The order of devices does NOT define the order of devices on the page.  Use css if you want to edit the positions of the devices on the page.  The parameters of each device is as follows:

Any Device
-----------
* _name_ - Defines the name of the vrpn device, and the id of the html dom element.
* _class_ - Which type of device it is.  This defines both of what type of control to use on the webpage, and also what kind of vrpn device it is.
* _displayName_ - The name or label to display on the webpage.

Class "button"
-----------
A vrpn_button with one channel.  Displays as an html button, activating when pressed.

_No additional parameters_

Example:

	{
		"name":"Button0",
		"class":"button",
		"displayName":"Clickable Button"
	}

Class "toggle"
-----------
A vrpn_button with one channel.  Displays as a checkbox, activating when the checkbox is checked.

* _initial_ - Whether to start checked or not.

Example:

	{
		"name":"Button0",
		"class":"toggle",
		"displayName":"Checkbox",
		"initial":true
	}

Class "slider"
-----------
A vrpn_analog with one channel.  Displays as a range slider, returning the value of the slider.

* _range_ - An array of the minumum and maximum values.
* _initial_ - The initial value of the slider.  Must be an integer value.
* _step_ - How much the value changes by moving slider one tick.

Example:

	{
		"name":"Analog0",
		"class":"slider",
		"displayName":"Example SLider",
		"range":[0,100],
		"initial":50,
		"step": 1
	}

Class "spinner"
-----------
A vrpn_analog with one channel.  Displays as a textbox with a number in it with up and down buttons on the side, returning the value of the textbox.

* _range_ - An array of the minumum and maximum values.  NOTE: Users can enter values outside this range, but using the up and down arrows is limited by this range.
* _initial_ - The initial value of the textbox.
* _step_ - How much the value changes by one click.

Example:

	{
		"name":"Analog1",
		"class":"spinner",
		"displayName":"A textbox with number controls",
		"range":[0,100],
		"initial":50,
		"step": 1
	}
