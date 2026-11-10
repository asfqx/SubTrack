package subscriptions

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/asfqx/SubTrack/internal/adapter/postgres"
	"github.com/asfqx/SubTrack/internal/domain"
	errs "github.com/asfqx/SubTrack/internal/errors"
)

type repository struct {
	db *postgres.Postgres
}

func New(db *postgres.Postgres) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByUUID(ctx context.Context, uuid uuid.UUID) (domain.Subscription, error) {
	query := `
		SELECT 
			uuid, 
			name, 
			description, 
			amount,
			currency,
			next_billing_at,
			started_at,
			canceled_at,
			status,
			user_uuid,
			created_at,
			updated_at
		FROM subscriptions
		WHERE uuid = $1 
	`

	var sub domain.Subscription

	err := r.db.Pool().QueryRow(ctx, query, uuid).Scan(
		&sub.UUID,
		&sub.Name,
		&sub.Description,
		&sub.Amount,
		&sub.Currency,
		&sub.NextBillingAt,
		&sub.StartedAt,
		&sub.CanceledAt,
		&sub.Status,
		&sub.UserUUID,
		&sub.CreatedAt,
		&sub.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Subscription{}, errs.ErrNoRows
		}
		return domain.Subscription{}, fmt.Errorf("get subscription by uuid: %w", err)
	}

	return sub, nil
}

func (r *repository) GetByUserUUID(ctx context.Context, userUUID uuid.UUID) ([]domain.Subscription, error) {
	query := `
		SELECT 
			uuid,
			name,
			description,
			amount,
			currency
			next_billing_at,
			started_at,
			canceled_at,
			status,
			user_uuid,
			created_at,
			update_at
		FROM subscriptions
		WHERE user_uuid = $1
	`

	subscriptions := make([]domain.Subscription, 0)

	rows, err := r.db.Pool().Query(ctx, query, userUUID)
	if err != nil {
		return nil, fmt.Errorf("get subscriptions by user: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var sub domain.Subscription

		err := rows.Scan(
			&sub.UUID,
			&sub.Name,
			&sub.Description,
			&sub.Amount,
			&sub.Currency,
			&sub.NextBillingAt,
			&sub.StartedAt,
			&sub.CanceledAt,
			&sub.Status,
			&sub.UserUUID,
			&sub.CreatedAt,
			&sub.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan subscription: %w", err)
		}

		subscriptions = append(subscriptions, sub)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate subscriptions: %w", err)
	}

	return subscriptions, nil
}

func (r *repository) Create(ctx context.Context, sub *domain.Subscription) error {
	query := `
		INSERT INTO subscriptions (
			name,
			description,
			amount,
			currency
			next_billing_at,
			started_at,
			status,
			user_uuid
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`

	err := r.db.Pool().QueryRow(
		ctx,
		query,
		sub.Name,
		sub.Description,
		sub.Amount,
		sub.Currency,
		sub.NextBillingAt,
		sub.StartedAt,
		sub.Status,
		sub.UserUUID,
	).Scan(
		&sub.UUID,
		&sub.CreatedAt,
		&sub.UpdatedAt,
	)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return errs.ErrSubscriptionAlreadyExists
			}
		}

		return fmt.Errorf("create subscription: %w", err)
	}

	return nil
}

func (r *repository) Update(ctx context.Context, sub domain.Subscription) error {
	query := `
		UPDATE subscriptions
		SET 
			name = $1,
			description = $2 
			amount = $3,
			currency = $4,
			next_billing_at = $5,
			started_at = $6,
			updated_at = NOW()
		WHERE uuid = &7
	`

	tag, err := r.db.Pool().Exec(
		ctx,
		query,
		sub.Name,
		sub.Description,
		sub.Amount,
		sub.Currency,
		sub.NextBillingAt,
		sub.StartedAt,
		sub.UUID,
	)

	if tag.RowsAffected() == 0 {
		return errs.ErrNoRows
	}

	if err != nil {
		return fmt.Errorf("update subscription: %w", err)
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, subUUID uuid.UUID) error {
	query := `
		DELETE FROM subscriptions
		WHERE uuid = $1
	`

	tag, err := r.db.Pool().Exec(ctx, query, subUUID)

	if tag.RowsAffected() == 0 {
		return errs.ErrNoRows
	}

	if err != nil {
		return fmt.Errorf("delete subscription: %w", err)
	}

	return nil
}
