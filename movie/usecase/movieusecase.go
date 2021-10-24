package usecase

import (
	"context"
	"log"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/movie"
	"github.com/andrefebrianto/Search-Movie-Service/searchlog"
)

type MovieUseCase struct {
	searchLogUseCase  searchlog.SearchLogUseCase
	movieDataProvider movie.MovieDataProvider
	contextTimeout    time.Duration
}

func CreateMovieUseCase(searchLogUseCase searchlog.SearchLogUseCase, movieDataProvider movie.MovieDataProvider, timeout time.Duration) MovieUseCase {
	return MovieUseCase{
		searchLogUseCase:  searchLogUseCase,
		movieDataProvider: movieDataProvider,
		contextTimeout:    timeout,
	}
}

func (useCase MovieUseCase) SearchMovies(ctx context.Context, keyword string, page int) (*movie.MovieSearch, error) {
	contextWithTimeout, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	resultError := make(chan error)
	resultChannel := make(chan *movie.MovieSearch)

	go func() {
		searchLog := searchlog.SearchLog{
			KeyWord:   keyword,
			Page:      page,
			Timestamp: time.Now().Local(),
		}
		err := useCase.searchLogUseCase.Create(contextWithTimeout, &searchLog)
		log.Println(err)
	}()

	go func() {
		result, err := useCase.movieDataProvider.SearchMovies(contextWithTimeout, keyword, page)
		resultChannel <- result
		resultError <- err
	}()

	movieSearchResult := <-resultChannel
	err := <-resultError
	return movieSearchResult, err
}

func (useCase MovieUseCase) GetMovieDetail(ctx context.Context, imdbId string) (*movie.Movie, error) {
	contextWithTimeout, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	resultError := make(chan error)
	resultChannel := make(chan *movie.Movie)

	// go func() {
	// 	searchLog := searchlog.SearchLog{
	// 		FullUrl:   "",
	// 		Path:      "",
	// 		Method:    "",
	// 		UserAgent: "",
	// 		Status:    200,
	// 		Timestamp: time.Now().Local(),
	// 	}
	// 	err := useCase.searchLogUseCase.Create(contextWithTimeout, &searchLog)
	// 	log.Println(err)
	// }()

	go func() {
		result, err := useCase.movieDataProvider.GetMovieDetail(contextWithTimeout, imdbId)
		resultChannel <- result
		resultError <- err
	}()

	movie := <-resultChannel
	err := <-resultError
	return movie, err
}
