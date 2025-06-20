package app

import (
	"log"
	"net/http"
	"os"

	"github.com/ewatgar/size-calc/internal/handler"
)

type app struct {
	name   string
	addr   string
	router *http.ServeMux
	log    *log.Logger
}

var App = app{
	"size-calc",
	":8000",
	http.NewServeMux(),
	log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime|log.Lshortfile),
}

func Run() {
	App.log.Println("Run App")
	ConnectServer()
}

func ConnectServer() {
	App.log.Println("Start server")
	router := App.router

	//Brands
	router.HandleFunc("GET /brands/", handler.GetBrands)
	router.HandleFunc("POST /brands/", handler.CreateBrand)
	router.HandleFunc("GET /brands/{id}", handler.GetBrand)
	router.HandleFunc("DELETE /brands/{id}", handler.DeleteBrand)

	//Size
	router.HandleFunc("GET /sizes/", handler.GetSizes)
	router.HandleFunc("POST /sizes/", handler.CreateSize)
	router.HandleFunc("GET /sizes/{id}", handler.GetSize)
	router.HandleFunc("DELETE /sizes/{id}", handler.DeleteSize)

	//Code
	router.HandleFunc("GET /codes/", handler.GetCodes)
	router.HandleFunc("POST /codes/", handler.CreateCode)
	router.HandleFunc("GET /codes/{id}", handler.GetCode)
	router.HandleFunc("DELETE /codes/{id}", handler.DeleteCode)

	//Users
	router.HandleFunc("GET /users/", handler.GetUsers)
	router.HandleFunc("POST /users/", handler.CreateSize)
	router.HandleFunc("GET /users/{id}/", handler.GetUser)
	router.HandleFunc("DELETE /users/{id}/", handler.DeleteUser)

	//Measurements
	router.HandleFunc("GET /measurements/", handler.GetMeasurements)
	router.HandleFunc("POST /measurements/", handler.CreateMeasurement)
	router.HandleFunc("GET /measurements/{id}/", handler.GetMeasurement)
	router.HandleFunc("DELETE /measurements/{id}/", handler.DeleteMeasurement)

	err := http.ListenAndServe(App.addr, App.router)
	if err != nil {
		App.log.Fatal(err)
	}
}
