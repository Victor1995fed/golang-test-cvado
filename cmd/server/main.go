package main

import (
	"fmt"
	"github.com/Victor1995fed/golang-test-cvado/config"
	"github.com/Victor1995fed/golang-test-cvado/internal/app"
	"github.com/Victor1995fed/golang-test-cvado/internal/repository/database"
	cvadoProto "github.com/Victor1995fed/golang-test-cvado/proto"
	"github.com/caarlos0/env/v6"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {

	//Загрузка параметров из env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load env file: %e", err)
	}

	cfg := config.Config{}
	err = env.Parse(&cfg)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.App.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//Подключение БД
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//Запуск сервера
	s := grpc.NewServer()
	cvadoProto.RegisterCvadoServer(s, app.NewServer(database.New(db)))
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
