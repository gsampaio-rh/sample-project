package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tomasen/realip"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func driver(w http.ResponseWriter, r *http.Request) {
	clientIP := realip.FromRequest(r)
	log.Println("GET /driver from", clientIP)
	message := "Hello Sergio, Your driver is arriving soon in a Hyundai HB20. After arriving, they'll wait 2 minutes before charges begin for time."
	enableCors(&w)
	w.Write([]byte(message))
}

func path(w http.ResponseWriter, r *http.Request) {
	clientIP := realip.FromRequest(r)
	log.Println("GET /path from", clientIP)
	message := "https://media.giphy.com/media/iNkZ9134iEdpu/source.gif"
	enableCors(&w)
	w.Write([]byte(message))
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	logger.Println("Server is starting...")
	http.HandleFunc("/", driver)
	http.HandleFunc("/driver", driver)
	http.HandleFunc("/path", path)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
