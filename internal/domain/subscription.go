package domain

import (
	"time"

	errs "github.com/asfqx/SubTrack/internal/errors"
	"github.com/google/uuid"
)

type Subscription struct {
	UUID          uuid.UUID
	Name          string
	Description   string
	Amount        int64
	Currency      Currency
	NextBillingAt int64
	StartedAt     int64
	CanceledAt    int64
	Status        SubscriptionStatus
	UserUUID      uuid.UUID
	CreatedAt     int64
	UpdatedAt     int64
}

type CreateSubscriptionParams struct {
	Name          string
	Description   string
	Amount        int64
	Currency      Currency
	NextBillingAt int64
	StartedAt     int64
	Status        SubscriptionStatus
	UserUUID      uuid.UUID
}

func NewSubscription(params CreateSubscriptionParams) (Subscription, error) {
	now := time.Now()
	started := time.UnixMilli(params.StartedAt)

	if started.After(now) {
		return Subscription{}, errs.ErrInvalidStatedAt
	}

	if len(params.Name) > maxNameLen {
		return Subscription{}, errs.ErrNameTooLong
	}

	if len(params.Name) < minNameLen {
		return Subscription{}, errs.ErrNameTooShort
	}

	if len(params.Description) > maxDescriptionLen {
		return Subscription{}, errs.ErrDescriptionTooLong
	}

	return Subscription{
		Name:          params.Name,
		Description:   params.Description,
		Amount:        params.Amount,
		Currency:      params.Currency,
		NextBillingAt: params.NextBillingAt,
		StartedAt:     params.StartedAt,
		Status:        params.Status,
		UserUUID:      params.UserUUID,
	}, nil
}
