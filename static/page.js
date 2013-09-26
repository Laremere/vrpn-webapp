function Page(){
	var page = this;
	this.view = {};
	this.view.buttons = ko.observableArray();
	this.view.sliders = ko.observableArray();

	ko.applyBindings(this.view);

	$.getJSON("/config/config.json")
		.done(function(d){page.ProcessConfig(d);})
		.fail(function(d){page.FailedConfig(d);});
}

Page.prototype.FailedConfig = function( obj, status, error ){
	console.log("Loading config error: ");
	console.log(status);
	console.log(error);
	alert("Error loading config.json, js console for details.")
} 

Page.prototype.ProcessConfig = function( data ){
	var device;
	for (var i = 0; i < data.devices.length; i++){
		device = data.devices[i];
		switch (device["type"]){
			case "button":
				this.NewButton(device)
				break;
			case "slider":
				this.NewSlider(device)
				break;
			default:
				console.log("Unkown device type: " + device["type"])
		}
	}
}

Page.prototype.NewButton = function(device){
	this.view.buttons.push(device);
}

Page.prototype.NewSlider = function(device){
	device.val = ko.observable();
	device.val(device.initial);
	this.view.sliders.push(device);
}

$(document).ready(function(){page = new Page();})