package mysql

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type GromEntity[T comparable] interface {
	TableName() string
	GetID() T
}

type GormRepository[E GromEntity[T], T comparable] struct {
	log *log.Helper
	db  *gorm.DB
}

func NewGormRepository[E GromEntity[T], T comparable](
	logger log.Logger,
	db *gorm.DB,
) *GormRepository[E, T] {
	return &GormRepository[E, T]{
		log: log.NewHelper(log.With(logger, "x_module", "api_service/PingService")),
		db:  db,
	}
}

func (gr *GormRepository[E, ID]) rwdb(ctx context.Context) *gorm.DB {
	return gr.db.WithContext(ctx)
}

func (gr *GormRepository[E, ID]) Create(ctx context.Context, data *E) (ID, error) {
	err := gr.rwdb(ctx).Create(data).Error
	if err != nil {
		gr.log.WithContext(ctx).Errorf("Create failed %+v - %+v", data, err)
	}
	return (*data).GetID(), err
}

func (gr *GormRepository[E, ID]) Update(ctx context.Context, data *E) (ID, error) {
	err := gr.rwdb(ctx).Create(data).Error
	if err != nil {
		gr.log.WithContext(ctx).Errorf("Create failed %+v - %+v", data, err)
	}
	return (*data).GetID(), err
}
