package http

import (
	"github.com/gin-gonic/gin"
	"github.com/klepon46/edot-user-service/common/response"
	"github.com/klepon46/edot-user-service/model"
	"github.com/klepon46/edot-user-service/service"
)

type IUserDelivery interface {
	Login(c *gin.Context)
	Registry(c *gin.Context)
	Validate(c *gin.Context)
}

type user struct {
	serviceRegistry *service.Registry
}

func NewUserDelivery(serviceRegistry *service.Registry) *user {
	return &user{
		serviceRegistry: serviceRegistry,
	}
}

func (u user) Login(c *gin.Context) {
	ctx := c.Request.Context()
	var payload model.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		return
	}

	if payload.Email == "" || payload.Phone == "" || payload.Password == "" {
		c.JSON(response.BadRequest(ctx).ToHTTPCodeAndMap())
		return
	}

	login, err := u.serviceRegistry.GetUserService().Login(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(response.ParseErrorToHTTPCode(ctx, err))
		return
	}

	c.JSON(response.OK(ctx, login).ToHTTPCodeAndMap())
}

func (u user) Registry(c *gin.Context) {
	ctx := c.Request.Context()
	var payload model.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		return
	}

	if payload.Email == "" || payload.Phone == "" || payload.Password == "" {
		c.JSON(response.BadRequest(ctx).ToHTTPCodeAndMap())
		return
	}

	login, err := u.serviceRegistry.GetUserService().Register(ctx, payload)
	if err != nil {
		c.AbortWithStatusJSON(response.ParseErrorToHTTPCode(ctx, err))
		return
	}

	c.JSON(response.OK(ctx, login).ToHTTPCodeAndMap())
}

func (u user) Validate(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
