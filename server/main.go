package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/wafer-bw/versioning-prototyping/pkg/grouped/c"
	"github.com/wafer-bw/versioning-prototyping/pkg/grouped/d"
	"github.com/wafer-bw/versioning-prototyping/pkg/individual/a"
	"github.com/wafer-bw/versioning-prototyping/pkg/individual/b"
)

const (
	host    string        = "localhost"
	port    string        = "2023"
	timeout time.Duration = 10 * time.Second
)

func main() {
	r := http.NewServeMux()

	r.Handle("/a", http.HandlerFunc(a.HandleA))
	r.Handle("/b", http.HandlerFunc(b.HandleB))
	r.Handle("/c", http.HandlerFunc(c.HandleC))
	r.Handle("/d", http.HandlerFunc(d.HandleD))

	s := http.Server{
		Addr:              net.JoinHostPort(host, port),
		ReadHeaderTimeout: timeout,
		WriteTimeout:      timeout,
		IdleTimeout:       timeout,
		Handler:           r,
	}

	log.Println("Starting server on", s.Addr)
	log.Fatal(s.ListenAndServe())
}
