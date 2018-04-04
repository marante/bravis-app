package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/marante/bravis-app/config"
	"github.com/marante/bravis-app/dao"
	"github.com/marante/bravis-app/models"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

var workorderDao = dao.WorkorderDao{}
var dbConfig = config.Config{}

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
	defer r.Body.Close()
	var order models.Workorder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to decode json body.", http.StatusBadRequest}
	}
	order.ID = bson.NewObjectId()
	if err := workorderDao.Insert(order); err != nil {
		fmt.Println(err)
		return &appError{err, "There was an error when trying to insert into database.", http.StatusBadRequest}
	}
	return nil
}

// AllWorkorders return all worksorders from the DB.
func AllWorkorders(w http.ResponseWriter, r *http.Request) *appError {
	workorders, err := workorderDao.FindAll()
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
	params := mux.Vars(r)
	order, err := workorderDao.FindById(params["id"])
	if err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to retrieve record from database", http.StatusBadRequest}
	}
	w.Header().Set("Content-type", "application/json")
	err = json.NewEncoder(w).Encode(order)
	return nil
}

// UpdateWorkorder updates a specific workorder by ID
func UpdateWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	defer r.Body.Close()
	var order models.Workorder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to decode json body.", http.StatusBadRequest}
	}
	if err := workorderDao.Update(order); err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to update workorder.", http.StatusBadRequest}
	}
	return nil
}

// DeleteWorkorder deletes a speicifc workorder by ID
func DeleteWorkorder(w http.ResponseWriter, r *http.Request) *appError {
	defer r.Body.Close()
	var order models.Workorder
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to decode json body.", http.StatusBadRequest}
	}
	if err := workorderDao.Delete(order); err != nil {
		fmt.Println(err)
		return &appError{err, "Error trying to delete workorder.", http.StatusBadRequest}
	}
	return nil
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	dbConfig.Read()
	workorderDao.Server = dbConfig.Server
	workorderDao.Database = dbConfig.Database
	workorderDao.Connect()
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
