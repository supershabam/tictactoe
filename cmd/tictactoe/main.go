package main

import (
	"flag"
	"log"

	"github.com/supershabam/tictactoe"
)

var (
	addr   = flag.String("addr", ":8080", "address to listen on")
	public = flag.String("public", "./", "path to public static assets")
)

func main() {
	flag.Parse()
	s := tictactoe.Server{
		Addr:   *addr,
		Public: *public,
	}
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
