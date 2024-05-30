package repository

import (
	"be_demo/internal/repository/api_repo"
	"be_demo/internal/repository/app_repo"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	app_repo.NewGreeterRepo,
	api_repo.NewDmActivityRepo,
)
