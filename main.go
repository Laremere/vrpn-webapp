package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"golang.org/x/net/websocket"
	"log"
	"net"
	"net/http"
)

func main() {
	//The main config object
	var config Config

	//Decode the config file
	_, err := toml.DecodeFile("config/config.toml", &config)
	if err != nil {
		//Report error and exit.
		fmt.Println("Error reading the config file, exiting.")
		fmt.Println(err)
		return
	}

	var devices []DeviceConfig
	for _, button := range config.Buttons {
		devices = append(devices, button)
	}
	for _, toggle := range config.Toggles {
		devices = append(devices, toggle)
	}
	for _, slider := range config.Sliders {
		devices = append(devices, slider)
	}
	for _, spinner := range config.Spinners {
		devices = append(devices, spinner)
	}

	go VrpnServe(config.VrpnPort, devices)
	//Start connection and event manager.
	go Manager(devices)

	var portStr string
	if config.HTTPPort != 80 {
		portStr = fmt.Sprintf(":%d", config.HTTPPort)
	}
	{
		log.Println("Server starting.  Point a web browser to one of the following:")
		log.Println("On this machine: localhost" + portStr)
		inters, err := net.Interfaces()
		if err != nil {
			fmt.Println("Error getting addresses, exiting")
			fmt.Println(err)
			return
		}
		for _, inter := range inters {
			addrs, err := inter.Addrs()
			if err != nil {
				fmt.Println("Error getting addresses, exiting")
				fmt.Println(err)
				return
			}
			for _, addr := range addrs {
				log.Printf("Via %s: %s%s", inter.Name, addr.String(), portStr)
			}
		}
	}

	http.Handle("/", http.FileServer(http.Dir("static")))
	wsh := WebSocketHandler{devices}
	http.Handle("/sock/", websocket.Handler(wsh.HandleSocket))

	err = http.ListenAndServe(portStr, nil)
	fmt.Println("Error starting webserver, exiting.")
	fmt.Println(err)
}

//Config holds the data from the configuration file
type Config struct {
	VrpnPort uint16     `toml:"vrpn_port"`
	HTTPPort uint16     `toml:"http_port"`
	Buttons  []*Button  `toml:"button"`
	Toggles  []*Toggle  `toml:"toggle"`
	Sliders  []*Slider  `toml:"slider"`
	Spinners []*Spinner `toml:"spinner"`
}

//DeviceConfig allows generic access devices
type DeviceConfig interface {
	GetName() string    //Device Name
	GetInitial() string //Initial Value
	VrpnType() VrpnType //VRPN analog or button
	WebType() string    //info for web browser
}

//Button represents a clickable button.
//See the default config.toml for documentation on components.
type Button struct {
	Name    string
	Display string
}

//GetName gets the device's vrpn name.
func (b *Button) GetName() string {
	return b.Name
}

//GetInitial gets the initial state of the device.
func (b *Button) GetInitial() string {
	return "false"
}

//VrpnType gets the vrpn representation of the device.
func (*Button) VrpnType() VrpnType {
	return VrpnButton
}

//WebType gets the webpage representation of the device.
func (*Button) WebType() string {
	return "button"
}

//Toggle epresents a toggle box.
//See the default config.toml for documentation on components.
type Toggle struct {
	Name    string
	Display string
	Initial bool
}

//GetName gets the device's vrpn name.
func (t *Toggle) GetName() string {
	return t.Name
}

//GetInitial gets the initial state of the device.
func (t *Toggle) GetInitial() string {
	return fmt.Sprint(t.Initial)
}

//VrpnType gets the vrpn representation of the device.
func (*Toggle) VrpnType() VrpnType {
	return VrpnButton
}

//WebType gets the webpage representation of the device.
func (*Toggle) WebType() string {
	return "toggle"
}

//Slider represents a slider bar.
//See the default config.toml for documentation on components.
type Slider struct {
	Name    string
	Display string
	Range   [2]float64
	Initial float64
	Step    float64
}

//GetName gets the device's vrpn name.
func (s *Slider) GetName() string {
	return s.Name
}

//GetInitial gets the initial state of the device.
func (s *Slider) GetInitial() string {
	return fmt.Sprint(s.Initial)
}

//VrpnType gets the vrpn representation of the device.
func (*Slider) VrpnType() VrpnType {
	return VrpnAnalog
}

//WebType gets the webpage representation of the device.
func (*Slider) WebType() string {
	return "slider"
}

//Spinner represents a number selector.
//See the default config.toml for documentation on components.
type Spinner struct {
	Name    string
	Display string
	Range   [2]float64
	Initial float64
	Step    float64
}

//GetName gets the device's vrpn name.
func (s *Spinner) GetName() string {
	return s.Name
}

//GetInitial gets the initial state of the device.
func (s *Spinner) GetInitial() string {
	return fmt.Sprint(s.Initial)
}

//VrpnType gets the vrpn representation of the device.
func (*Spinner) VrpnType() VrpnType {
	return VrpnAnalog
}

//WebType gets the webpage representation of the device.
func (*Spinner) WebType() string {
	return "spinner"
}
