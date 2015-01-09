function Page(){
	var page = this;
	this.view = {};
	$.ajaxSetup({ cache: false });

	this.sock = new WebSocket("ws://" + window.location.host + "/sock/");
	this.sock.onerror = function(event){
		//Todo: clear buttons, try for reconnect
	}

	this.sock.onopen = function(){
	}
}

Page.prototype.Send = function(device){
	var message = device;
	for (var i = 1; i < arguments.length; i ++){
		message += "@" + arguments[i];
	}
	this.sock.send(message + ";");
}

$(document).ready(function(){new Page();})