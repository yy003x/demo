package cli

import (
	"be_demo/internal/infrastructure/command"
	"be_demo/internal/service/cli_serv"

	"github.com/go-kratos/kratos/v2/middleware/metadata"
)

func NewCommandServer(
	demoCmdService *cli_serv.DemoCmdService,
) *command.Server {
	var opts = []command.ServerOption{
		command.WithMiddleware(
			metadata.Server(metadata.WithPropagatedPrefix("x-")),
		),
	}
	srv := command.NewServer(opts...)
	srv.RegisterCmdService(demoCmdService)
	return srv
}
