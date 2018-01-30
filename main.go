package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	repo := DynamoRepo{
		Region:       "eu-west-1",
		TableName:    "test-gossad",
		IDColumnName: "id",
	}
	repo.GetItem("222")

	//fmt.Printf("%v", value)
	//http.ListenAndServe(":5000", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")
}
