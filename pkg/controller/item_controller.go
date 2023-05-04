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
	Sort string `json:"sort" binding:"required"`
}

type filterParams struct {
	Sort string `json:"sort" binding:"required"`
}

type gradeParams struct {
	Rating float32 `json:"rating"`
}

func (c *Controller) purchaseItem(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("userId").(int)

	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])

	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = c.services.Items.Purchase(productId, userId)

	if err != nil {
		w.WriteHeader(500)
		return
	}
}

func (c *Controller) gradeItem(w http.ResponseWriter, r *http.Request) {
	var params gradeParams
	err := json.NewDecoder(r.Body).Decode(&params)

	if err != nil || params.Rating > 5.0 {
		w.WriteHeader(400)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	err = c.services.Items.Grade(id, params.Rating)

	if err != nil {
		w.WriteHeader(500)
		return
	}
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

func (c *Controller) FilterByRating(w http.ResponseWriter, r *http.Request) {
	var params filterParams
	err := json.NewDecoder(r.Body).Decode(&params)

	items, err := c.services.Items.FilterByRating(params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(items)
}

func (c *Controller) FilterByPrice(w http.ResponseWriter, r *http.Request) {
	var params filterParams
	err := json.NewDecoder(r.Body).Decode(&params)

	items, err := c.services.Items.FilterByPrice(params.Sort)
	if err != nil {
		log.Print(err)
	}

	err = json.NewEncoder(w).Encode(items)
}
