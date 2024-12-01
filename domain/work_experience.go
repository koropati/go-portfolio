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
	ID             uuid.UUID  `gorm:"primaryKey;type:char(36)" validate:"required" json:"id" form:"id"`
	CompanyName    string     `gorm:"size:255;index" validate:"required" json:"company_name" form:"company_name"`
	Role           string     `gorm:"size:255;index" validate:"required" json:"role" form:"role"`
	StartDate      time.Time  `gorm:"type:date" json:"start_date" validate:"required" form:"start_date"`
	EndDate        *time.Time `gorm:"type:date" json:"end_date" form:"end_date"`
	Location       string     `json:"location" form:"location"`
	Description    string     `json:"description" form:"description"`
	CompanyLogoURL string     `json:"company_logo_url" form:"company_logo_url"`
	IsActive       bool       `gorm:"index" json:"is_active" form:"is_active"`
}

func (data *WorkExperience) GenerateID() (err error) {
	data.ID, err = uuid.NewUUID()
	if err != nil {
		return err
	}
	return nil
}

func (data *WorkExperience) SetFileURL(fileURL string) {
	data.CompanyLogoURL = fileURL
}

func (data *WorkExperience) SetActive() {
	data.IsActive = true
}

func (data *WorkExperience) SetNonActive() {
	data.IsActive = false
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
