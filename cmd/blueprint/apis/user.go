package apis

import (
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/daos"
	"github.com/MartinHeinz/go-project-blueprint/cmd/blueprint/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetUser is function for endpoint /api/v1/users to get User by ID
func GetUser(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
