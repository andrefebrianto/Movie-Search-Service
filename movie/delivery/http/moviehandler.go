package http

import (
	"net/http"
	"strconv"

	"github.com/andrefebrianto/Search-Movie-Service/movie"
	"github.com/labstack/echo"
)

type HttpResponseObject struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type MovieHttpHandler struct {
	MovieUseCase movie.MovieUseCase
}

func HandleHttpRequest(ech *echo.Echo, movieUseCase movie.MovieUseCase) {
	handler := MovieHttpHandler{
		MovieUseCase: movieUseCase,
	}

	ech.GET("/api/v1/movies", handler.SearchMovies)
	ech.GET("/api/v1/movies/:id", handler.GetMovieDetail)
}

func (handler MovieHttpHandler) SearchMovies(context echo.Context) error {
	searchKeyword := context.QueryParam("searchword")
	page, _ := strconv.Atoi(context.QueryParam("pagination"))
	if searchKeyword == "" || page == 0 {
		return context.JSON(http.StatusBadRequest, HttpResponseObject{Message: "Invalid input parameter"})
	}

	ctx := context.Request().Context()

	movieMetaDatas, err := handler.MovieUseCase.SearchMovies(ctx, searchKeyword, page)
	if err != nil {
		if err.Error() == "not found" {
			return context.JSON(http.StatusNotFound, HttpResponseObject{Message: "Movie not found"})
		}
		return context.JSON(http.StatusInternalServerError, HttpResponseObject{Message: "internal server error"})
	}

	return context.JSON(http.StatusOK, HttpResponseObject{Message: "Movie search result retrieved", Data: movieMetaDatas})
}

func (handler MovieHttpHandler) GetMovieDetail(context echo.Context) error {
	movieId := context.Param("id")
	ctx := context.Request().Context()

	movie, err := handler.MovieUseCase.GetMovieDetail(ctx, movieId)
	if err != nil {
		if err.Error() == "not found" {
			return context.JSON(http.StatusNotFound, HttpResponseObject{Message: "Movie not found"})
		}
		return context.JSON(http.StatusInternalServerError, HttpResponseObject{Message: err.Error()})
	}

	return context.JSON(http.StatusOK, HttpResponseObject{Message: "Movie detail retrieved", Data: movie})
}
