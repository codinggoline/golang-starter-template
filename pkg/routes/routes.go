package routes

import (
	"golang_starter_template/pkg/global"
	"golang_starter_template/pkg/jobs/controller"
	"golang_starter_template/pkg/jobs/repository"
	"golang_starter_template/pkg/jobs/service/impl"
	"net/http"
)

func Routes(routes *http.ServeMux) *http.ServeMux {
	// Repository
	userRepo := repository.NewUserRepoImpl(*global.DB)

	// Service
	userService := impl.UserServiceImpl{Repository: userRepo}

	// Controller
	userController := controller.UserController{UserService: &userService}

	// Routes
	routes = userController.UserRoutes(routes)

	return routes
}
