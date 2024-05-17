package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"andersoncoimbra/goapi/handlers"
	"andersoncoimbra/goapi/utils"
)	
func main()  {
	utils.ConnectDB()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	

	e.GET("/produtos", handlers.GetProdutos)
	e.GET("/produtos/:id", handlers.GetProduto)
	e.POST("/produtos", handlers.CreateProduto)
	e.PUT("/produtos/:id", handlers.UpdateProduto)
	e.DELETE("/produtos/:id", handlers.DeleteProduto)

	e.Logger.Fatal(e.Start(":8080"))
}