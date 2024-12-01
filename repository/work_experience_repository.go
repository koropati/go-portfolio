package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/koropati/go-portfolio/domain"
	"gorm.io/gorm"
)

type workExperienceRepository struct {
	database  *gorm.DB
	table     string
	pageInit  int64
	limitInit int64
}

func NewWorkExperienceRepository(db *gorm.DB, table string, pageInit int64, limitInit int64) domain.WorkExperienceRepository {
	return &workExperienceRepository{
		database:  db,
		table:     table,
		pageInit:  pageInit,
		limitInit: limitInit,
	}
}

func (u *workExperienceRepository) Create(c context.Context, data domain.WorkExperience) error {
	result := u.database.WithContext(c).Table(u.table).Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *workExperienceRepository) Retrieve(c context.Context, filter domain.Filter) (workExperiences []domain.WorkExperience, meta domain.MetaResponse, err error) {
	query := u.database.WithContext(c).Table(u.table)

	if filter.Search != "" {
		query = query.Where("company_name LIKE ?", "%"+filter.Search+"%")
	}

	if filter.Limit <= 0 {
		filter.Limit = 5
	}

	if filter.WithPagination {
		offset := (filter.Page - 1) * filter.Limit
		query = query.Offset(int(offset)).Limit(int(filter.Limit))
	}

	result := query.Find(&workExperiences)
	if result.Error != nil {
		return nil, domain.MetaResponse{}, result.Error
	}

	// Hitung total records
	var totalRecords int64
	u.database.Table(u.table).Count(&totalRecords)

	meta = domain.MetaResponse{
		TotalRecords:    totalRecords,
		FilteredRecords: int64(result.RowsAffected),
		Page:            filter.Page,
		PerPage:         filter.Limit,
		TotalPages:      (totalRecords + filter.Limit - 1) / filter.Limit,
	}

	return workExperiences, meta, nil
}

func (u *workExperienceRepository) Update(c context.Context, id uuid.UUID, data domain.WorkExperience) (user domain.WorkExperience, err error) {
	result := u.database.WithContext(c).Table(u.table).Where(queryFindByID, id).Updates(data)
	if result.Error != nil {
		return domain.WorkExperience{}, result.Error
	}
	if result.RowsAffected == 0 {
		return domain.WorkExperience{}, errors.New("no data was updated")
	}
	err = u.database.WithContext(c).Table(u.table).Where(queryFindByID, id).First(&user).Error
	if err != nil {
		return domain.WorkExperience{}, err
	}
	return user, nil
}

func (u *workExperienceRepository) Delete(c context.Context, id uuid.UUID) error {
	result := u.database.WithContext(c).Table(u.table).Where(queryFindByID, id).Delete(&domain.WorkExperience{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("no data was deleted")
	}
	return nil
}

func (u *workExperienceRepository) GetById(c context.Context, id uuid.UUID) (user domain.WorkExperience, err error) {
	result := u.database.WithContext(c).Table(u.table).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return domain.WorkExperience{}, result.Error
	}
	return user, nil
}
