package daos

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/config"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserDAO_Get(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewUserDAO()

	user, err := dao.Get(1)

	expected := map[string]string{"First Name": "John", "Last Name": "Doe", "Email": "john.doe@gmail.com"}

	assert.Nil(t, err)
	assert.Equal(t, expected["First Name"], user.FirstName)
	assert.Equal(t, expected["Last Name"], user.LastName)
	assert.Equal(t, expected["Email"], user.Email)
}

func TestUserDAO_GetNotPresent(t *testing.T) {
	config.Config.DB = test_data.ResetDB()
	dao := NewUserDAO()

	user, err := dao.Get(9999)

	assert.NotNil(t, err)
	assert.Equal(t, "", user.FirstName)
	assert.Equal(t, "", user.LastName)
	assert.Equal(t, "", user.Email)
}
