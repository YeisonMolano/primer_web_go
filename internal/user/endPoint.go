package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Controller func(response http.ResponseWriter, request *http.Request)

	EndPoints struct {
		Create Controller
		Get    Controller
		GetAll Controller
		Updte  Controller
		Delete Controller
	}

	CreateRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Telefono  string `json:"telefono"`
	}

	UpdateRequest struct {
		FirstName *string `json:"firstName"`
		LastName  *string `json:"lastName"`
		Email     *string `json:"email"`
		Telefono  *string `json:"telefono"`
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Err    string      `json:"error,omitempty"`
	}
)

func MakeEndpoints(service Service) EndPoints {
	return EndPoints{
		Create: makeCreateEndPoint(service),
		Get:    makeGetEndPoint(service),
		GetAll: makeGetAllEndPoint(service),
		Updte:  makeUpdateEndPoint(service),
		Delete: makeDeleteEndPoint(service),
	}
}

func makeCreateEndPoint(service Service) Controller {
	return func(response http.ResponseWriter, request *http.Request) {
		var req CreateRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "Error al recibir el dato JSON"})
			return
		}
		if req.FirstName == "" {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "El campo First name no puede estar vacío"})
			return
		}

		if req.LastName == "" {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "El campo Last name no puede estar vacío"})
			return
		}
		user, err := service.Create(req.FirstName, req.LastName, req.Email, req.Telefono)
		if err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "Error al crear el servicio"})
			return
		}
		json.NewEncoder(response).Encode(&Response{Status: 200, Data: user})
	}
}

func makeGetEndPoint(service Service) Controller {
	return func(response http.ResponseWriter, request *http.Request) {
		path := mux.Vars(request)
		id := path["id"]
		user, err := service.GetById(id)
		if err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}
		json.NewEncoder(response).Encode(user)
	}
}

func makeGetAllEndPoint(service Service) Controller {
	return func(response http.ResponseWriter, request *http.Request) {
		values := request.URL.Query()

		filter := filter{
			First_name: values.Get("firstName"),
			Last_name:  values.Get("lastName"),
		}

		users, err := service.GetAll(filter)
		if err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}
		json.NewEncoder(response).Encode(&Response{Status: 200, Data: users})
	}
}

func makeUpdateEndPoint(service Service) Controller {
	return func(response http.ResponseWriter, request *http.Request) {
		var req UpdateRequest
		if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "Error al recibir el dato JSON"})
			return
		}
		if req.FirstName != nil && *req.FirstName == "" {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "El campo First name no puede estar vacío"})
			return
		}

		if req.LastName != nil && *req.LastName == "" {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: "El campo Last name no puede estar vacío"})
			return
		}
		path := mux.Vars(request)
		id := path["id"]
		err := service.Update(id, req.FirstName, req.LastName, req.Email, req.Telefono)
		log.Println("LLego")
		if err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}
		json.NewEncoder(response).Encode(&Response{Status: 200, Data: req})
	}
}

func makeDeleteEndPoint(service Service) Controller {
	return func(response http.ResponseWriter, request *http.Request) {
		path := mux.Vars(request)
		id := path["id"]
		value, err := service.Delete(id)
		if err != nil {
			response.WriteHeader(400)
			json.NewEncoder(response).Encode(&Response{Status: 400, Err: err.Error()})
			return
		}
		json.NewEncoder(response).Encode(&Response{Status: 200, Data: value})
	}
}
