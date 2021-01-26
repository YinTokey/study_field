// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"unsplash_server/internal/dao"
	"unsplash_server/internal/service"
	"unsplash_server/internal/server/grpc"
	"unsplash_server/internal/server/http"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
