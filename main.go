package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("EggTrail Edge Agent v2.6.1-spring starting...")
	
	// Internal telemetry endpoint
	http.HandleFunc("/v1/telemetry", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Telemetry Node Active. Mode: Spring-2026")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
