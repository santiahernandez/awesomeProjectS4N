package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func apiStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Classroom management ms working")
}

func HandleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//test
	myRouter.HandleFunc("/", apiStatus).Methods("GET")

	//basic crud
	myRouter.HandleFunc("/", createUser).Methods("POST")

	fmt.Println("Port 8080 is listening")
	log.Fatal(http.ListenAndServe(":8080", myRouter))

}
