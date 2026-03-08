package middleware

import "ecommerce/config"

type Middlewares struct {
	configuration *config.Config
}

func NewMiddlewares(configuration *config.Config) *Middlewares {
	return &Middlewares{
		configuration: configuration,
	}
}
