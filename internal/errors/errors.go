package errors

import "errors"

var (
	ErrInvalidCurrency           = errors.New("Currency is invalid")
	ErrInvalidSubscriptionStatus = errors.New("Status of subscription is invalid")
	ErrInvalidStartedAt          = errors.New("StartedAt of subscription is invalid")
	ErrNameTooLong               = errors.New("Name too long")
	ErrNameTooShort              = errors.New("Name too short")
	ErrDescriptionTooLong        = errors.New("Description too long")
	ErrNoRows                    = errors.New("No rows to return")
	ErrSubscriptionAlreadyExists = errors.New("Subscription already exists")
	ErrLimitSubscriptionPerUser  = errors.New("The limit for creating a subscription for the user has been exceeded")
)
