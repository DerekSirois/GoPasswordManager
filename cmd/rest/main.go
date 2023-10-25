package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"passwordManager/pkg/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = database.InitDb()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
