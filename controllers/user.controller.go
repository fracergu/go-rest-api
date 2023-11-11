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
// @Param   user  body    models.UserModel    true  "User's info"
// @Success 200 {object}  models.UserModel
// @Router /users [post]
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserModel
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Password encryption
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error al codificar la contraseña", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// CreateUser returns all users
// @Summary Returns all users
// @Description Returns all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array}  models.UserModel
// @Router /users [get]
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.UserModel
	database.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UserModel
	database.DB.First(&user, id)

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UserModel
	database.DB.First(&user, id)

	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user models.UserModel
	database.DB.Delete(&user, id)

	w.WriteHeader(http.StatusNoContent)
}
