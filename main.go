package main

import (
	"net/http"

	"./API"
	"github.com/gorilla/mux"
)

var (
	ServerPort = ":8080"
)

/*
end point for entering a/multple game results
check if the request json doesn`t have ant errors and
then isert the data to the
*/

func insertDataEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonBody, _ := insertdata.InsertData(r)
	w.Write(jsonBody)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/insert", insertDataEndpoint).Methods("POST")
	http.ListenAndServe(ServerPort, router)
}
