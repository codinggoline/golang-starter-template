package server

import (
	"errors"
	"fmt"
	"golang_starter_template/pkg/config/database"
	"golang_starter_template/pkg/global"
	"golang_starter_template/pkg/middleware"
	"golang_starter_template/pkg/routes"
	"golang_starter_template/pkg/utils"
	"log"
	"net/http"
)

func Start(tab []string) error {
	// Check if the number of arguments is correct
	if len(tab) != 0 {
		return errors.New("too many arguments")
	}

	// Read the .env file
	err := utils.Environment()
	if err != nil {
		return err
	}

	// Initialize the database
	global.DB, err = database.Connect()
	if err != nil {
		return err
	}
	defer global.DB.Close()

	// Migration
	if err := database.Migrate(global.DB.GetDB()); err != nil {
		return err
	}

	// Initialize the route
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	// Add the routes
	mux = routes.Routes(mux)

	// Add the middleware
	wrapperMux := middleware.LoggingMiddleware(mux)

	// Set the server structure
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", utils.GetEnv("PORT")),
		Handler: wrapperMux,
	}

	log.Default().Printf("%sThe server is listening at http://localhost%s%s\n", utils.Info, server.Addr, utils.Reset)
	utils.LoggerInfo.Printf("%sThe server is listening at http://localhost%s%s\n", utils.Info, server.Addr, utils.Reset)
	err = server.ListenAndServe()

	return err
}
