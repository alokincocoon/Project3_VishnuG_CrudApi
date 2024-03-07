package main

import (
	"log"
	"net/http"

	"github.com/edgedb/edgedb-go"
	"github.com/gorilla/mux"
	"github.com/vishnu-g-k/student-management/dbhandler"
	"github.com/vishnu-g-k/student-management/handlers"
)

type App struct {
	Router   *mux.Router
	DBClient *edgedb.Client
}

func (a *App) createOrListStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GetStudents(a.DBClient, w, r)
	} else if r.Method == "POST" {
		handlers.CreateStudent(a.DBClient, w, r)
	}
}

func (a *App) listOrUpdateStudents(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handlers.GetStudentDetails(a.DBClient, w, r)
	} else if r.Method == "PUT" {
		handlers.UpdateStudentDetails(a.DBClient, w, r)
	} else if r.Method == "DELETE" {
		handlers.DeleteStudentDetails(a.DBClient, w, r)
	}
}

func routerNotFound(w http.ResponseWriter, r *http.Request) {
	handlers.NotFoundResponse(w, r)
}

// func logMW(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Printf("%s - %s (%s)", r.Method, r.URL.Path, r.RemoteAddr)
// 		next.ServeHTTP(w, r)
// 	})
// }

// Routes
func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/students", a.createOrListStudents).Methods("GET", "POST")
	a.Router.HandleFunc("/students/{id}", a.listOrUpdateStudents).Methods("GET", "PUT", "DELETE")

	a.Router.NotFoundHandler = http.HandlerFunc(routerNotFound)
	// a.Router.Use(logMW)
}

func (a *App) Initialize() {
	dbClient, err := dbhandler.ConnectDb()
	a.DBClient = dbClient
	if err != nil {
		log.Println("Database failed to load, err: ", err)
	}
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(port string) {
	http.ListenAndServe(port, a.Router)
}
