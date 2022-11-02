package engine

import (
	"context"
	"time"
)

type UploadJobStatus string

const (
	UploadJobStatusPending   UploadJobStatus = "pending"
	UploadJobStatusCompleted UploadJobStatus = "completed"
	UploadJobStatusFailed    UploadJobStatus = "failed"
)

type UploadJob struct {
	ID                          string
	Status                      UploadJobStatus
	Description                 string
	PriceChangeRequestTotal     uint64
	PriceChangeRequestProcessed uint64
	UpdatedAt                   time.Time
	CreatedAt                   time.Time
}

type UploadJobRepository interface {
	Create(context.Context, UploadJob) error
	GetByID(context.Context, string) (UploadJob, error)
}
