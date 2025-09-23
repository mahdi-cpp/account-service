package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	newItem, err := handler.appManager.CreateUser(&u)
	if err != nil {
		SendError(c, err.Error(), 2)
		return
	}

	c.JSON(http.StatusCreated, newItem)
}

func (handler *AccountHandler) Read(c *gin.Context) {

	id, err := uuid.Parse("0198adfd-c0ca-7151-990f-b50956fc7f27")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}

	item, err := handler.appManager.ReadUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (handler *AccountHandler) ReadAll(c *gin.Context) {

	with := &user.SearchOptions{}

	item2, err := handler.appManager.ReadUsers(with)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *AccountHandler) Update(c *gin.Context) {

	fmt.Println("update")
	
	var with user.UpdateOptions
	if err := c.ShouldBindJSON(&with); err != nil {
		SendError(c, "Invalid with", http.StatusBadRequest)
		return
	}

	newItem, err := handler.appManager.UpdateUsers(&with)
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
