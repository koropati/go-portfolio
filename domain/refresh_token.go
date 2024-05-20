package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	RefreshTokenTable = "refresh_tokens"
)

type RefreshToken struct {
	Token     string `gorm:"primaryKey" json:"token"`
	PairToken string `json:"pair_token"`
	UserID    uuid.UUID
	Revoked   bool
	CreatedAt time.Time
	ExpiresAt int64
}

type RefreshTokenRepository interface {
	Create(c context.Context, refreshToken RefreshToken) error
	Revoke(c context.Context, token string) error
	RevokeByPairToken(c context.Context, pairToken string) error
	RevokeByUserID(c context.Context, userID uuid.UUID) error
	IsValid(c context.Context, token string) bool
	Delete(c context.Context, token string) error
}

type RefreshTokenUsecase interface {
	Create(c context.Context, refreshToken RefreshToken) error
	Revoke(c context.Context, token string) error
	RevokeByPairToken(c context.Context, pairToken string) error
	RevokeByUserID(c context.Context, userID uuid.UUID) error
	IsValid(c context.Context, token string) bool
	Delete(c context.Context, token string) error
}
