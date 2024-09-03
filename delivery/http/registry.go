package http

// @Notice: Register your http deliveries here

type IRegistry interface {
}

type Registry struct {
}

func NewRegistry() *Registry {
	return &Registry{}
}
