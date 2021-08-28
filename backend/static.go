package main

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

//go:embed static
var staticFs embed.FS

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
		index, err := staticFs.Open("static/index.html")
		if err != nil {
			panic(fmt.Errorf("Couldn't open index.html: %v", err))
		}
		if r.URL.Path == "/" {
			http.ServeContent(w, r, "index.html", time.Unix(0, 0), index.(io.ReadSeeker))
		} else {
			file, err := staticFs.Open("static" + r.URL.Path)
			if err != nil {
				http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), index.(io.ReadSeeker))
			} else {
				http.ServeContent(w, r, r.URL.Path, time.Unix(0, 0), file.(io.ReadSeeker))
			}
		}
	})

	localIP := getLocalIP()

	fmt.Println("Brood Clash is up and running!")
	fmt.Printf("Navigate to \"%s:5000/presenter\" on the presenter.\n", localIP)
	fmt.Printf("Then, navigate to \"%s:5000/controller\" on the controller.\n", localIP)
	http.ListenAndServe(":5000", handler)
}
