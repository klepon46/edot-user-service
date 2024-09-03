package repository

// @Notice: Register your repositories here

type IRegistry interface {
	GetUserRepository() IUserRepository
}

type Registry struct {
	UserRepository IUserRepository
}

func NewRegistryRepository(
	userRepository IUserRepository,
) *Registry {
	return &Registry{
		UserRepository: userRepository,
	}
}

func (r *Registry) GetUserRepository() IUserRepository {
	return r.UserRepository
}
