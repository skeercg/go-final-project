package controller

import (
	"go-final-project/pkg/service"

	"github.com/gorilla/mux"
)

type Controller struct {
	services *service.Service
}

func NewController(services *service.Service) *Controller {
	return &Controller{services: services}
}

func (c *Controller) InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/sign-in", c.signIn).Methods("POST")

	auth.HandleFunc("/sign-up", c.signUp).Methods("POST")

	return router
}
