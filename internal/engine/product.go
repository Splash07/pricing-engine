package engine

import (
	"context"
	"time"
)

// Product - domain model
type Product struct {
	ID                     uint
	SkuID                  string
	BlackPrice             Money
	PermanentMarkdownPrice Money
	TemporaryMarkdownPrice Money
	StartDate              time.Time
	EndDate                time.Time
	UpdatedAt              time.Time
}

type ProductRepository interface {
	GetOneBySkuID(context.Context, string) (Product, error)
	GetOneBySkuIDs(context.Context, []string) ([]Product, error)
}
