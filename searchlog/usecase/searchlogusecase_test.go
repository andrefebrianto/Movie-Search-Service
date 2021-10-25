package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/mocks"
	"github.com/andrefebrianto/Search-Movie-Service/searchlog"
	searchLogUseCase "github.com/andrefebrianto/Search-Movie-Service/searchlog/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreate(t *testing.T) {
	mockSearchLogCommandRepository := new(mocks.SearchLogCommandRepository)
	timeout := time.Duration(2 * time.Second)

	t.Run("should success to create search link", func(t *testing.T) {
		mockSearchLogCommandRepository.On("Create", mock.Anything, mock.AnythingOfType("*searchlog.SearchLog")).Return(nil).Once()

		searchLogUseCase := searchLogUseCase.CreateSearchLogUseCase(mockSearchLogCommandRepository, timeout)
		searchLog := &searchlog.SearchLog{
			Url:          "https://youtube.com",
			ResponseData: "Dummy Text",
			Status:       200,
			Timestamp:    time.Now().Local(),
		}
		err := searchLogUseCase.Create(context.TODO(), searchLog)
		assert.Nil(t, err)
		mockSearchLogCommandRepository.AssertExpectations(t)
	})

	t.Run("should fail to create search link (database error)", func(t *testing.T) {
		errorMock := errors.New("database error")
		mockSearchLogCommandRepository.On("Create", mock.Anything, mock.AnythingOfType("*searchlog.SearchLog")).Return(errorMock).Once()

		searchLogUseCase := searchLogUseCase.CreateSearchLogUseCase(mockSearchLogCommandRepository, timeout)
		searchLog := &searchlog.SearchLog{
			Url:          "https://youtube.com",
			ResponseData: "Dummy Text",
			Status:       200,
			Timestamp:    time.Now().Local(),
		}
		err := searchLogUseCase.Create(context.TODO(), searchLog)
		if assert.Error(t, err) {
			assert.Equal(t, errorMock, err)
		}
		mockSearchLogCommandRepository.AssertExpectations(t)
	})
}
