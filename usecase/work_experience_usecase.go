package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/koropati/go-portfolio/domain"
)

type workExperienceUsecase struct {
	workExperienceRepository domain.WorkExperienceRepository
	contextTimeout           time.Duration
}

func NewWorkExperienceUsecase(workExperienceRepository domain.WorkExperienceRepository, timeout time.Duration) domain.WorkExperienceUsecase {
	return &workExperienceUsecase{
		workExperienceRepository: workExperienceRepository,
		contextTimeout:           timeout,
	}
}

func (u *workExperienceUsecase) Create(c context.Context, workExperience domain.WorkExperience) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.workExperienceRepository.Create(ctx, workExperience)
}

func (u *workExperienceUsecase) Retrieve(c context.Context, filter domain.Filter) (workExperiences []domain.WorkExperience, meta domain.MetaResponse, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.workExperienceRepository.Retrieve(ctx, filter)
}

func (u *workExperienceUsecase) Update(c context.Context, id uuid.UUID, workExperience domain.WorkExperience) (domain.WorkExperience, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.workExperienceRepository.Update(ctx, id, workExperience)
}

func (u *workExperienceUsecase) Delete(c context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.workExperienceRepository.Delete(ctx, id)
}

func (u *workExperienceUsecase) GetById(c context.Context, id uuid.UUID) (workExperience domain.WorkExperience, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()
	return u.workExperienceRepository.GetById(ctx, id)
}
