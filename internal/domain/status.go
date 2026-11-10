package domain

import errs "github.com/asfqx/SubTrack/internal/errors"

type SubscriptionStatus string

const (
	Active   SubscriptionStatus = "ACTIVE"
	Canceled SubscriptionStatus = "CANCELED"
	Paused   SubscriptionStatus = "PAUSED"
)

func NewSubscriptionStatus(value string) (SubscriptionStatus, error) {
	status := SubscriptionStatus(value)
	switch status {
	case Active, Canceled, Paused:
		return status, nil
	default:
		return "", errs.ErrInvalidSubscriptionStatus
	}
}
