package api

import (
	"log"
	"net/http"
	"os"

	"github.com/ComputePractice2017/adressbook-server/model"
	"github.com/gorilla/handlers"
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
	r.HandleFunc("/persons", firstOptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/persons/{guid}", editPersonHandler).Methods("PUT")
	r.HandleFunc("/persons/{guid}", deletePersonHandler).Methods("DELETE")
	r.HandleFunc("/persons/{guid}", secondOptionsHandler).Methods("OPTIONS")

	log.Println("Running the server on port 8000...")

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r))
}
