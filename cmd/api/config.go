package api

import "github.com/EnriqueBravo115/ecommerce/db"

type Config struct {
	port int
	env  string
	db   db.DB
}
