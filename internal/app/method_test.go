package app

import (
	"context"
	"github.com/Victor1995fed/golang-test-cvado/config"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
)

const (
	bookId   = 1
	authorId = 1
	envPath  = "../../.env"
)

// Тест метода на получение авторов
func TestApp_GetAuthorByBook(t *testing.T) {

	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()
	a, _ := client.GetAuthorByBook(ctx, &cvadoProto.RequestBook{Id: bookId})
	assert.NotEmpty(t, a)
}

// Тест метода на получение книг
func TestApp_GetBooksByAuthor(t *testing.T) {

	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()
	a, _ := client.GetBooksByAuthor(ctx, &cvadoProto.RequestAuthor{Id: authorId})
	assert.NotEmpty(t, a)
}

// Метод запуска сервера для тестов
func server(ctx context.Context) (cvadoProto.CvadoClient, func()) {

	lis := bufconn.Listen(1024 * 1024)
	cfg := config.ParseConfig(envPath)
	db := database.Connect(&cfg)
	s := grpc.NewServer()
	cvadoProto.RegisterCvadoServer(s, NewServer(database.New(db)))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		s.Stop()
	}

	client := cvadoProto.NewCvadoClient(conn)

	return client, closer
}
