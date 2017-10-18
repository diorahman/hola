package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/hooqtv/hola/hola"
	x "golang.org/x/net/context"
)

// Service ...
type Service struct {
}

// Get ...
func (s *Service) Get(ctx x.Context, in *hola.GetRequest) (*hola.GetResponse, error) {
	log.Printf("get: %s", in.Key)
	res := new(hola.GetResponse)
	res.Value = in.Key
	return res, nil
}

// NewService ...
func NewService() (s *Service) {
	s = &Service{}
	return
}

func main() {
	port := flag.Int("port", 8000, "grpc port")

	flag.Parse()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		os.Exit(1)
	}

	gs := grpc.NewServer()
	hola.RegisterHolaServer(gs, NewService())

	log.Printf("starting grpc on :%d\n", *port)

	gs.Serve(listener)
}
