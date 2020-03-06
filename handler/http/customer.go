package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/zlp-ecommerce/customer-service/driver"
	models "github.com/zlp-ecommerce/customer-service/models"
	repository "github.com/zlp-ecommerce/customer-service/repository"
	"github.com/zlp-ecommerce/customer-service/repository/customer"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Customer {
	return &Customer{
		repo: customer.NewSQLCustomerRepo(db.SQL),
	}
}

// Customers ...
type Customer struct {
	repo repository.CustomerRepo
}

// Fetch all customer data
func (p *Customer) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}

// Create a new customer
func (p *Customer) Create(w http.ResponseWriter, r *http.Request) {
	customer := models.Customer{}
	json.NewDecoder(r.Body).Decode(&customer)

	newID, err := p.repo.Create(r.Context(), &customer)
	fmt.Println(newID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully Created"})
}

// Update a customer by id
func (p *Customer) Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Customer{ID: int64(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a customer details
func (p *Customer) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// Delete a customer
func (p *Customer) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Server Error")
	}

	respondwithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete Successfully"})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
