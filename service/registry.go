package service

type IRegistry interface {
	GetUserService() IUserService
}

type Registry struct {
	userService IUserService
}

func NewRegistry(userService IUserService) *Registry {
	return &Registry{
		userService: userService,
	}
}

func (r *Registry) GetUserService() IUserService {
	return r.userService
}
