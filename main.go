package main

import (
	"fmt"
	"github.com/ameniGa/timeTracker/config"
	"github.com/ameniGa/timeTracker/database"
	"github.com/ameniGa/timeTracker/helpers"
	"github.com/ameniGa/timeTracker/server/http/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	//faceRecognition.Register()
	r := mux.NewRouter()
	conf, err := config.LoadConfig()
	if err != nil {
		log.Panic(err)
	}
	logger := helpers.GetLogger()
	db, err := database.Create(&conf.Database, conf.Server.Deadline, logger)
	if err != nil {
		log.Panic(err)
	}
	runner := services.Runner{Config: conf, Db: db}
	r.HandleFunc(fmt.Sprintf("/api/auth"), runner.GetUser).Methods(http.MethodGet)
	r.HandleFunc(fmt.Sprintf("/api/user"), runner.UpdateUser).Methods(http.MethodPut)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", conf.Server.Host, conf.Server.Port), r))
}
