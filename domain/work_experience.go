package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	WorkExperienceTable = "work_experiences"
)

type WorkExperience struct {
	ID             uuid.UUID  `gorm:"primaryKey;type:char(36)" json:"id"`
	CompanyName    string     `gorm:"size:255;index" json:"company_name"`
	Role           string     `gorm:"size:255;index" json:"role"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
	Location       string     `json:"location"`
	Description    string     `json:"description"`
	CompanyLogoURL string     `json:"company_logo_url"`
	IsActive       bool       `gorm:"index" json:"is_active"`
}

type WorkExperienceRepository interface {
	Create(c context.Context, workExperience WorkExperience) error
	Retrieve(c context.Context, filter Filter) (workExperiences []WorkExperience, meta MetaResponse, err error)
	GetById(c context.Context, id uuid.UUID) (workExperience WorkExperience, err error)
	Update(c context.Context, id uuid.UUID, data WorkExperience) (workExperience WorkExperience, err error)
	Delete(c context.Context, id uuid.UUID) error
}

type WorkExperienceUsecase interface {
	Create(c context.Context, workExperience WorkExperience) error
	Retrieve(c context.Context, filter Filter) (workExperiences []WorkExperience, meta MetaResponse, err error)
	GetById(c context.Context, id uuid.UUID) (workExperience WorkExperience, err error)
	Update(c context.Context, id uuid.UUID, data WorkExperience) (workExperience WorkExperience, err error)
	Delete(c context.Context, id uuid.UUID) error
}
