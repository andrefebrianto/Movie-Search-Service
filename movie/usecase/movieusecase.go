package usecase

import (
	"context"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/movie"
)

type MovieUseCase struct {
	movieDataProvider movie.MovieDataProvider
	contextTimeout    time.Duration
}

func CreateMovieUseCase(movieDataProvider movie.MovieDataProvider, timeout time.Duration) MovieUseCase {
	return MovieUseCase{
		movieDataProvider: movieDataProvider,
		contextTimeout:    timeout,
	}
}

func (useCase MovieUseCase) SearchMovies(ctx context.Context, keyword string, page int) (*movie.MovieSearch, error) {
	contextWithTimeout, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	result, err := useCase.movieDataProvider.SearchMovies(contextWithTimeout, keyword, page)

	return result, err
}

func (useCase MovieUseCase) GetMovieDetail(ctx context.Context, imdbId string) (*movie.Movie, error) {
	contextWithTimeout, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	result, err := useCase.movieDataProvider.GetMovieDetail(contextWithTimeout, imdbId)

	return result, err
}
