package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(" Starting server on port 8080")
	log.Print("GET localhost:8080/greeting/ is handled")
	mux := http.NewServeMux()
	mux.HandleFunc("/greeting/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", RequestLogger(mux)))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second) //Doing hard work
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Hello Page", "Hello Web!")
}

func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}