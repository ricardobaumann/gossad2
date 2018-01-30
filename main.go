package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/ricardobaumann/gossad2/piwik"
)

var (
	repo DynamoRepo
)

func schedule(what func(), delay time.Duration) chan bool {
	stop := make(chan bool)

	go func() {
		for {
			what()
			select {
			case <-time.After(delay):
			case <-stop:
				return
			}
		}
	}()

	return stop
}

func extractFromPiwik() {
	os.Setenv("piwikToken", "325b6226f6b06472e78e6da694999486")
	os.Setenv("limitPercontentID", "10")
	os.Setenv("limitPerPage", "500")
	os.Setenv("maxPages", "1")
	os.Setenv("throughput", "1")

	results := piwik.GetRecommendations()
	repo.Init()
	for k, v := range results {
		bytes, _ := json.Marshal(v)
		repo.PutItem(k, string(bytes))
	}

	println("Finished\n\n")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/{id}", IDHandler)
	http.Handle("/", r)
	repo = DynamoRepo{
		Region:       "eu-west-1",
		TableName:    "test-gossad",
		IDColumnName: "id",
	}

	stop := schedule(extractFromPiwik, 15*time.Minute)
	defer func() {
		stop <- true
	}()

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
