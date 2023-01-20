package router

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/marceloemanoel/movieDB/db"
	"go.mongodb.org/mongo-driver/bson"
)

func ListMovies(w http.ResponseWriter, r *http.Request) {
	client, error := db.Init()
	if error != nil {
		log.Fatal(error)
	}

	collection := client.Database("movies").Collection("movies")
	// bson.D{{}} selects all documents
	cursor, err := collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		w.WriteHeader(500)
		return
	}

	results := []*db.Movie{}
	if error := cursor.All(context.TODO(), &results); error != nil {
		fmt.Fprint(w, error)
	}

	json.NewEncoder(w).Encode(results)
}

func ListCSVMovies(w http.ResponseWriter, r *http.Request) {
	results, err := decode("rotten_tomatoes_top_movies.csv", db.ParseMovie)

	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	json.NewEncoder(w).Encode(results)
}

func CountMovies(w http.ResponseWriter, req *http.Request) {
	client, error := db.Init()
	if error != nil {
		log.Fatal(error)
	}

	collection := client.Database("movies").Collection("movies")

	// bson.D{{}} selects all documents
	count, err := collection.CountDocuments(context.TODO(), bson.D{{}})
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, error)
		return
	}

	json.NewEncoder(w).Encode(count)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("testing log!")
	fmt.Fprintf(w, "Hello from router package!")
}

type Any interface{}

type Decoder[T Any] func([]string) *T

func decode[T Any](fileName string, parse Decoder[T]) ([]*T, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	results := []*T{}
	for line, record := range records {
		// skip csv header line
		if line == 0 {
			continue
		}

		value := parse(record)
		results = append(results, value)
	}

	return results, nil
}
