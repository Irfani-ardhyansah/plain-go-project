package main

import (
	"golang-api-udemy/app"
	"golang-api-udemy/controller"
	exception "golang-api-udemy/excception"
	"golang-api-udemy/helper"
	"golang-api-udemy/middleware"
	"golang-api-udemy/repository"
	"golang-api-udemy/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/category/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/category/:categoryId", categoryController.Update)
	router.DELETE("/api/category/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
