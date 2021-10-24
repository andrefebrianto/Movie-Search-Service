package usecase

import (
	"context"
	"time"

	"github.com/andrefebrianto/Search-Movie-Service/searchlog"
)

type SearchLogUseCase struct {
	mysqlCommandRepository searchlog.SearchLogCommandRepository
	contextTimeout         time.Duration
}

func CreateSearchLogUseCase(command searchlog.SearchLogCommandRepository, timeout time.Duration) SearchLogUseCase {
	return SearchLogUseCase{
		mysqlCommandRepository: command,
		contextTimeout:         timeout,
	}
}

func (useCase SearchLogUseCase) Create(ctx context.Context, searchLog *searchlog.SearchLog) error {
	contextWithTimeOut, cancel := context.WithTimeout(ctx, useCase.contextTimeout)
	defer cancel()

	err := useCase.mysqlCommandRepository.Create(contextWithTimeOut, searchLog)
	if err != nil {
		return err
	}

	return nil
}
