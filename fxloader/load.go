package fxloader

import (
	"rice-wine-shop/api/controllers"
	"rice-wine-shop/api/routers"
	"rice-wine-shop/core/adapters"
	"rice-wine-shop/core/adapters/repository"
	"rice-wine-shop/core/services"

	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
)

func Load() []fx.Option {
	return []fx.Option{
		fx.Options(loadAdapter()...),
		fx.Options(loadUseCase()...),
		fx.Options(loadValidator()...),
		fx.Options(loadEngine()...),
	}
}
func loadUseCase() []fx.Option {
	return []fx.Option{
		fx.Provide(services.NewServiceUser),
	}
}

func loadValidator() []fx.Option {
	return []fx.Option{
		fx.Provide(validator.New),
	}
}
func loadEngine() []fx.Option {
	return []fx.Option{
		fx.Provide(routers.NewApiRouter),
		fx.Provide(controllers.NewControllerSaveFile),
		fx.Provide(controllers.NewControllerUser),
	}
}
func loadAdapter() []fx.Option {
	return []fx.Option{
		fx.Provide(repository.NewDBHelper),
		fx.Provide(repository.NewUserRepository),
		fx.Provide(repository.NewFileRepository),
		fx.Provide(adapters.ConnectPgsql),
	}
}
