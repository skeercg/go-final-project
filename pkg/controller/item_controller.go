package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-final-project/pkg/model"
	"log"
	"net/http"
	"strconv"
)

type searchParams struct {
	Name string `json:"name" binding:"required"`
	Sort  string `json:"sort" binding:"required"`
}
type filterParams struct {
	Sort  string `json:"sort" binding:"required"`
}

func (c *Controller) getItems(w http.ResponseWriter, r *http.Request) {
	var params searchParams
	err := json.NewDecoder(r.Body).Decode(&params)

	items, err := c.services.Items.GetAll(params.Name, params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(items)
}

func (c *Controller) getItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	item, err := c.services.Items.GetById(id)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = json.NewEncoder(w).Encode(item)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func (c *Controller) deleteItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Items.Delete(id)
	if err != nil {
		return
	}
}

func (c *Controller) createItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Items.Create(item)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}

func (c *Controller) updateItemById(w http.ResponseWriter, r *http.Request) {

	var item model.Item
	err := json.NewDecoder(r.Body).Decode(&item)

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Items.Update(item, id)

	if err != nil {
		w.WriteHeader(400)
		return
	}
}


func (c *Controller) FilterbyRating(w http.ResponseWriter, r *http.Request) {
	var params filterParams
	err := json.NewDecoder(r.Body).Decode(&params)

	items, err := c.services.Items.FilterbyRating(params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(items)
}

func (c *Controller) FilterbyPrice(w http.ResponseWriter, r *http.Request) {
	var params filterParams
	err := json.NewDecoder(r.Body).Decode(&params)

	items, err := c.services.Items.FilterbyPrice(params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(items)
}