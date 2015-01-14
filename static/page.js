function Page(){
	$.ajaxSetup({ cache: false });
	this.NotConnected();
	this.Connect();
	this.val = false;
}

Page.prototype.Send = function(device, val){
	this.sock.send(device + ";" + val + ";");
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
		setTimeout(page.Connect, 1000 * 3);
	}
	this.sock.onclose = this.sock.onerror;

	this.sock.onopen = function(){
		$("body").empty();
		page.toggle();
	}

	this.sock.onmessage = function(message){
		var data = $.parseJSON(message.data);
		console.log(data);
	}
}

Page.prototype.toggle = function(){
	this.val = !this.val

	this.Send("button1", "" + this.val);
	var page = this;
	setTimeout(function(){page.toggle()}, 1000 * 1);
}

$(document).ready(function(){new Page();})