package searchlog

import (
	"context"
	"time"
)

type SearchLog struct {
	KeyWord   string
	Page      int
	Timestamp time.Time
}

type SearchLogUseCase interface {
	Create(ctx context.Context, searchLog *SearchLog) error
}

type SearchLogCommandRepository interface {
	Create(ctx context.Context, searchLog *SearchLog) error
}
