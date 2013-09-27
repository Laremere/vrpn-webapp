function Page(){
	var page = this;
	this.view = {};
	this.view.buttons = ko.observableArray();
	this.view.toggles = ko.observableArray();
	this.view.sliders = ko.observableArray();
	ko.applyBindings(this.view);

	this.sock = new WebSocket("ws://" + window.location.host + "/sock/");
	this.sock.onerror = function(event){
		console.log(event);
	}

	this.sock.onopen = function(){
		$.getJSON("/config/config.json")
			.done(function(d){page.ProcessConfig(d);})
			.fail(function(d){page.FailedConfig(d);});
	}

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
		switch (device["class"]){
			case "button":
				this.NewButton(device)
				break;
			case "toggle":
				this.NewToggle(device)
				break;
			case "slider":
				this.NewSlider(device)
				break;
			default:
				console.log("Unkown device class: " + device["class"])
		}
	}
}

Page.prototype.Send = function(device){
	var message = device;
	for (var i = 1; i < arguments.length; i ++){
		message += "@" + arguments[i];
	}
	this.sock.send(message + ";");
}

Page.prototype.NewButton = function(device){
	this.view.buttons.push(device);
}

Page.prototype.NewSlider = function(device){
	device.val = ko.observable();
	this.view.sliders.push(device);

	var page = this;
	device.val.subscribe(function(val){
		page.Send(device.name, 0, val);
	});
	device.val(device.initial);
}

Page.prototype.NewToggle = function(device){
	device.val = ko.observable();
	this.view.toggles.push(device);

	var page = this;
	device.val.subscribe(function(val){
		page.Send(device.name, 0, val);
		console.log(val);
	});
	device.val(device.initial);	
}

$(document).ready(function(){new Page();})