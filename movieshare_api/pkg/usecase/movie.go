package usecase

import (
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/model"
)

func GetMovieAtRandom() (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.FindMovieAtRandom()

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}

func GetMovieList(title *string) ([]model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movieList, error := moviePersistence.GetMovieList(title)

	if error != nil {
		return []model.Movie{}, error
	}
	return movieList, nil
}

func GetMovieID(ID uint64) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movieID, error := moviePersistence.GetMovieID(ID)

	if error != nil {
		return model.Movie{}, error
	}
	return movieID, nil
}
