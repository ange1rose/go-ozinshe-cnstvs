package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goozinshe/handlers"
)

func main() {
	r := gin.Default()
	corsConfig := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"*"},
		AllowMethods:    []string{"*"},
	}
	r.Use(cors.New(corsConfig))

	moviesHandler := handlers.NewMovieHandler()

	r.POST("/movies", moviesHandler.Create)
	r.PUT("/movies/:id", moviesHandler.Update)
	r.GET("/movies/", moviesHandler.HandlerFindAll)

	r.Run(":8081")

}
