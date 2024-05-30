package api_biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Test struct {
	Ping string
}

// GreeterRepo is a Greater repo.

// GreeterUsecase is a Greeter usecase.
type PingLogic struct {
	log *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewPingLogic(
	logger log.Logger,
) *PingLogic {
	return &PingLogic{
		log: log.NewHelper(logger),
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *PingLogic) Ping(ctx context.Context, g *Test) (*Test, error) {
	uc.log.WithContext(ctx).Infof("Ping: %v", g.Ping)
	return nil, nil
}
