package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/marceloemanoel/movieDB/router"
)

const DefaultPort int64 = 3000

func readPortConfig() int64 {
	portStr := os.Getenv("PORT")

	if portStr == "" {
		fmt.Println("PORT is not defined, using default port: ", DefaultPort)
		return DefaultPort
	}

	portNumber, error := strconv.ParseInt(portStr, 10, 0)
	if error != nil {
		fmt.Println("Invalid port number defaulting to 3000")
		return DefaultPort
	}

	return portNumber
}

func main() {
	
	http.HandleFunc("/", router.HelloHandler)
	http.HandleFunc("/movies", router.ListMovies)
	http.HandleFunc("/count-movies", router.CountMovies)
	http.HandleFunc("/csv-movies", router.ListCSVMovies)

	port := readPortConfig()
	fmt.Printf("Starting server at port %d \n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
