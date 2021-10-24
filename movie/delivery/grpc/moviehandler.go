package grpc

import (
	"context"
	"errors"

	"github.com/andrefebrianto/Search-Movie-Service/movie"
	"github.com/andrefebrianto/Search-Movie-Service/movie/delivery/grpc/model"
	"google.golang.org/grpc"
)

type MovieGrpcServer struct {
	MovieUseCase movie.MovieUseCase
	model.UnimplementedMoviesServer
}

func (handler MovieGrpcServer) SearchMovies(contex context.Context, param *model.MovieKeywordAndPage) (*model.MovieSearch, error) {
	if param.Keyword == "" || param.Page == 0 {
		// return context.JSON(http.StatusBadRequest, HttpResponseObject{Message: "Invalid input parameter"})
		return nil, errors.New("")
	}

	movieMetaDatas, err := handler.MovieUseCase.SearchMovies(contex, param.Keyword, int(param.Page))
	if err != nil {
		if err.Error() == "not found" {
			return nil, errors.New("movie not found")
		}
		return nil, errors.New("internal server error")
	}

	var result = new(model.MovieSearch)

	result.TotalResults = int32(movieMetaDatas.TotalResults)
	result.SearchResult = make([]*model.MovieSearch_MovieMetaData, 0)

	for _, metaData := range movieMetaDatas.SearchResult {
		entry := &model.MovieSearch_MovieMetaData{
			Title:  metaData.Title,
			Year:   metaData.Year,
			ImdbID: metaData.ImdbID,
			Type:   metaData.Type,
			Poster: metaData.Poster,
		}
		result.SearchResult = append(result.SearchResult, entry)
	}

	return result, nil
}

func (handler MovieGrpcServer) GetMovieDetail(ctx context.Context, param *model.MovieId) (*model.Movie, error) {
	movie, err := handler.MovieUseCase.GetMovieDetail(ctx, param.Id)
	if err != nil {
		if err.Error() == "not found" {
			return nil, errors.New("movie not found")
		}
		return nil, errors.New("internal server error")
	}

	var result = new(model.Movie)

	result.Title = movie.Title
	result.Year = movie.Year
	result.Rated = movie.Rated
	result.Released = movie.Released
	result.Runtime = movie.Runtime
	result.Genre = movie.Genre
	result.Director = movie.Director
	result.Writer = movie.Writer
	result.Actors = movie.Actors
	result.Plot = movie.Plot
	result.Language = movie.Language
	result.Country = movie.Country
	result.Awards = movie.Awards
	result.Poster = movie.Poster
	result.Metascore = movie.Metascore
	result.ImdbRating = movie.ImdbRating
	result.ImdbVotes = movie.ImdbVotes
	result.ImdbID = movie.ImdbID
	result.Type = movie.Type
	result.DVD = movie.DVD
	result.BoxOffice = movie.BoxOffice
	result.Production = movie.Production
	result.Website = movie.Website

	return result, nil
}

func HandleGrpcRequest(server *grpc.Server, movieUseCase movie.MovieUseCase) {
	handler := MovieGrpcServer{
		MovieUseCase: movieUseCase,
	}

	model.RegisterMoviesServer(server, handler)
}
