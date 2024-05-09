package db

type DB struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}
