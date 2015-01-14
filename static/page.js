function Page(){
	$.ajaxSetup({ cache: false });
	this.NotConnected();
	this.Connect();
}

Page.prototype.Send = function(device){
	var message = device;
	for (var i = 1; i < arguments.length; i ++){
		message += "@" + arguments[i];
	}
	this.sock.send(message + ";");
}

Page.prototype.NotConnected = function(){
	var body = $("body");
	body.empty();
	body.append("<p>Not connected to server</p>")
}

Page.prototype.Connect = function(){
	var page = this;
	this.sock = new WebSocket("ws://" + window.location.host + "/sock/");
	this.sock.onerror = function(event){
		page.NotConnected();
		setTimeout(page.Connect(), 1000 * 3);
	}
	this.sock.onclose = this.sock.onerror;

	this.sock.onopen = function(){
	}

	this.sock.onmessage = function(){
	}
}

$(document).ready(function(){new Page();})