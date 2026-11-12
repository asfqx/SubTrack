package subscription

import (
	"context"
	"fmt"

	"github.com/asfqx/SubTrack/internal/domain"
	errs "github.com/asfqx/SubTrack/internal/errors"
	"github.com/google/uuid"
)

type Repository interface {
	GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Subscription, error)

	GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.Subscription, error)

	Create(ctx context.Context, sub *domain.Subscription) error

	Update(ctx context.Context, sub domain.UpdateSubscriptionParams, subUUID uuid.UUID) error

	Delete(ctx context.Context, subUUID uuid.UUID) error

	UpdateStatus(ctx context.Context, status domain.SubscriptionStatus, subUUID uuid.UUID) error

	CountByUserUUID(ctx context.Context, userUUID uuid.UUID) (int64, error)
}

type Service struct {
	repository Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repository: repo,
	}
}

func (s *Service) Create(ctx context.Context, params domain.CreateSubscriptionParams) (domain.Subscription, error) {
	count, err := s.repository.CountByUserUUID(ctx, params.UserUUID)
	if err != nil {
		return domain.Subscription{}, fmt.Errorf("create subscription: %w", err)
	}

	if count >= maxSubsPerUser {
		return domain.Subscription{}, errs.ErrLimitSubscriptionPerUser
	}

	sub, err := domain.NewSubscription(params)
	if err != nil {
		return sub, fmt.Errorf("create subscription: %w", err)
	}

	err = s.repository.Create(ctx, &sub)
	if err != nil {
		return sub, fmt.Errorf("create subscription: %w", err)
	}

	return sub, nil
}

func (s *Service) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Subscription, error) {
	sub, err := s.repository.GetByUUID(ctx, uuid)
	if err != nil {
		return sub, fmt.Errorf("get subscription by uuid: %w", err)
	}

	return sub, nil
}

func (s *Service) GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.Subscription, error) {
	subs, err := s.repository.GetByUserUUID(ctx, userUUID)
	if err != nil {
		return subs, fmt.Errorf("get subscription by user uuid: %w", err)
	}

	return subs, nil
}

func (s *Service) Update(ctx context.Context, params domain.UpdateSubscriptionParams, subUUID uuid.UUID) error {
	err := s.repository.Update(ctx, params, subUUID)
	if err != nil {
		return fmt.Errorf("update subscription: %w", err)
	}

	return nil
}

func (s *Service) Delete(ctx context.Context, subUUID uuid.UUID) error {
	err := s.repository.Delete(ctx, subUUID)
	if err != nil {
		return fmt.Errorf("delete subscription: %w", err)
	}

	return nil
}

func (s *Service) UpdateStatus(ctx context.Context, status domain.SubscriptionStatus, subUUID uuid.UUID) error {
	err := s.repository.UpdateStatus(ctx, status, subUUID)
	if err != nil {
		return fmt.Errorf("update subscription status: %w", err)
	}

	return nil
}
