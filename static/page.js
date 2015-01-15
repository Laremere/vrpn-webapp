function Page(){
	$.ajaxSetup({ cache: false });
	this.NotConnected();
	this.Connect();
	this.devices = {};
	this.deviceLastChange = {}
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
	}

	this.sock.onmessage = function(message){
		var data = $.parseJSON(message.data);
		switch(data.MessageType){
			case "new_button":
				page.New_Button(data.Data);
				return;
			case "new_toggle":
				page.New_Toggle(data.Data);
				return;
			case "new_slider":
				page.New_Slider(data.Data);
				return;
			case "new_spinner":
				page.New_Spinner(data.Data);
				return;
			case "event":
				page.Event(data.Data);
				return;	
		}
		console.log("Unknown Message Type");
		console.log(data);
	}
}

Page.prototype.New_Button = function(data){
	var page = this;
	page.deviceLastChange[data.Name] = false;
	var button = $('<button/>', {
		id: data.Name,
		text: data.Display
	});
	page.devices[data.Name] = function(){};
	button.mousedown(function(){
		page.Send(data.Name, "true");
	});
	button.mouseup(function(){
		page.Send(data.Name, "false");
	});
	button.appendTo('body');
	$("<br/>").appendTo('body');
}

Page.prototype.New_Toggle = function(data){
	var page = this;
	page.deviceLastChange[data.Name] = false;
	var toggle = $("<input/>", {
		type: "checkbox",
		id: data.Name
	});
	toggle.change(function(){
		page.Send(data.Name, this.checked);
	});
	page.devices[data.Name] = function(newValue){
		toggle.prop('checked', newValue === "true");
	};
	toggle.appendTo('body');
	$("<label/>", {
		text: data.Display,
		for: data.Name
	}).appendTo('body');
	$("<br/>").appendTo('body');
}

Page.prototype.New_Slider = function(data){
	var page = this;
	page.deviceLastChange[data.Name] = false;
	var slider = $("<input/>", {
		type: "range",
		id: data.Name,
		min: data.Range[0],
		max: data.Range[1],
		step: data.Step
	});

	slider.on("input", function(){
		page.Send(data.Name, this.value);
	});
	page.devices[data.Name] = function(newValue){
		slider.prop('value', newValue);
	};

	$("<label/>", {
		text: data.Display,
		for: data.Name
	}).appendTo('body');
	slider.appendTo('body');
	$("<br/>").appendTo('body');
}

Page.prototype.New_Spinner = function(data){
	var page = this;
	page.deviceLastChange[data.AName] = false;
	var spinner = $("<input/>", {
		type: "number",
		id: data.Name,
		min: data.Range[0],
		max: data.Range[1],
		step: data.Step
	});

	spinner.on("input", function(){
		page.Send(data.Name, this.value);
	});
	page.devices[data.Name] = function(newValue){
		spinner.prop('value', newValue);
	};

	$("<label/>", {
		text: data.Display,
		for: data.Name
	}).appendTo('body');
	spinner.appendTo('body');
	$("<br/>").appendTo('body');
}

Page.prototype.Event = function(data){
	console.log(data);
	this.devices[data.Device](data.Value);
}

$(document).ready(function(){new Page();})