package listing

import "github.com/iqbalmaulanaardi/hex-domain-example/hex-domain/pkg/models"

type Repository interface {
	GetUsersRepo() ([]models.User, error)
}
type Service interface {
	GetUsersService() ([]models.User, error)
}
type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}
