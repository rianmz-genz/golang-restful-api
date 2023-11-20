package main

import (
	"net/http"
	"rianmz-genz/golang-restful-api/app"
	"rianmz-genz/golang-restful-api/controller"
	"rianmz-genz/golang-restful-api/exception"
	"rianmz-genz/golang-restful-api/helper"
	"rianmz-genz/golang-restful-api/middleware"
	"rianmz-genz/golang-restful-api/repository"
	"rianmz-genz/golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
	
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr: "localhost:7611",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}