// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import movie "github.com/andrefebrianto/Search-Movie-Service/movie"

// MovieDataProvider is an autogenerated mock type for the MovieDataProvider type
type MovieDataProvider struct {
	mock.Mock
}

// GetMovieDetail provides a mock function with given fields: ctx, imdbId
func (_m *MovieDataProvider) GetMovieDetail(ctx context.Context, imdbId string) (*movie.Movie, error) {
	ret := _m.Called(ctx, imdbId)

	var r0 *movie.Movie
	if rf, ok := ret.Get(0).(func(context.Context, string) *movie.Movie); ok {
		r0 = rf(ctx, imdbId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*movie.Movie)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, imdbId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchMovies provides a mock function with given fields: ctx, keyword, page
func (_m *MovieDataProvider) SearchMovies(ctx context.Context, keyword string, page int) (*movie.MovieSearch, error) {
	ret := _m.Called(ctx, keyword, page)

	var r0 *movie.MovieSearch
	if rf, ok := ret.Get(0).(func(context.Context, string, int) *movie.MovieSearch); ok {
		r0 = rf(ctx, keyword, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*movie.MovieSearch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, keyword, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
