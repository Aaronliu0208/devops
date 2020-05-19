// +build wireinject
// The build tag makes sure the stub is not built in the final build.
// https://github.com/google/wire/issues/117 一定和package 声明之间空行

package injector

import (
	"casicloud.com/ylops/backend/app/router"
	"github.com/google/wire"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGinEngine,
		router.RouterSet,
		InjectorSet,
	)

	return new(Injector), nil, nil
}
