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

	items := router.PathPrefix("/items").Subrouter()

	items.Use(c.userIdentity)

	items.HandleFunc("", c.getItems).Methods("GET")

	items.HandleFunc("/{id}", c.getItemById).Methods("GET")

	items.HandleFunc("/{id}", c.deleteItemById).Methods("DELETE")

	items.HandleFunc("/grade/{id}", c.gradeItem).Methods("POST")

	items.HandleFunc("/purchase/{id}", c.purchaseItem).Methods("POST")

	items.HandleFunc("", c.createItem).Methods("POST")

	items.HandleFunc("/{id}", c.updateItemById).Methods("PUT")

	items.HandleFunc("/sort/rating", c.FilterByRating).Methods("POST")

	items.HandleFunc("/sort/price", c.FilterByPrice).Methods("POST")

	return router
}
