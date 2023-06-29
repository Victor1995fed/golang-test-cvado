package main

import (
	"context"
	"flag"
	"log"
	"time"

	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Id = 1
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := cvadoProto.NewCvadoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	a, err := c.GetAuthorByBook(ctx, &cvadoProto.RequestBook{Id: Id})
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}

	log.Printf("Response: %s", a.GetAuthors())

	b, err := c.GetBooksByAuthor(ctx, &cvadoProto.RequestAuthor{Id: Id})
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	log.Printf("Response: %s", b.GetBooks())
}
