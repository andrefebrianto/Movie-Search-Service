package omdbapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/movie"
	"github.com/andrefebrianto/Search-Movie-Service/searchlog"
)

type MovieDataProvider struct {
	BaseUrl   string `json:"baseUrl"`
	ApiKey    string `json:"apiKey"`
	SearchLog searchlog.SearchLogUseCase
}

type SearchResult struct {
	Search       []movie.MovieMetaData `json:"Search"`
	TotalResults string                `json:"totalResults"`
	Response     string                `json:"Response"`
	Error        string                `json:"Error"`
}

type MovieDetailResult struct {
	movie.Movie
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

func (dataProvider MovieDataProvider) SearchMovies(ctx context.Context, keyword string, page int) (*movie.MovieSearch, error) {
	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", dataProvider.BaseUrl, nil)
	if err != nil {
		return nil, err
	}
	query := request.URL.Query()
	query.Add("apikey", dataProvider.ApiKey)
	query.Add("s", keyword)
	query.Add("page", strconv.Itoa(page))
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("internal server error")
	}
	defer response.Body.Close()

	var responseData SearchResult
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	go func() {
		contex, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
		defer cancel()
		plainResponseData, _ := json.Marshal(responseData)
		searchLog := searchlog.SearchLog{
			Url:          request.RequestURI,
			Status:       response.StatusCode,
			ResponseData: string(plainResponseData),
			Timestamp:    time.Now().Local(),
		}
		err := dataProvider.SearchLog.Create(contex, &searchLog)
		if err != nil {
			fmt.Println("Search Log Database Error: ", err)
		}
	}()

	if responseData.Error == "Movie not found!" {
		return nil, errors.New("not found")
	} else if responseData.Error != "" {
		return nil, errors.New(responseData.Error)
	}

	totalResult, _ := strconv.Atoi(responseData.TotalResults)
	movieSearchResult := &movie.MovieSearch{
		SearchResult: responseData.Search,
		TotalResults: totalResult,
	}

	return movieSearchResult, nil
}

func (dataProvider MovieDataProvider) GetMovieDetail(ctx context.Context, imdbId string) (*movie.Movie, error) {
	client := http.Client{}
	request, err := http.NewRequestWithContext(ctx, "GET", dataProvider.BaseUrl, nil)
	if err != nil {
		return nil, err
	}
	query := request.URL.Query()
	query.Add("apikey", dataProvider.ApiKey)
	query.Add("i", imdbId)
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, errors.New("internal server error")
	}
	defer response.Body.Close()

	var responseData MovieDetailResult
	if err := json.NewDecoder(response.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	go func() {
		contex, cancel := context.WithTimeout(context.Background(), time.Duration(2*time.Second))
		defer cancel()
		plainResponseData, _ := json.Marshal(responseData)
		searchLog := searchlog.SearchLog{
			Url:          request.URL.String(),
			Status:       response.StatusCode,
			ResponseData: string(plainResponseData),
			Timestamp:    time.Now().Local(),
		}
		err := dataProvider.SearchLog.Create(contex, &searchLog)
		if err != nil {
			fmt.Println("Search Log Database Error: ", err)
		}
	}()

	if responseData.Error == "Incorrect IMDb ID." {
		return nil, errors.New("not found")
	} else if responseData.Error != "" {
		return nil, errors.New(responseData.Error)
	}

	movieDetail := responseData.Movie

	return &movieDetail, nil
}
