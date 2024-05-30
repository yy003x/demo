package app_repo

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type greeterRepo struct {
	log *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(
	logger log.Logger,
) GreeterRepo {
	return &greeterRepo{
		log: log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*Greeter, error) {
	return nil, nil
}
