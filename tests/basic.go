package main

import (
	"github.com/zachrip/express"
	"net/http"
	"log"
)

func main() {
	s := express.Server(":8080")
	
	s.Get("/testing", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"));
	})
	
	s.Get("/test", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"));
	});
	
	go s.Listen()
	log.Printf("Listening");
	select {}
}