package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	. "github.com/marante/bravis-app/config"
	. "github.com/marante/bravis-app/dao"
	"log"
	"net/http"
	"os"
)

var config = Config{}
var dao = WorkorderDao{}

// Below code simplifies and makes error handling for handlers more concrete.
type appError struct {
	Error   error
	Message string
	Code    int
}

// custom handlerfunc so I can return a custom error message.
type appHandler func(http.ResponseWriter, *http.Request) *appError

// definition of the custom ServeHTTP method.
func (ah appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := ah(w, r); err != nil {
		http.Error(w, err.Message, err.Code)
	}
}

// CreateWorkorder creates a new workorder
func CreateWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	return nil
}

// AllWorkorders return all worksorders from the DB.
func AllWorkorders(w http.ResponseWriter, r *http.Request) *appError {
	workorders, err := dao.FindAll()
	if err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to retrieve records from database", http.StatusBadRequest}
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(workorders)
	return nil
}

// FindWorkorder finds a specific workorder by ID
func FindWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	return nil
}

// UpdateWorkorder updates a specific workorder by ID
func UpdateWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	return nil
}

// DeleteWorkorder deletes a speicifc workorder by ID
func DeleteWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	return nil
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := mux.NewRouter()
	r.Handle("/workorders", appHandler(AllWorkorders)).Methods("GET")
	r.Handle("/workorders/{id}", appHandler(FindWorkorder)).Methods("GET")
	r.Handle("/workorders", appHandler(UpdateWorkorder)).Methods("PUT")
	r.Handle("/workorders", appHandler(CreateWorkorder)).Methods("POST")
	r.Handle("/workorders", appHandler(DeleteWorkorder)).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r)))
}
