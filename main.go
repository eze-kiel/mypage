package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/eze-kiel/hblanc/handlers"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handlers.Handle(),
	}

	fmt.Printf("[%v] : Serving...\n", time.Now().Format("Mon Jan 2 15:04:05"))
	srv.ListenAndServe()
}
