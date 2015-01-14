package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
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

	devices := make([]DeviceConfig, 0)
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

	go vrpnServe(config.VrpnPort, devices)
	//Start connection and event manager.
	Manager()
}

//The root for configuration
type Config struct {
	VrpnPort uint16     `toml:"vrpn_port"`
	HttpPort uint16     `toml:"http_port"`
	Buttons  []*Button  `toml:"button"`
	Toggles  []*Toggle  `toml:"toggle"`
	Sliders  []*Slider  `toml:"slider"`
	Spinners []*Spinner `toml:"spinner"`
}

//Interface for reading device's information
type DeviceConfig interface {
	GetName() string    //Device Name
	GetInitial() string //Initial Value
	VrpnType() vrpnType
}

//Represents a clickable button.
//See the default config.toml for documentation on components.
type Button struct {
	Name    string
	Display string
}

func (b *Button) GetName() string {
	return b.Name
}

func (b *Button) GetInitial() string {
	return "false"
}

func (*Button) VrpnType() vrpnType {
	return vrpnButton
}

//Represents a toggle box.
//See the default config.toml for documentation on components.
type Toggle struct {
	Name    string
	Display string
	Initial bool
}

func (t *Toggle) GetName() string {
	return t.Name
}

func (t *Toggle) GetInitial() string {
	return fmt.Sprint(t.Initial)
}

func (*Toggle) VrpnType() vrpnType {
	return vrpnButton
}

//Represents a slider bar.
//See the default config.toml for documentation on components.
type Slider struct {
	Name    string
	Display string
	Range   [2]float64
	Initial float64
	Step    float64
}

func (s *Slider) GetName() string {
	return s.Name
}

func (s *Slider) GetInitial() string {
	return fmt.Sprint(s.Initial)
}

func (*Slider) VrpnType() vrpnType {
	return vrpnAnalog
}

//Represents a number selector.
//See the default config.toml for documentation on components.
type Spinner struct {
	Name    string
	Display string
	Range   [2]float64
	Initial float64
	Step    float64
}

func (s *Spinner) GetName() string {
	return s.Name
}

func (s *Spinner) GetInitial() string {
	return fmt.Sprint(s.Initial)
}

func (*Spinner) VrpnType() vrpnType {
	return vrpnAnalog
}
