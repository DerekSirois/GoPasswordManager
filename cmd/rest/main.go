package main

import (
	"log"
	"os"
	"passwordManager/pkg/database"
)

func main() {
	err := database.InitDb()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
