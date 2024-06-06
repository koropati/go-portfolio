package domain

import (
	"context"

	"github.com/google/uuid"
)

const (
	AccessTokenTable = "access_tokens"
)

type AccessToken struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"id"`
	Token     string    `gorm:"type:longtext" json:"token"`
	UserID    uuid.UUID `gorm:"type:char(36);not null;index;foreignKey:ID" json:"user_id"`
	Revoked   bool      `gorm:"default:false" json:"revoked"`
	CreatedAt int64     `gorm:"autoCreateTime" json:"created_at"`
	ExpiresAt int64     `gorm:"index" json:"expires_at"`
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
