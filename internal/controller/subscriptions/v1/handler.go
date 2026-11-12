package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/asfqx/SubTrack/internal/domain"
	"github.com/google/uuid"
)

type Service interface {
	Create(ctx context.Context, params domain.CreateSubscriptionParams) (domain.Subscription, error)

	GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Subscription, error)

	GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.Subscription, error)

	Update(ctx context.Context, params domain.UpdateSubscriptionParams, subUUID uuid.UUID) error

	Delete(ctx context.Context, subUUID uuid.UUID) error

	UpdateStatus(ctx context.Context, status domain.SubscriptionStatus, subUUID uuid.UUID) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) GetByUUID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var input GetSubscriptionRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
	}

	subscription, err := h.service.GetByUUID(ctx, input.UUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
