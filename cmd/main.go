package main

import (
	"fmt"
	"log"

	"github.com/mahdi-cpp/account-service/account"
)

func main() {

	ginInit()

	// Create account manager
	manager, err := account.NewAccountManager()
	if err != nil {
		log.Fatalf("failed to create account manager: %v", err)
		return
	}
	defer func(manager *account.ServiceManager) {
		err := manager.Close()
		if err != nil {

		}
	}(manager) // Ensure proper cleanup

	users, err := manager.UserCollection.GetAll()
	if err != nil {
		return
	}

	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName)
	}

	err = manager.Publish()
	if err != nil {
		log.Fatalf("failed to publish manager: %v", err)
		return
	}

	h := account.NewUserHandler(manager)
	userRoute(h)

	startServer(router)

}

func userRoute(h *account.UserHandler) {

	api := router.Group("/api/v1/user")

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("get_user", h.GetUser)
	api.POST("list", h.GetList)
}
