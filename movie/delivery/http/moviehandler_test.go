package http_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andrefebrianto/Search-Movie-Service/mocks"
	"github.com/andrefebrianto/Search-Movie-Service/movie"
	movieHttpHandler "github.com/andrefebrianto/Search-Movie-Service/movie/delivery/http"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSearchMovies(t *testing.T) {
	mockMovieUseCase := new(mocks.MovieUseCase)

	t.Run("should fail to get search movie result (invalid input parameter)", func(t *testing.T) {
		keyword := ""
		page := ""
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies", nil)
		assert.NoError(t, err)
		query := request.URL.Query()
		query.Add("searchword", keyword)
		query.Add("pagination", page)
		request.URL.RawQuery = query.Encode()

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies")

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.SearchMovies(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	})

	t.Run("should fail to get search movie result (movies not found)", func(t *testing.T) {
		keyword := "Batman"
		page := "1"
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies", nil)
		assert.NoError(t, err)
		query := request.URL.Query()
		query.Add("searchword", keyword)
		query.Add("pagination", page)
		request.URL.RawQuery = query.Encode()

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies")

		mockError := errors.New("not found")
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, mockError).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.SearchMovies(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should fail to get search movie result (internal server error)", func(t *testing.T) {
		keyword := "Batman"
		page := "1"
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies", nil)
		assert.NoError(t, err)
		query := request.URL.Query()
		query.Add("searchword", keyword)
		query.Add("pagination", page)
		request.URL.RawQuery = query.Encode()

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies")

		mockError := errors.New("internal server error")
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil, mockError).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.SearchMovies(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should success to get search movie result", func(t *testing.T) {
		keyword := "Batman"
		page := "1"
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies", nil)
		assert.NoError(t, err)
		query := request.URL.Query()
		query.Add("searchword", keyword)
		query.Add("pagination", page)
		request.URL.RawQuery = query.Encode()

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies")

		movieSearchResult := &movie.MovieSearch{
			SearchResult: []movie.MovieMetaData{{Title: "Incredible 2", Year: "2018", ImdbID: "tt3606756", Type: "Animation"}},
			TotalResults: 1,
		}
		mockMovieUseCase.On("SearchMovies", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(movieSearchResult, nil).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.SearchMovies(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})
}

func TestGetMovieDetail(t *testing.T) {
	mockMovieUseCase := new(mocks.MovieUseCase)
	movieId := "tt3606756"

	t.Run("should fail to get movie detail (invalid input parameter)", func(t *testing.T) {
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies"+movieId, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies/:id")
		context.SetParamNames("id")
		context.SetParamValues(movieId)

		mockError := errors.New("not found")
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(nil, mockError).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.GetMovieDetail(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should fail to get movie detail (internal server error)", func(t *testing.T) {
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies"+movieId, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies/:id")
		context.SetParamNames("id")
		context.SetParamValues(movieId)

		mockError := errors.New("interna server error")
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(nil, mockError).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.GetMovieDetail(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})

	t.Run("should success to get movie detail", func(t *testing.T) {
		ech := echo.New()
		request, err := http.NewRequest(echo.POST, "/api/v1/movies"+movieId, nil)
		assert.NoError(t, err)

		responseRecorder := httptest.NewRecorder()
		context := ech.NewContext(request, responseRecorder)
		context.SetPath("/api/v1/movies/:id")
		context.SetParamNames("id")
		context.SetParamValues(movieId)

		movie := &movie.Movie{
			Title:    "Incredibles 2",
			Year:     "2018",
			Rated:    "PG",
			Released: "2018",
			Genre:    "Animation, Action, Adventure",
			Director: "Brad Bird",
			Writer:   "Brad Bird",
			ImdbID:   "tt3606756",
		}
		mockMovieUseCase.On("GetMovieDetail", mock.Anything, mock.AnythingOfType("string")).Return(movie, nil).Once()

		httpHandler := movieHttpHandler.MovieHttpHandler{
			MovieUseCase: mockMovieUseCase,
		}
		err = httpHandler.GetMovieDetail(context)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, responseRecorder.Code)
		mockMovieUseCase.AssertExpectations(t)
	})
}
