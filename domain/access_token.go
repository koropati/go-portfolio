package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	AccessTokenTable = "access_tokens"
)

type AccessToken struct {
	Token     string `gorm:"primaryKey" json:"token"`
	UserID    uuid.UUID
	Revoked   bool
	CreatedAt time.Time
	ExpiresAt int64
}

type AccessTokenRepository interface {
	Create(c context.Context, accessToken AccessToken) error
	Revoke(c context.Context, token string) error
	RevokeByUserID(c context.Context, userID uuid.UUID) error
	IsValid(c context.Context, token string) bool
	Delete(c context.Context, token string) error
}

type AccessTokenUsecase interface {
	Create(c context.Context, accessToken AccessToken) error
	Revoke(c context.Context, token string) error
	RevokeByUserID(c context.Context, userID uuid.UUID) error
	IsValid(c context.Context, token string) bool
	Delete(c context.Context, token string) error
}
