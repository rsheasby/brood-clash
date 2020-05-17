package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/markbates/pkger"
)

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

func serveStatic() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir := pkger.Dir("/static")
		index, err := dir.Open("/index.html")
		if err != nil {
			panic(fmt.Errorf("Couldn't open index.html: %v", err))
		}
		if r.URL.Path == "/" {
			http.ServeContent(w, r, "index.html", time.Unix(0, 0), index)
		} else {
			file, err := dir.Open(r.URL.Path)
			if err != nil {
				http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), index)
			} else {
				http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), file)
			}
		}
	})

	localIP := getLocalIP()

	fmt.Println("Brood Clash is up and running!")
	fmt.Printf("Navigate to \"%s:5000/presenter\" on the presenter.\n", localIP)
	fmt.Printf("Then, navigate to \"%s:5000/controller\" on the controller.\n", localIP)
	http.ListenAndServe(":5000", handler)
}
