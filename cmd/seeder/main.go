package main

import (
	"context"
	"github.com/dennis0126/network-monitor/internal/config"
	"github.com/dennis0126/network-monitor/internal/db"
	"github.com/dennis0126/network-monitor/internal/repository"
	"github.com/dennis0126/network-monitor/internal/service"
	"github.com/jackc/pgx/v5"
	"log"
)

func main() {
	cfg := config.InitConfig()

	// database connection
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.DbString)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := db.New(conn)

	// repository
	userRepository := repository.NewUserRepository(ctx, queries)

	// service
	userService := service.NewUserService(userRepository)

	user, err := userService.GetUserByName("admin")
	if err != nil {
		log.Fatal(err)
	}
	if user != nil {
		log.Println("`admin` user already exists")
		return
	}
	if _, err := userService.CreateUser("admin", "12345678"); err != nil {
		log.Fatal(err)
	}
	log.Println("`admin` user is created")
}
