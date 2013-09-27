package main

import (
	"encoding/json"
	"github.com/Laremere/vrpn-webapp/vrpn"
	"io/ioutil"
	"log"
)

type Config struct {
	VrpnPort int
	HttpPort int
	Devices  []ConfigDevice
}

type ConfigDevice struct {
	Name  string
	Class string
}

func ReadConfig() *Config {
	contents, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Fatal("Error opening config file: ", err)
	}
	results := new(Config)
	err = json.Unmarshal(contents, results)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}
	_ = contents
	return results
}

func CreateDevice(conn *vrpn.Connection, c ConfigDevice) *Device {
	switch {
	case c.Class == "button", c.Class == "toggle":
		return NewButtonDevice(conn, c.Name, 1)
	case c.Class == "slider":
		return NewAnalogDevice(conn, c.Name, 1)
	}

	log.Fatal("Unkown device class: ", c.Class)
	return nil
}
