package test

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"rianmz-genz/golang-restful-api/app"
	"rianmz-genz/golang-restful-api/controller"
	"rianmz-genz/golang-restful-api/helper"
	"rianmz-genz/golang-restful-api/middleware"
	"rianmz-genz/golang-restful-api/repository"
	"rianmz-genz/golang-restful-api/service"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)


func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(60 * time.Minute)
	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	requestBody := strings.NewReader(`{"name": "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:7611/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	truncateCategory(db)
	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:7611/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	
	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {

}
func TestUpdateCategoryFailed(t *testing.T) {
	
}

func TestGetCategorySuccess(t *testing.T) {

}
func TestGetCategoryFailed(t *testing.T) {
	
}

func TestDeleteCategorySuccess(t *testing.T) {

}
func TestDeleteCategoryFailed(t *testing.T) {
	
}

func TestGetListCategorySuccess(t *testing.T) {

}
func TestGetListCategoryFailed(t *testing.T) {
	
}

func TestUnauthorized(t *testing.T) {
	
}