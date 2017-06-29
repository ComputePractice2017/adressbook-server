package api

import (
	"log"
	"net/http"

	"github.com/ComputePractice2017/adressbook-server/model"
	"github.com/gorilla/mux"
)

//Run runs the server
func Run() {
	log.Println("Connecting to rethinkDB on localhost...")
	err := model.InitSesson()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected")

	r := mux.NewRouter()
	r.HandleFunc("/", helloWorldHandler).Methods("GET")
	r.HandleFunc("/persons", getAllPersonsHandler).Methods("GET")
	r.HandleFunc("/persons", newPersonHandler).Methods("POST")
	r.HandleFunc("/persons/{guid}", editPersonHandler).Methods("PUT")

	log.Println("Running the server on port 8000...")
	http.ListenAndServe(":8000", r)
}
