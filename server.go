package boxcars

import (
	"fmt"
	"log"
	"net/http"
)

func Listen(port int) {
	log.Printf("Starting at %d", port)
	http.HandleFunc("/", OnRequest)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
