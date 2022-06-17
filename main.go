package main

import (
	"net/http"

	controller "kwai/controller/user"
	repository "kwai/repository/user"
	service "kwai/service/user"

	"github.com/gorilla/mux"
)

var (
	listUserRepository repository.ListUserRepository = repository.NewListUserRepository()
	listUserService    service.ListUserService       = service.NewListUserService(listUserRepository)
	listUserController controller.ListUserController = controller.NewListUserController(listUserService)
)

func main() {
	const port string = ":8081"

	r := mux.NewRouter()

	r.HandleFunc("/users", listUserController.ListUsers).Methods("GET")
	r.HandleFunc("/users/{username}", listUserController.FindByName).Methods("POST")

	http.ListenAndServe(port, r)
}
