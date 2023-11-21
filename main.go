package main

import (
	"net/http"
	"rianmz-genz/golang-restful-api/app"
	"rianmz-genz/golang-restful-api/controller"
	"rianmz-genz/golang-restful-api/helper"
	"rianmz-genz/golang-restful-api/middleware"
	"rianmz-genz/golang-restful-api/repository"
	"rianmz-genz/golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr: "localhost:7611",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
