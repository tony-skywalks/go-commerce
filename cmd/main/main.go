package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/tony-skywalks/my-web/pkg/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	fmt.Println(fmt.Sprintf("Server Listening On Port :: %v", port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%v", port), r))
}
