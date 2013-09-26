package main

import (
	"log"
	"net/http"
	"strconv"
)

func StartHttp(port int) {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/config/", http.StripPrefix("/config/", http.FileServer(http.Dir("config"))))

	log.Fatal(http.ListenAndServe(":"+strconv.FormatInt(int64(port), 10), nil))
}
