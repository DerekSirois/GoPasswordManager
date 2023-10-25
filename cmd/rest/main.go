package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"passwordManager/pkg/database"
	"passwordManager/pkg/handler"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = database.InitDb()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/password", handler.CreatePassword).Methods("POST")

	log.Println("Serving on port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
