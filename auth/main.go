package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rs401/myauth/auth/models"
	"github.com/rs401/myauth/auth/repository"
	"github.com/rs401/myauth/db"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}
}

func main() {
	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	conn.DB().AutoMigrate(&models.User{})
	usersRepository := repository.NewUsersRepository(conn)
	users, err := usersRepository.GetAll()
	if err != nil {
		log.Fatalf("Error retrieving users: %v\n", err)
	}

	fmt.Println(users)
}
