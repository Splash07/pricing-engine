package engine

import (
	"context"
	"fmt"
	"time"
)

type PriceChangeRequestStatus string

const (
	PriceChangeRequestStatusPendingApproval PriceChangeRequestStatus = "pending_approval"
	PriceChangeRequestStatusPendingInvalid  PriceChangeRequestStatus = "invalid"
	PriceChangeRequestStatusApproved        PriceChangeRequestStatus = "approved"
	PriceChangeRequestStatusRejected        PriceChangeRequestStatus = "rejected"
)

type PriceChangeRequest struct {
	ID                    uint
	SkuID                 string
	PriceType             PriceType
	Price                 Money
	Status                PriceChangeRequestStatus
	ValidationDescription string
	UpdatedAt             time.Time
	CreatedAt             time.Time
}

type PriceChangeRequestRepository interface {
	GetByID(context.Context, uint) (PriceChangeRequest, error)
	Create(context.Context, PriceChangeRequest) error
	Update(context.Context, PriceChangeRequest) error
}

func (r *PriceChangeRequest) Approves() error {
	if r.Status != PriceChangeRequestStatusPendingApproval {
		return fmt.Errorf("cannot accept invalid price change request")
	}
	r.Status = PriceChangeRequestStatusApproved
	return nil
}

func (r *PriceChangeRequest) Rejects() error {
	if r.Status != PriceChangeRequestStatusPendingApproval {
		return fmt.Errorf("cannot reject invalid price change request")
	}
	r.Status = PriceChangeRequestStatusRejected
	return nil
}
