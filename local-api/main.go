// api go with endpoint health-check and return 200 OK
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func health(w http.ResponseWriter, r *http.Request) {
	log.Println("Info: Health check")
	fmt.Fprintf(w, "OK %s", time.Now())
	log.Println("Info: Health check completed")
}

func main() {
	// get port from environment variable with default 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	log.Println("Info: Starting server")
	http.HandleFunc("/health-check", health)
	log.Printf("Server listening on port %s", port)
	http.ListenAndServe(":"+port, nil)
	log.Println("Info: Server stopped")
}
