package main

import (
	"log"

	"github.com/EnriqueBravo115/ecommerce/db"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type config struct {
	db db.DB
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.Use(middleware.Logger())

	e.Logger.Fatal(e.Start(":3000"))
}
