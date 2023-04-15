package main

import (
	"backend/cmd/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
	router := gin.Default()

	router.Use(app.enableCORS())

	// router.GET("/", controller.Home)
	// router.GET("/movies", controller.AllMovies)
	router.GET("/users", controller.AllUsers)
	router.GET("/users/:id", controller.GetUser)
	router.PUT("/users/:id", controller.UpdateUser)

	return router
}
