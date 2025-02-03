package controller

import (
	"encoding/json"
	"golang_starter_template/pkg/jobs/entity"
	"golang_starter_template/pkg/jobs/service"
	"golang_starter_template/pkg/utils"
	"net/http"
)

type UserController struct {
	UserService service.UserService
}

func (uc *UserController) UserRoutes(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("POST "+utils.GetEnv("API_PREFIX")+"/signup", uc.Signup)
	return mux
}

func (uc *UserController) Signup(w http.ResponseWriter, r *http.Request) {
	user := new(entity.User)
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.LoggerError.Println(utils.Error+"Invalid Request Body", err, utils.Reset)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	if err := uc.UserService.CreateUser(user); err != nil {
		utils.LoggerError.Println(utils.Error+"Failed to create user", err, utils.Reset)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusCreated,
		"message": "User created successfully",
	})
	if err != nil {
		utils.LoggerError.Println(utils.Error+"Failed to encode response", err, utils.Reset)
		return
	}
	utils.LoggerInfo.Println(utils.Info+"User created successfully", utils.Reset)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.LoggerError.Println(utils.Error+"Invalid Request Body", err, utils.Reset)
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	// Check if user exists

}
