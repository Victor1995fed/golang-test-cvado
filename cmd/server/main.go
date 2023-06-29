package main

import (
	"fmt"
	"github.com/Victor1995fed/golang-test-cvado/config"
	"github.com/Victor1995fed/golang-test-cvado/internal/app"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	//Загрузка параметров из env
	cfg := config.ParseConfig(".env")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//Подключение БД
	db := database.Connect(&cfg)
	//Запуск сервера
	s := grpc.NewServer()
	cvadoProto.RegisterCvadoServer(s, app.NewServer(database.New(db)))
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
