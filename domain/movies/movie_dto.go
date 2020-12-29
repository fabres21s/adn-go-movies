package movies

import (
	"strings"

	"github.com/fabres21s/adn-go-movies/utils/errors"
)

const (
	StatusActive = "active"
)

type Movie struct {
	Id    int64  `json:"id"`
	Code  string `json:"code"`
	Title string `json:"title"`
}

func (movie *Movie) Validate() *errors.RestErr {
	movie.Code = strings.TrimSpace(strings.ToLower(movie.Code))

	if movie.Code == "" {
		return errors.NewBadRequestError("invalid code movie")
	}

	movie.Title = strings.TrimSpace(movie.Title)
	if movie.Title == "" {
		return errors.NewBadRequestError("invalid title movie")
	}

	return nil
}
