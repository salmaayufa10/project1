package main

import (
	"database/sql"
	"fmt"
	"library/config"
	"library/internal/handler"
	"library/internal/repository"
	"library/internal/service"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()
	fmt.Println("abc", cfg.DatabaseURL)
	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Printf("database ping failed (check env for real connection): %v", err)
	}

	repo := repository.NewBookRepository(db)
	svc := service.NewBookService(repo)
	userrepo := repository.NewUserRepository(db)
	usersvc := service.NewUserService(userrepo)
	h := handler.NewBookHandler(svc, usersvc)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler.RegisterRoutes(e, h)

	log.Printf("listening on :%s", cfg.Port)
	if err := e.Start(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
