package service

import (
	"api_temankopi/dto"
	"api_temankopi/model"
	"api_temankopi/repository"
	"errors"
)

type StaffService interface {
	AddStaff(product dto.StaffCreate) (model.Staff, error)
	EditStaff(id int, product dto.StaffUpdate) (model.Staff, error)
	DeleteStaff(id int) error
	ViewStaff() ([]model.Staff, error)
}

type staffService struct {
	staffRepository repository.StaffRepository
}

func NewStaffService(staffRepository repository.StaffRepository) *staffService {
	return &staffService{staffRepository}
}

func (s *staffService) AddStaff(staff dto.StaffCreate) (model.Staff, error) {
	var newStaff = model.Staff{
		Nama:   staff.Nama,
		Posisi: staff.Posisi,
		Shift:  staff.Shift,
	}

	response, err := s.staffRepository.AddStaff(newStaff)
	if err != nil {
		return model.Staff{}, err
	}

	return response, nil
}

func (s *staffService) EditStaff(id int, staff dto.StaffUpdate) (model.Staff, error) {
	editedStaff, err := s.staffRepository.ViewStaffByID(id)
	if err != nil {
		return model.Staff{}, err
	}

	editedStaff.Nama = staff.Nama
	editedStaff.Posisi = staff.Posisi
	editedStaff.Shift = staff.Shift

	result, err := s.staffRepository.EditStaff(editedStaff)
	return result, err
}

func (s *staffService) DeleteStaff(id int) error {
	deletedStaff, err := s.staffRepository.ViewStaffByID(id)
	if err != nil {
		return errors.New("There's nothing!")
	}

	err = s.staffRepository.DeleteStaff(deletedStaff)
	if err != nil {
		return err
	}
	return nil
}

func (s *staffService) ViewStaff() ([]model.Staff, error) {
	staff, err := s.staffRepository.ViewStaff()
	if err != nil {
		return staff, err
	}
	return staff, nil
}
