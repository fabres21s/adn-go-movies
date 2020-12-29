package services

import (
	"github.com/fabres21s/adn-go-movies/domain/movies"
	"github.com/fabres21s/adn-go-movies/utils/errors"
)

func GetMovie(movieId int64) (*movies.Movie, *errors.RestErr) {

	result := &movies.Movie{Id: movieId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateMovie(movie movies.Movie) (*movies.Movie, *errors.RestErr) {
	if err := movie.Validate(); err != nil {
		return nil, err
	}

	if err := movie.Save(); err != nil {
		return nil, err
	}
	return &movie, nil
}

func UpdateMovie(isPartial bool, movie movies.Movie) (*movies.Movie, *errors.RestErr) {
	current, err := GetMovie(movie.Id)

	if err != nil {
		return nil, err
	}

	if isPartial {
		if movie.Code != "" {
			current.Code = movie.Code
		} else if movie.Title != "" {
			current.Title = movie.Title
		}
	} else {
		current.Code = movie.Code
		current.Title = movie.Title
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteMovie(movieId int64) *errors.RestErr {
	movie := &movies.Movie{Id: movieId}
	return movie.Delete()
}

func Search() ([]movies.Movie, *errors.RestErr) {
	dao := &movies.Movie{}
	return dao.Find()
}
