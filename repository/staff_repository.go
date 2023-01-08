package repository

import (
	"api_temankopi/model"

	"gorm.io/gorm"
)

type StaffRepository interface {
	AddStaff(staff model.Staff) (model.Staff, error)
	EditStaff(staff model.Staff) (model.Staff, error)
	DeleteStaff(staff model.Staff) error
	ViewStaff() ([]model.Staff, error)
	ViewStaffByID(id int) (model.Staff, error)
}

type staffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *staffRepository {
	return &staffRepository{
		db: db,
	}
}

func (s *staffRepository) AddStaff(staff model.Staff) (model.Staff, error) {
	err := s.db.Create(&staff).Error
	return staff, err
}

func (s *staffRepository) EditStaff(staff model.Staff) (model.Staff, error) {
	err := s.db.Save(&staff).Error
	return staff, err
}

func (s *staffRepository) DeleteStaff(staff model.Staff) error {
	err := s.db.Debug().Delete(&staff).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *staffRepository) ViewStaff() ([]model.Staff, error) {
	var staffs []model.Staff
	err := s.db.Debug().Find(&staffs).Error
	if err != nil {
		return staffs, err
	}
	return staffs, nil
}

func (s *staffRepository) ViewStaffByID(id int) (model.Staff, error) {
	var staff model.Staff
	err := s.db.Where("id = ?", id).First(&staff).Error
	if err != nil {
		return staff, err
	}
	return staff, nil
}
