//go:build wireinject
// +build wireinject

package todos

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/google/wire"
	"github.com/shripadmhetre/golang-wire-demo/src/config"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/apm"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/todo"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/todo/todoservice"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/user"
	"github.com/shripadmhetre/golang-wire-demo/src/domain/user/userservice"
	"github.com/shripadmhetre/golang-wire-demo/src/infra/server"
	"github.com/shripadmhetre/golang-wire-demo/src/infra/somerealapm"
	"github.com/shripadmhetre/golang-wire-demo/src/infra/sqldb"
	"github.com/shripadmhetre/golang-wire-demo/src/infra/todorepo"
	"github.com/shripadmhetre/golang-wire-demo/src/infra/userrepo"
	"github.com/shripadmhetre/golang-wire-demo/src/todos/handler/api"
)

func Wire(enver config.Enver, logger *log.Logger) (*App, error) {
	wire.Build(

		// sqldb to DB.
		sqldb.NewConfig,
		sqldb.New,
		wire.Bind(new(DB), new(*sql.DB)),

		// Our imposter APM.
		somerealapm.NewConfig,
		somerealapm.New,
		wire.Bind(new(apm.APM), new(*somerealapm.APM)),
		wire.Bind(new(APM), new(*somerealapm.APM)),

		// User Repo.
		userrepo.Wired,
		wire.Bind(new(user.Repo), new(*userrepo.Repo)),

		// User Service.
		userservice.Wired,
		wire.Bind(new(user.Service), new(*userservice.Service)),

		// Todo Repo.
		todorepo.Wired,
		wire.Bind(new(todo.Repo), new(*todorepo.Repo)),

		// Todo Service.
		todoservice.Wired,
		wire.Bind(new(todo.Service), new(*todoservice.Service)),

		// API as our http.Handler.
		api.Wired,
		wire.Bind(new(http.Handler), new(*api.API)),

		// Http server.
		server.Wired,

		// This package - the application.
		newApp,
	)

	// Requirement for compilation.
	return nil, nil
}
