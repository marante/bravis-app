package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/marante/bravis-app/models"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func Index(w http.ResponseWriter, r *http.Request) *appError {
	file, err := http.Get("http://localhost:8080/assets/css/styles.css")
	if err != nil {
		panic(err)
	}
	defer file.Body.Close()
	data, err := ioutil.ReadAll(file.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data from the route", string(data))
	return nil
}

func Workorder(w http.ResponseWriter, r *http.Request) *appError {
	workorders := models.Workorder{
		ObjNr:       "B400",
		Adress:      "Pilevallsvägen 19",
		Description: "Vi bygger grejer",
		Start:       "Påbörjat",
		Status:      "IDK",
		Invoice:     "Faktureras inte",
		Worker: []models.Worker{
			models.Worker{
				Name:        "Johan",
				HoursWorked: 45,
				Description: "Rev vägg. Målade vägg.",
			},
			models.Worker{
				Name:        "Johan",
				HoursWorked: 45,
				Description: "Rev vägg. Målade vägg.",
			},
		},
	}
	tpl.ExecuteTemplate(w, "index.gohtml", workorders)
	return nil
}

func main() {
	// Get port from env variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	// Initializing a gorilla mux router.
	r := mux.NewRouter()

	// Creating a fileserver to serve my assets.
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	r.Handle("/", appHandler(Index)).Methods("GET")
	r.Handle("/workorder", appHandler(Workorder)).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r)))
}
