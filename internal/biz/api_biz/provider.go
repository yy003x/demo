package api_biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	NewPingLogic,
	NewActivityLogic,
)
