package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/eze-kiel/hblanc/handlers"
)

func main() {
	var prod bool

	flag.BoolVar(&prod, "prod", false, "production mode")
	flag.Parse()

	if prod {
		srv := &http.Server{
			Addr:    ":80",
			Handler: handlers.Handle(),
		}
		fmt.Printf("[PROD | %v] : Serving...\n", time.Now().Format("Mon Jan 2 15:04:05"))
		srv.ListenAndServe()
	} else {
		srv := &http.Server{
			Addr:    ":8080",
			Handler: handlers.Handle(),
		}
		fmt.Printf("DEV | [%v] : Serving...\n", time.Now().Format("Mon Jan 2 15:04:05"))
		srv.ListenAndServe()
	}

}
