package router

import (
	"github.com/alexdemen/auto_catalog/middleware"
	"github.com/alexdemen/auto_catalog/model"
	"github.com/alexdemen/auto_catalog/storage"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	*mux.Router
	storage storage.Storable
}

func NewHandler(s storage.Storable) *Handler {
	handler := Handler{storage: s}
	handler.configureRouter()
	return &handler
}

func (h *Handler) configureRouter() {
	h.Router = mux.NewRouter()
	h.Router.Path("/api/v1/cars").Methods("GET").HandlerFunc(h.carsList)
	h.Router.Path("/api/v1/cars/{id}").
		Methods("GET").
		HandlerFunc(middleware.Identity(h.getCar, "/api/v1/cars/{id}"))
	h.Router.Path("/api/v1/cars").Methods("POST").HandlerFunc(h.addCar)
	h.Router.Path("/api/v1/cars/{id}").Methods("PATCH").HandlerFunc(h.updateCar)
}

func (h *Handler) carsList(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", jsonapi.MediaType)
	err := jsonapi.MarshalPayload(writer, h.storage.GetCars())
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getCar(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id")
	print(id)
	writer.Header().Set("Content-Type", jsonapi.MediaType)
	err := jsonapi.MarshalPayload(writer, h.storage.GetCars()[0])
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) addCar(w http.ResponseWriter, r *http.Request) {
	car := new(model.Car)
	err := jsonapi.UnmarshalPayload(r.Body, car)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func (h *Handler) updateCar(writer http.ResponseWriter, request *http.Request) {

}
