package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"html/template"
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
	fmt.Fprintln(w, "Yo dude! Sup?")
	return nil
}

func WorkOrder(w http.ResponseWriter, r *http.Request) *appError {
	//workOrders := models.Workorder{
	//	ObjNr:       "B400",
	//	Adress:      "Pilevallsvägen 19",
	//	Description: "Vi bygger grejer",
	//	Start:       "Påbörjat",
	//	Status:      "IDK",
	//	Invoice:     "Faktureras inte",
	//	Worker: []models.Worker{
	//		{
	//			Name:        "Johan",
	//			HoursWorked: 45,
	//			Description: "Rev vägg. Målade vägg.",
	//		},
	//		{
	//			Name:        "Johan",
	//			HoursWorked: 45,
	//			Description: "Rev vägg. Målade vägg.",
	//		},
	//	},
	//}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
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
	r.Handle("/workorder", appHandler(WorkOrder)).Methods("GET")
	r.Handle("/favicon.ico", r.NotFoundHandler)
	log.Fatal(http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r)))
}
