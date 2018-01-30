package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	repo DynamoRepo
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{id}", IDHandler)
	http.Handle("/", r)
	repo = DynamoRepo{
		Region:       "eu-west-1",
		TableName:    "test-gossad",
		IDColumnName: "id",
	}

	//fmt.Printf("%v", value)
	http.ListenAndServe(":5000", r)
}

//IDHandler Handles per id requests
func IDHandler(w http.ResponseWriter, r *http.Request) {

	value := repo.GetItem(mux.Vars(r)["id"])
	if value != "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, value)
	} else {
		fmt.Fprintf(w, "[]")
	}
}
