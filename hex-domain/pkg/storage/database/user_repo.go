package database

import "github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"

func (s *Storage) GetUsersRepo() ([]models.User, error) {
	records := []models.User{}
	err := s.DB.Find(&records).Error
	if err != nil {
		return []models.User{}, err
	}
	return records, nil
}
