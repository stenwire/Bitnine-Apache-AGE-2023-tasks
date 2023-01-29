package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host = ""
	port = 5432
	user = ""
	password = ""
	dbname = ""
)

type sandbox struct {
	user_id int
	name string
	age int
	phone string
}

// load .env file
func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

// main function
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	goDotEnvVariable("host"), port, goDotEnvVariable("user"), goDotEnvVariable("password"), goDotEnvVariable("dbname"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

