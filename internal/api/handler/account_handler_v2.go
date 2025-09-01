package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Account represents the structure of a user account.
type Account struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Omit from JSON output
	CreatedAt time.Time `json:"createdAt"`
}

// In-memory database for demonstration purposes.
var accounts = make(map[string]Account)

// Request bodies for different endpoints.
type createAccountRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type updateAccountRequest struct {
	Email    string `json:"email,omitempty" binding:"omitempty,email"`
	Password string `json:"password,omitempty" binding:"omitempty,min=8"`
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Handlers for the API endpoints.

// createAccount godoc
// @Summary Create a new user account
// @Description Creates a new account with a unique ID and stores it.
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body createAccountRequest true "Account data"
// @Success 201 {object} map[string]string "Successful creation"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 409 {object} map[string]string "Duplicate email"
// @Router /api/v1/accounts [post]
func createAccount(c *gin.Context) {
	var req createAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON or missing fields."})
		return
	}

	for _, acc := range accounts {
		if acc.Email == req.Email {
			c.JSON(http.StatusConflict, gin.H{"error": "Email address already in use."})
			return
		}
	}

	id := uuid.New().String()
	newAccount := Account{
		ID:        id,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
	}

	accounts[id] = newAccount
	c.Header("Location", fmt.Sprintf("/api/v1/accounts/%s", id))
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// getAccount godoc
// @Summary Get account details by ID
// @Description Retrieves an account's details using a unique ID.
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 200 {object} Account "Successful retrieval"
// @Failure 400 {object} map[string]string "Invalid ID format"
// @Failure 404 {object} map[string]string "Account not found"
// @Router /api/v1/accounts/{id} [get]
func getAccount(c *gin.Context) {
	id := c.Param("id")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID format."})
		return
	}

	if account, ok := accounts[id]; ok {
		c.JSON(http.StatusOK, account)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found."})
}

// updateAccount godoc
// @Summary Update an account
// @Description Updates an existing account's details.
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "Account ID"
// @Param account body updateAccountRequest true "Account data to update"
// @Success 200 {object} Account "Successful update"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 404 {object} map[string]string "Account not found"
// @Router /api/v1/accounts/{id} [put]
func updateAccount(c *gin.Context) {
	id := c.Param("id")
	if _, ok := accounts[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Account not found."})
		return
	}

	var req updateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON or invalid fields."})
		return
	}

	existingAccount := accounts[id]
	if req.Email != "" {
		existingAccount.Email = req.Email
	}
	if req.Password != "" {
		existingAccount.Password = req.Password
	}

	accounts[id] = existingAccount
	c.JSON(http.StatusOK, existingAccount)
}

// deleteAccount godoc
// @Summary Delete an account by ID
// @Description Permanently deletes an account.
// @Tags accounts
// @Produce json
// @Param id path string true "Account ID"
// @Success 204 "No content"
// @Failure 404 {object} map[string]string "Account not found"
// @Router /api/v1/accounts/{id} [delete]
func deleteAccount(c *gin.Context) {
	id := c.Param("id")

	if _, ok := accounts[id]; ok {
		delete(accounts, id)
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Account not found."})
}

// login godoc
// @Summary Authenticate user
// @Description Logs in a user by checking credentials and returning a token.
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body loginRequest true "User credentials"
// @Success 200 {object} map[string]string "Successful login"
// @Failure 400 {object} map[string]string "Invalid input"
// @Failure 401 {object} map[string]string "Invalid credentials"
// @Router /api/v1/auth/login [post]
func login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON or missing fields."})
		return
	}

	for _, acc := range accounts {
		if acc.Email == req.Email && acc.Password == req.Password {
			c.JSON(http.StatusOK, gin.H{"token": "dummy-jwt-token"})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials."})
}
