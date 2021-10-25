package grpc_test

import (
	"context"
	"errors"
	"testing"

	"github.com/andrefebrianto/Search-Movie-Service/mocks"
	"github.com/andrefebrianto/Search-Movie-Service/movie"
	grpcMovieHandler "github.com/andrefebrianto/Search-Movie-Service/movie/delivery/grpc"
	grpcModel "github.com/andrefebrianto/Search-Movie-Service/movie/delivery/grpc/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchMovies(t *testing.T) {
	mockMovieUseCase := new(mocks.MovieUseCase)

	t.Run("should return invalid input parameter error", func(t *testing.T) {
		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieKeywordAndPage{
			Keyword: "",
			Page:    0,
		}

		result, err := grpcHandler.SearchMovies(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("invalid input parameter")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return not found error", func(t *testing.T) {
		mockError := errors.New("not found")
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, mockError).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieKeywordAndPage{
			Keyword: "Thor",
			Page:    1,
		}

		result, err := grpcHandler.SearchMovies(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("movie not found")
			assert.Equal(t, mockError, err)
		}
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should return internal server error", func(t *testing.T) {
		mockError := errors.New("internal server error")
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, mockError).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieKeywordAndPage{
			Keyword: "Thor",
			Page:    1,
		}

		result, err := grpcHandler.SearchMovies(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("internal server error")
			assert.Equal(t, mockError, err)
		}
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should return movie search result", func(t *testing.T) {
		movieSearchResult := &movie.MovieSearch{
			SearchResult: []movie.MovieMetaData{{Title: "Incredibles 2", Year: "2018", ImdbID: "tt3606756", Type: "game", Poster: ""}},
			TotalResults: 1,
		}
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(movieSearchResult, nil).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieKeywordAndPage{
			Keyword: "Thor",
			Page:    1,
		}

		result, err := grpcHandler.SearchMovies(context.TODO(), params)
		assert.Nil(t, err)
		assert.Equal(t, int(result.TotalResults), 1)
		mockMovieUseCase.AssertExpectations(t)
	})
}

func TestGetMovieDetail(t *testing.T) {
	mockMovieUseCase := new(mocks.MovieUseCase)

	t.Run("should return invalid input parameter error", func(t *testing.T) {
		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieId{
			Id: "",
		}

		result, err := grpcHandler.GetMovieDetail(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("invalid input parameter")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return not found error", func(t *testing.T) {
		mockError := errors.New("not found")
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(nil, mockError).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieId{
			Id: "tt3606756",
		}

		result, err := grpcHandler.GetMovieDetail(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("movie not found")
			assert.Equal(t, mockError, err)
		}
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should return internal server error", func(t *testing.T) {
		mockError := errors.New("internal server error")
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(nil, mockError).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieId{
			Id: "tt3606756",
		}

		result, err := grpcHandler.GetMovieDetail(context.TODO(), params)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("internal server error")
			assert.Equal(t, mockError, err)
		}
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should return movie search result", func(t *testing.T) {
		movieDetail := &movie.Movie{
			Title:    "Incredibles 2",
			Year:     "2018",
			Rated:    "PG",
			Genre:    "Animation, Action, Adventure",
			ImdbID:   "tt3606756",
			Writer:   "Brad Bird",
			Director: "Brad Bird",
		}
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(movieDetail, nil).Once()

		grpcHandler := grpcMovieHandler.MovieGrpcServer{
			MovieUseCase: mockMovieUseCase,
		}

		params := &grpcModel.MovieId{
			Id: "tt3606756",
		}

		result, err := grpcHandler.GetMovieDetail(context.TODO(), params)
		assert.Nil(t, err)
		assert.Equal(t, result.Title, movieDetail.Title)
		assert.Equal(t, result.Year, movieDetail.Year)
		mockMovieUseCase.AssertExpectations(t)
	})
}
