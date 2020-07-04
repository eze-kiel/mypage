package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/caddyserver/certmagic"
	"github.com/eze-kiel/hblanc/handlers"
)

func main() {
	var prod bool
	flag.BoolVar(&prod, "prod", false, "production mode")
	flag.Parse()
	switch prod {
	case true:
		// read and agree to your CA's legal documents
		certmagic.DefaultACME.Agreed = true

		// provide an email address
		certmagic.DefaultACME.Email = "hugoblanc@fastmail.com"

		fmt.Println("[PROD] Server is starting, wish me luck boys")
		certmagic.HTTPS([]string{"hugoblanc.com"}, handlers.Handle())

	case false:
		srv := &http.Server{
			Addr:         ":8080",
			Handler:      handlers.Handle(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}
		fmt.Println("[DEV] Server is starting, wish me luck boys")
		log.Println(srv.ListenAndServe())
	}

}
