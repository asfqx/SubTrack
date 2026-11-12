package v1

import "github.com/google/uuid"

type GetSubscriptionResponse struct {
	UUID          uuid.UUID `json:"uuid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Amount        int64     `json:"amount"`
	Currency      string    `json:"currency"`
	NextBillingAt int64     `json:"next_billing_at"`
	StartedAt     int64     `json:"started_at"`
	CanceledAt    int64     `json:"canceled_at"`
	Status        string    `json:"status"`
	UserUUID      uuid.UUID `json:"user_uuid"`
	CreatedAt     int64     `json:"created_at"`
	UpdatedAt     int64     `json:"updated_at"`
}

type GetSubscriptionRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type DeleteSubscriptionRequest struct {
	UUID uuid.UUID `json:"uuid"`
}

type GetSubscriptionByUserRequest struct {
	UserUUID uuid.UUID `json:"user_uuid"`
}

type UpdateSubscriptionRequest struct {
	UUID          uuid.UUID `json:"uuid"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Amount        int64     `json:"amount"`
	Currency      string    `json:"currency"`
	NextBillingAt int64     `json:"next_billing_at"`
	StartedAt     int64     `json:"started_at"`
}
