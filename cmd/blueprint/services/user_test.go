package services

import (
	"errors"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUserService(t *testing.T) {
	dao := newMockUserDAO()
	s := NewUserService(dao)
	assert.Equal(t, dao, s.dao)
}

func TestUserService_Get(t *testing.T) {
	s := NewUserService(newMockUserDAO())
	user, err := s.Get(2)
	if assert.Nil(t, err) && assert.NotNil(t, user) {
		assert.Equal(t, "Ben", user.FirstName)
		assert.Equal(t, "Doe", user.LastName)
	}

	user, err = s.Get(100)
	assert.NotNil(t, err)
}

func newMockUserDAO() userDAO {
	return &mockUserDAO{
		records: []models.User{
			{Model: models.Model{ID: 1}, FirstName: "John", LastName: "Smith", Email: "john.smith@gmail.com", Address: "Dummy Value"},
			{Model: models.Model{ID: 2}, FirstName: "Ben", LastName: "Doe", Email: "ben.doe@gmail.com", Address: "Dummy Value"},
		},
	}
}

// Mock Get function that replaces real User DAO
func (m *mockUserDAO) Get(id uint) (*models.User, error) {
	for _, record := range m.records {
		if record.ID == id {
			return &record, nil
		}
	}
	return nil, errors.New("not found")
}

type mockUserDAO struct {
	records []models.User
}
