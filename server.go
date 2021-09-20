package main

import (
	"fmt"
	"log"
	"net/http"

	_ "go_rest_api/docs"
	"go_rest_api/routes"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Posts API
// @version 1.0
// @description This is a sample rest api to add and retrieve posts
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email abc@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /

func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running...")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
