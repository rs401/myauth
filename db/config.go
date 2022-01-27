package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	ConnStr() string
	DbName() string
}

type config struct {
	dbUser  string
	dbPass  string
	dbHost  string
	dbPort  int
	dbName  string
	connStr string
}

func NewConfig() Config {
	var cfg config
	var err error
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}
	cfg.connStr = fmt.Sprintf("host=%s user=%s  password=%s  dbname=%s port=%d",
		cfg.dbHost, cfg.dbUser, cfg.dbPass, cfg.dbName, cfg.dbPort)
	return &cfg
}

func (c *config) ConnStr() string {
	return c.connStr
}
func (c *config) DbName() string {
	return c.dbName
}
