package movies

import (
	"github.com/fabres21s/adn-go-movies/datasources/mysql/movies_db"
	"github.com/fabres21s/adn-go-movies/utils/errors"
	"github.com/fabres21s/adn-go-movies/utils/mysql_utils"
)

const (
	queryInsertMovie = "INSERT INTO movie (code, title) VALUES (?,?)"
	queryGetMovie    = "SELECT id, code, title FROM movie WHERE id = ?"
	queryUpdateMovie = "UPDATE movie SET code = ?, title = ? WHERE id = ?"
	queryDeleteMovie = "DELETE FROM movie WHERE id = ?"
	queryFindMovies  = "SELECT id, code, title FROM movie ORDER BY code "
)

func (movie *Movie) Get() *errors.RestErr {

	stmt, err := movies_db.Client.Prepare(queryGetMovie)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	result := stmt.QueryRow(movie.Id)

	if getErr := result.Scan(&movie.Id, &movie.Code, &movie.Title); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}

	return nil
}

func (movie *Movie) Save() *errors.RestErr {
	stmt, err := movies_db.Client.Prepare(queryInsertMovie)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(movie.Code, movie.Title)

	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	movieId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	movie.Id = movieId
	return nil
}

func (movie *Movie) Update() *errors.RestErr {
	stmt, err := movies_db.Client.Prepare(queryUpdateMovie)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(movie.Code, movie.Title, movie.Id)

	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (movie *Movie) Delete() *errors.RestErr {
	stmt, err := movies_db.Client.Prepare(queryDeleteMovie)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(movie.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (movie *Movie) Find() ([]Movie, *errors.RestErr) {
	stmt, err := movies_db.Client.Prepare(queryFindMovies)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]Movie, 0)

	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Id, &movie.Code, &movie.Title); err != nil {
			return nil, mysql_utils.ParseError(err)
		}

		results = append(results, movie)
	}

	return results, nil
}
