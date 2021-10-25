package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/mocks"
	"github.com/andrefebrianto/Search-Movie-Service/movie"
	movieUseCase "github.com/andrefebrianto/Search-Movie-Service/movie/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchMovies(t *testing.T) {
	mockDataProvider := new(mocks.MovieDataProvider)
	timeout := time.Duration(2 * time.Second)

	t.Run("should return movies search result", func(t *testing.T) {
		movieSearchResult := &movie.MovieSearch{
			SearchResult: []movie.MovieMetaData{{Title: "ABC", Year: "2020", ImdbID: "hs2d332", Type: "game", Poster: ""}},
			TotalResults: 1,
		}
		mockDataProvider.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(movieSearchResult, nil).Once()

		movieUseCase := movieUseCase.CreateMovieUseCase(mockDataProvider, timeout)
		result, err := movieUseCase.SearchMovies(context.TODO(), "ABC", 1)
		assert.Equal(t, result.TotalResults, movieSearchResult.TotalResults)
		assert.Nil(t, err)
		mockDataProvider.AssertExpectations(t)
	})

	t.Run("should return error movie not found", func(t *testing.T) {
		mockError := errors.New("movie not found")
		mockDataProvider.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, mockError).Once()

		movieUseCase := movieUseCase.CreateMovieUseCase(mockDataProvider, timeout)
		result, err := movieUseCase.SearchMovies(context.TODO(), "ABC", 1)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
		mockDataProvider.AssertExpectations(t)
	})
}

func TestGetMovieDetail(t *testing.T) {
	mockDataProvider := new(mocks.MovieDataProvider)
	timeout := time.Duration(2 * time.Second)

	t.Run("should return movies search result", func(t *testing.T) {
		movieResult := &movie.Movie{
			Title:  "Incredibles 2",
			Year:   "2018",
			ImdbID: "tt3606756",
			Genre:  "Animation, Action, Adventure",
			Rated:  "PG",
		}
		mockDataProvider.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(movieResult, nil).Once()

		movieUseCase := movieUseCase.CreateMovieUseCase(mockDataProvider, timeout)
		result, err := movieUseCase.GetMovieDetail(context.TODO(), "tt3606756")
		assert.Equal(t, result.Title, movieResult.Title)
		assert.Equal(t, result.ImdbID, movieResult.ImdbID)
		assert.Nil(t, err)
		mockDataProvider.AssertExpectations(t)
	})

	t.Run("should return error movie not found", func(t *testing.T) {
		mockError := errors.New("movie not found")
		mockDataProvider.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(nil, mockError).Once()

		movieUseCase := movieUseCase.CreateMovieUseCase(mockDataProvider, timeout)
		result, err := movieUseCase.GetMovieDetail(context.TODO(), "tt3606756")
		assert.Nil(t, result)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
		mockDataProvider.AssertExpectations(t)
	})
}
