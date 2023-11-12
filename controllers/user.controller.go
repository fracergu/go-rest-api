package controllers

import (
	"encoding/json"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a new user
// @Summary Creates a new user
// @Description creates a new user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user  body    models.UserCreateModel    true  "User's info"
// @Success 200 {object}  models.UserViewModel
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUser models.UserCreateModel
	err := json.NewDecoder(r.Body).Decode(&createUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Password hashing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al codificar la contraseña", http.StatusInternalServerError)
		return
	}

	// CCreate user model
	user := models.UserModel{
		Name:     createUser.Name,
		Surname:  createUser.Surname,
		Username: createUser.Username,
		Email:    createUser.Email,
		Password: string(hashedPassword),
	}

	// Save user in database
	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	// Convert to UserViewModel before sending the response
	userView := models.UserViewModel{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Username: user.Username,
		Email:    user.Email,
	}

	json.NewEncoder(w).Encode(userView)
}

// CreateUser returns all users
// @Summary Returns all users
// @Description Returns all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array}  models.UserViewModel
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.UserModel
	database.DB.Find(&users)

	// Convert to UserViewModel before sending the response
	var usersView []models.UserViewModel
	for _, user := range users {
		userView := models.UserViewModel{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Name:      user.Name,
			Surname:   user.Surname,
			Username:  user.Username,
			Email:     user.Email,
		}
		usersView = append(usersView, userView)
	}

	// Return empty array if no users found
	if len(usersView) == 0 {
		usersView = []models.UserViewModel{}
	}

	json.NewEncoder(w).Encode(usersView)
}

// GetUser returns a user
// @Summary Returns a user
// @Description returns a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id      path    int     true  "User ID"
// @Success 200 {object}  models.UserViewModel
// @Failure 404
// @Router /users/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UserModel
	database.DB.First(&user, id)

	// Return 404 if user not found
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Convert to UserViewModel before sending the response
	userView := models.UserViewModel{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Surname:   user.Surname,
		Username:  user.Username,
		Email:     user.Email,
	}

	json.NewEncoder(w).Encode(userView)
}

// UpdateUser updates a user
// @Summary Updates a user
// @Description updates a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id      path    int     true  "User ID"
// @Param   user  body    models.UserUpdateModel    true  "User's info"
// @Success 200 {object}  models.UserViewModel
// @Failure 404
// @Router /users/{id} [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Get user from database
	var user models.UserModel
	database.DB.First(&user, id)

	// Return 404 if user not found
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Update user with new data
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)

	// Convert to UserViewModel before sending the response
	userView := models.UserViewModel{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Surname:   user.Surname,
		Username:  user.Username,
		Email:     user.Email,
	}

	json.NewEncoder(w).Encode(userView)
}

// DeleteUser deletes a user
// @Summary Deletes a user
// @Description deletes a user with the input payload
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id      path    int     true  "User ID"
// @Success 204
// @Failure 404
// @Router /users/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UserModel
	database.DB.Delete(&user, id)

	w.WriteHeader(http.StatusNoContent)
}
