package movie

import "context"

type MovieMetaData struct {
	Title, Year, ImdbID, Type, Poster string
}

type MovieSearch struct {
	SearchResult []MovieMetaData
	TotalResults int
}

type Movie struct {
	Title, Year, Rated, Released, Runtime, Genre, Director, Writer, Actors, Plot, Language, Country, Awards, Poster,
	Metascore, ImdbRating, ImdbVotes, ImdbID, Type, DVD, BoxOffice, Production, Website string
	Ratings []struct {
		Source, Value string
	}
}

type MovieDataProvider interface {
	SearchMovies(ctx context.Context, keyword string, page int) (*MovieSearch, error)
	GetMovieDetail(ctx context.Context, imdbId string) (*Movie, error)
}

type MovieUseCase interface {
	SearchMovies(ctx context.Context, keyword string, page int) (*MovieSearch, error)
	GetMovieDetail(ctx context.Context, imdbId string) (*Movie, error)
}
