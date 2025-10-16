package main

import (
	"log"

	"github.com/mahdi-cpp/account-service/internal/api/handler"
	"github.com/mahdi-cpp/account-service/internal/application"
)

func main() {

	Init()

	manager, err := application.New()
	if err != nil {
		log.Fatalf("failed to create application manager: %v", err)
		return
	}

	//defer func(manager *application.AppManager) {
	//	err := manager.Close()
	//	if err != nil {
	//
	//	}
	//}(manager) // Ensure proper cleanup
	//
	//err = manager.Publish()
	//if err != nil {
	//	log.Fatalf("failed to publish manager: %v", err)
	//	return
	//}

	h := handler.NewAccountHandler(manager)
	userRoute(h)

	StartServer(Router)
}

func userRoute(h *handler.AccountHandler) {

	api := Router.Group("")
	//api.Use(middleware.AuthMiddleware())

	api.POST("/api/users", h.Create)
	api.POST("/api/users/search", h.ReadAll)

	api.GET("/api/users", h.Read)

	api.PATCH("/api/users", h.Update)
	api.DELETE("/api/users", h.Delete)
}
