package movies

import (
	"net/http"
	"strconv"

	"github.com/fabres21s/adn-go-movies/domain/movies"
	"github.com/fabres21s/adn-go-movies/services"
	"github.com/fabres21s/adn-go-movies/utils/errors"
	"github.com/gin-gonic/gin"
)

func getMovieId(movieIdParam string) (int64, *errors.RestErr) {
	movieId, movieErr := strconv.ParseInt(movieIdParam, 10, 64)

	if movieErr != nil {
		return 0, errors.NewBadRequestError("movie id should be a number")
	}

	return movieId, nil
}

func Get(c *gin.Context) {
	movieId, idErr := getMovieId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}
	movie, getErr := services.GetMovie(movieId)

	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	c.JSON(http.StatusOK, movie)
}

func Create(c *gin.Context) {
	var movie movies.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateMovie(movie)

	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func Search(c *gin.Context) {

	movies, err := services.Search()

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, movies)
}

func Update(c *gin.Context) {

	movieId, idErr := getMovieId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var movie movies.Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	movie.Id = movieId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UpdateMovie(isPartial, movie)

	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	movieId, idErr := getMovieId(c.Param("id"))

	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteMovie(movieId); err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}
