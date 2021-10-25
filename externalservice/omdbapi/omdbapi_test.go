package omdbapi_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/andrefebrianto/Search-Movie-Service/externalservice/omdbapi"
	"github.com/andrefebrianto/Search-Movie-Service/httpclient"
	"github.com/andrefebrianto/Search-Movie-Service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSearchMovies(t *testing.T) {
	mockSearchLogUseCase := new(mocks.SearchLogUseCase)
	keyword := "Batman"
	page := 1

	t.Run("should return fail create new request error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "123",
			ApiKey:  "123456",
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, &http.Client{})

		result, err := omdbApi.SearchMovies(context.TODO(), keyword, page)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			fmt.Println(err)
		}
	})

	t.Run("should return execute http request error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		mockError := errors.New("error http client")
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: nil,
			ErrorData:    mockError,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.SearchMovies(context.TODO(), keyword, page)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return internal server error after receiving status code not equal 200", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusBadRequest,
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.SearchMovies(context.TODO(), keyword, page)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("internal server error")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return parsing json error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Invalid JSON format")),
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.SearchMovies(context.TODO(), keyword, page)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("invalid json format")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return internal server error after receiving error response", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("{\"Error\":\"No API key provided.\",\"Response\":\"False\"}")),
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.SearchMovies(context.TODO(), keyword, page)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("No API key provided.")
			assert.Equal(t, mockError, err)
		}
	})
}

func TestGetMovieDetail(t *testing.T) {
	mockSearchLogUseCase := new(mocks.SearchLogUseCase)
	movieId := "tt3606756"

	t.Run("should return fail create new request error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "123",
			ApiKey:  "123456",
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, &http.Client{})

		result, err := omdbApi.GetMovieDetail(context.TODO(), movieId)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			fmt.Println(err)
		}
	})

	t.Run("should return execute http request error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		mockError := errors.New("error http client")
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: nil,
			ErrorData:    mockError,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.GetMovieDetail(context.TODO(), movieId)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return internal server error after receiving status code not equal 200", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusBadRequest,
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.GetMovieDetail(context.TODO(), movieId)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("internal server error")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return parsing json error", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("Invalid JSON format")),
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.GetMovieDetail(context.TODO(), movieId)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("invalid json format")
			assert.Equal(t, mockError, err)
		}
	})

	t.Run("should return internal server error after receiving error response", func(t *testing.T) {
		omdbApiConfig := omdbapi.OmdbConfig{
			BaseUrl: "http://www.omdbapi.com/",
			ApiKey:  "faf7e5bb",
		}
		responseData := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewBufferString("{\"Error\":\"No API key provided.\",\"Response\":\"False\"}")),
		}
		mockHttpClinet := httpclient.HttpClientMock{
			ResponseData: responseData,
			ErrorData:    nil,
		}
		omdbApi := omdbapi.CreateOmdbApiClient(omdbApiConfig, mockSearchLogUseCase, mockHttpClinet)

		result, err := omdbApi.GetMovieDetail(context.TODO(), movieId)
		assert.Nil(t, result)
		if assert.Error(t, err) {
			mockError := errors.New("No API key provided.")
			assert.Equal(t, mockError, err)
		}
	})
}
