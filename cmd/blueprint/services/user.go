package services

import "github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"

type userDAO interface {
	Get(id uint) (*models.User, error)
}

type UserService struct {
	dao userDAO
}

// NewUserService creates a new UserService with the given user DAO.
func NewUserService(dao userDAO) *UserService {
	return &UserService{dao}
}

// Get just retrieves user using User DAO, here can be additional logic for processing data retrieved by DAOs
func (s *UserService) Get(id uint) (*models.User, error) {
	return s.dao.Get(id)
}
