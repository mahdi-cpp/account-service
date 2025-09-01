package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/account-service/internal/account"
	"github.com/mahdi-cpp/account-service/internal/user"
)

type AccountHandler struct {
	manager *account.Manager
}

func NewAccountHandler(manager *account.Manager) *AccountHandler {
	return &AccountHandler{
		manager: manager,
	}
}

type requestBody struct {
	UserID string `json:"userID"`
}

func (handler *AccountHandler) Create(c *gin.Context) {

	var request user.Update
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	//userStorage, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}

	newItem, err := handler.manager.Create(&user.User{
		Username:    request.Username,
		FirstName:   request.FirstName,
		LastName:    request.LastName,
		AvatarURL:   request.AvatarURL,
		PhoneNumber: request.PhoneNumber,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//update := asset.AssetUpdate{AssetIds: request.AssetIds, AddAlbums: []int{newItem.ID}}
	//_, err = userStorage.UpdateAsset(update)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//userStorage.UpdateCollections()

	c.JSON(http.StatusCreated, newItem)
}

func (handler *AccountHandler) Update(c *gin.Context) {

	//userID, err := utils.GetUserId(c)
	//if err != nil {
	//	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	var itemUpdate user.Update

	if err := c.ShouldBindJSON(&itemUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	update, err := handler.manager.Update(itemUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, update)
}

func (handler *AccountHandler) GetCollectionList(c *gin.Context) {

	item2, err := handler.manager.UserCollection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, item2)
}

func (handler *AccountHandler) GetUserByID(c *gin.Context) {

	var request requestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := handler.manager.UserCollection.Get(request.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	//result := asset.PHCollectionList[*account.User]{
	//	Collections: make([]*asset.PHCollection[*account.User], len(items)),
	//}
	//
	//for i, item := range items {
	//	assets, _ := handler.manager.AccountManager.GetItemAssets(item.ID)
	//	result.Collections[i] = &asset.PHCollection[*account.User]{
	//		Item:   item,
	//		Assets: assets,
	//	}
	//}

	c.JSON(http.StatusOK, user)
}

func (handler *AccountHandler) GetUser(c *gin.Context) {

	item, err := handler.manager.UserCollection.Get("0198adfd-c0ca-7151-990f-b50956fc7f27")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (handler *AccountHandler) GetList(c *gin.Context) {

	//var with asset.PHFetchOptions
	//if err := c.ShouldBindJSON(&with); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	//	fmt.Println("Invalid request")
	//	return
	//}

	//userStorage, err := handler.manager.GetUserManager(c, userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//}

	//items, err := handler.manager.AccountManager.GetAllSorted(with.SortBy, with.SortOrder)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}

	items, err := handler.manager.UserCollection.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	//result := asset.PHCollectionList[*account.User]{
	//	Collections: make([]*asset.PHCollection[*account.User], len(items)),
	//}
	//
	//for i, item := range items {
	//	assets, _ := handler.manager.AccountManager.GetItemAssets(item.ID)
	//	result.Collections[i] = &asset.PHCollection[*account.User]{
	//		Item:   item,
	//		Assets: assets,
	//	}
	//}

	c.JSON(http.StatusOK, items)
}

func (handler *AccountHandler) Delete(c *gin.Context) {

	//err = handler.manager.UserCollection.Delete(userID)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err})
	//	return
	//}
	//
	//c.JSON(http.StatusCreated, "delete ok")
}
