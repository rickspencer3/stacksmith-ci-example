package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	listenAddr = flag.String("l", ":8080", "listen address and port number")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello WORLD\n")
}

func main() {
	flag.Parse()

	http.HandleFunc("/", handler)

	log.Print("listen on", *listenAddr)
	log.Fatal(http.ListenAndServe(*listenAddr, nil))
}
