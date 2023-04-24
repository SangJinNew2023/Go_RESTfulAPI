package main

import (
	"CRUD02/config"
	"CRUD02/controller"
	"CRUD02/helper"
	"CRUD02/model"
	"CRUD02/repo"
	"CRUD02/router"
	"CRUD02/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Started Server")

	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	tagsRepository := repo.NewTagsRepositoryImpl(db)

	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	tagsController := controller.NewTagsController(tagsService)

	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8888",
		Handler: routes,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
