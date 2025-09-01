package main

import (
	"fmt"
	"log"

	"github.com/mahdi-cpp/account-service/cmd/middleware"
	"github.com/mahdi-cpp/account-service/internal/account"
	"github.com/mahdi-cpp/account-service/internal/api/handler"
)

func main() {

	ginInit()

	// Create account manager
	manager, err := account.NewAccountManager()
	if err != nil {
		log.Fatalf("failed to create account manager: %v", err)
		return
	}
	defer func(manager *account.Manager) {
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

	h := handler.NewAccountHandler(manager)
	userRoute(h)

	startServer(router)

}

func userRoute(h *handler.AccountHandler) {

	api := router.Group("/api/v1/user")
	api.Use(middleware.AuthMiddleware())

	api.POST("create", h.Create)
	api.POST("update", h.Update)
	api.POST("delete", h.Delete)
	api.POST("get_user", h.GetUser)
	api.POST("list", h.GetList)
}

//
//func apiV2(h *api.AccountHandler) {
//
//	apiV2 := r.Group("/api/v2")
//	{
//		// Accounts Group
//		accountsGroup := apiV2.Group("/accounts")
//		{
//			accountsGroup.POST("", createAccount)
//			accountsGroup.GET("/:id", getAccount)
//			accountsGroup.PUT("/:id", updateAccount)
//			accountsGroup.DELETE("/:id", deleteAccount)
//		}
//
//		// Auth Group
//		authGroup := apiV2.Group("/auth")
//		{
//			authGroup.POST("/login", login)
//		}
//	}
//}
