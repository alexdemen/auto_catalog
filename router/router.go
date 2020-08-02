package router

import (
	"github.com/alexdemen/auto_catalog/middleware"
	"github.com/alexdemen/auto_catalog/model"
	"github.com/alexdemen/auto_catalog/storage"
	"github.com/google/jsonapi"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
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
		HandlerFunc(middleware.PathParams(h.getCar, "/api/v1/cars/{id}"))
	h.Router.Path("/api/v1/cars").Methods("POST").HandlerFunc(h.addCar)
	h.Router.Path("/api/v1/cars/{id}").
		Methods("PUT").
		HandlerFunc(middleware.PathParams(h.updateCar, "/api/v1/cars/{id}"))
	h.Router.Path("/api/v1/cars/{id}").
		Methods("DELETE").
		HandlerFunc(middleware.PathParams(h.deleteCar, "/api/v1/cars/{id}"))
}

func (h *Handler) carsList(writer http.ResponseWriter, _ *http.Request) {
	data, err := h.storage.GetCars()
	if err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", jsonapi.MediaType)
	err = jsonapi.MarshalPayload(writer, data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) getCar(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id")
	intId, err := strconv.Atoi(id.(string))
	if err != nil {
		http.Error(writer, "incorrect id", http.StatusBadRequest)
		return
	}

	data, err := h.storage.GetCar(intId)
	if err != nil {
		http.NotFound(writer, request)
	}
	writer.Header().Set("Content-Type", jsonapi.MediaType)
	err = jsonapi.MarshalPayload(writer, data)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *Handler) addCar(w http.ResponseWriter, r *http.Request) {
	car := new(model.Car)
	err := jsonapi.UnmarshalPayload(r.Body, car)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	if validation := car.Validate(); validation != "" {
		http.Error(w, validation, http.StatusBadRequest)
		return
	}

	err = h.storage.AddCar(car)

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
	}

}

func (h *Handler) updateCar(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id")
	intId, err := strconv.Atoi(id.(string))
	if err != nil {
		http.Error(writer, "incorrect id", http.StatusBadRequest)
		return
	}

	car := new(model.Car)
	err = jsonapi.UnmarshalPayload(request.Body, car)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	car.ID = intId
	err = h.storage.Update(car)
	if err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
	}
}

func (h *Handler) deleteCar(writer http.ResponseWriter, request *http.Request) {
	id := request.Context().Value("id")
	intId, err := strconv.Atoi(id.(string))
	if err != nil {
		http.Error(writer, "incorrect id", http.StatusBadRequest)
		return
	}

	err = h.storage.DeleteCar(intId)
	if err != nil {
		http.Error(writer, "", http.StatusInternalServerError)
		return
	}
}
