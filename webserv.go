package main

import (
	"bufio"
	"code.google.com/p/go.net/websocket"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func StartHttp(port int) {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("config"))))
	http.Handle("/sock/", websocket.Handler(handleWebsocket))

	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), nil))
}

func handleWebsocket(c *websocket.Conn) {
	log.Println("Websocket connected: ", c.RemoteAddr())
	buf := bufio.NewReader(c)
	for {
		command, err := buf.ReadString(byte(';'))
		if err != nil {
			break
		}
		splitCommand := strings.Split(command[0:len(command)-1], "@")
		target, ok := devices[splitCommand[0]]
		if ok {
			target.Update(splitCommand[1:])
		} else {
			log.Println("Uknown device: ", splitCommand[0])
		}

	}
	log.Println("Websocket disconnected: ", c.RemoteAddr())
}
