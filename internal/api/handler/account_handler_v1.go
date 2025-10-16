package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/mahdi-cpp/account-service/internal/application"
	"github.com/mahdi-cpp/account-service/internal/collections/user"
)

type AccountHandler struct {
	appManager *application.AppManager
}

func NewAccountHandler(appManager *application.AppManager) *AccountHandler {
	return &AccountHandler{
		appManager: appManager,
	}
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func SendError(c *gin.Context, message string, code int) {
	c.JSON(http.StatusBadRequest, gin.H{"message": message, "code": code})
}

func (handler *AccountHandler) Create(c *gin.Context) {

	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		SendError(c, "Invalid with", http.StatusBadRequest)
		return
	}

	newItem, err := handler.appManager.UserManager.Create(&u)
	if err != nil {
		SendError(c, err.Error(), 2)
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func (handler *AccountHandler) Read(c *gin.Context) {

	fmt.Println("Read...")

	idStr := c.Query("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	item, err := handler.appManager.UserManager.Read(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (handler *AccountHandler) ReadAll(c *gin.Context) {

	var with *user.SearchOptions
	err := json.NewDecoder(c.Request.Body).Decode(&with)
	if err != nil {
		SendError(c, err.Error(), http.StatusBadRequest)
		return
	}

	users, err := handler.appManager.UserManager.ReadAll(with)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	for _, u := range users {
		fmt.Println(u.ThumbnailURL)
	}

	c.JSON(http.StatusCreated, users)
}

func (handler *AccountHandler) Update(c *gin.Context) {

	fmt.Println("update")

	var with user.UpdateOptions
	if err := c.ShouldBindJSON(&with); err != nil {
		SendError(c, "Invalid with", http.StatusBadRequest)
		return
	}

	newItem, err := handler.appManager.UserManager.Update(with)
	if err != nil {
		SendError(c, err.Error(), 2)
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func (handler *AccountHandler) Delete(c *gin.Context) {

	//err = handler.appManager.UserCollection.Delete(userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}
	//
	//c.JSON(http.StatusCreated, "delete ok")
}
