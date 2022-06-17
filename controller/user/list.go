package controller

import (
	"net/http"

	service "kwai/service/user"

	"github.com/gorilla/mux"
)

type (
	listUserController struct{}
	ListUserController interface {
		ListUsers(rw http.ResponseWriter, r *http.Request)
		FindByName(rw http.ResponseWriter, r *http.Request)
	}
)

var listUserService service.ListUserService

func NewListUserController(service service.ListUserService) ListUserController {
	listUserService = service
	return &listUserController{}
}

func (*listUserController) ListUsers(rw http.ResponseWriter, r *http.Request) {
	listUserService.ListUsers()
}

func (*listUserController) FindByName(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	listUserService.FindByName(vars["username"])
}
