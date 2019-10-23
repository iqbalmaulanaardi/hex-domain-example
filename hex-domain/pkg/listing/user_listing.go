package listing

import "github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"

func (s *service) GetUsersService() ([]models.User, error) {
	result, err := s.r.GetUsersRepo()
	if err != nil {
		return []models.User{}, err
	}
	return result, nil
}
