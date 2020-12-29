package app

import (
	"github.com/fabres21s/adn-go-movies/controllers/movies"
	"github.com/fabres21s/adn-go-movies/controllers/ping"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/movies/:id", movies.Get)
	router.GET("/movies", movies.Search)
	router.POST("/movies", movies.Create)
	router.PUT("/movies/:id", movies.Update)
	router.PATCH("/movies/:id", movies.Update)
	router.DELETE("/movies/:id", movies.Delete)

}
