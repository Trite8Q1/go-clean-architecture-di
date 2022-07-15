package rest

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"

	"github.com/trite8q1/go-clean-architecture-di/internal/user/service"
	"github.com/trite8q1/go-clean-architecture-di/pkg/entity"
)

type RestUserController interface {
	GetUsers(c *gin.Context)
	AddUser(c *gin.Context)
}

type restUserController struct {
	userService service.UserService
}

//NewUserController: constructor, dependency injection from user service and firebase service
func NewUserController(s service.UserService) RestUserController {
	return &restUserController{
		userService: s,
	}
}
func (u *restUserController) GetUsers(c *gin.Context) {
	users, err := u.userService.FindAll()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error while getting users"})
		return
	}
	c.JSON(http.StatusOK, users)
}
func (u *restUserController) AddUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err1 := (u.userService.Validate(&user)); err1 != nil {
		sentry.CaptureException(err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	if ageValidation := (u.userService.ValidateAge(&user)); ageValidation != true {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB"})
		return
	}
	u.userService.Create(&user)
	c.JSON(http.StatusOK, user)
}
