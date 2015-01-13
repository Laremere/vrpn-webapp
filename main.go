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

	//Temporary debug
	fmt.Println(config)
}

//The root for configuration
type Config struct {
	VrpnPort uint16    `toml:"vrpn_port"`
	HttpPort uint16    `toml:"http_port"`
	Buttons  []Button  `toml:"button"`
	Toggles  []Toggle  `toml:"toggle"`
	Sliders  []Slider  `toml:"slider"`
	Spinners []Spinner `toml:"spinner"`
}

//Interface for reading device's information
type DeviceConfig interface {
	GetName() string    //Device Name
	GetInitial() string //Initial Value
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
