package main

import (
	"net/http"
	"sample/src/config"
	"sample/src/controller"
	"sample/src/helper"
	"sample/src/model"
	"sample/src/repository"
	"sample/src/router"
	"sample/src/service"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msg("Starting server...")
	db := config.DatabaseConnection()
	validate := validator.New()
	model.Migrate()

	repository := repository.NewUserRepositoryImpl(db)

	service := service.NewUserServiceImpl(repository, validate)

	controller := controller.NewUserController(service)

	routes := router.NewRouter(controller)

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.CheckErr(err)
}
