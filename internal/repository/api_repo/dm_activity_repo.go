package api_repo

import (
	"be_demo/internal/data"
	"be_demo/internal/entity/po"
	"be_demo/internal/infrastructure/mysql"

	"github.com/go-kratos/kratos/v2/log"
)

type DmActivityRepo struct {
	*mysql.GormRepository[po.DmActivity, int64]
	log *log.Helper
}

func NewDmActivityRepo(
	logger log.Logger,
	dbs *data.DBS,
) *DmActivityRepo {
	return &DmActivityRepo{
		mysql.NewGormRepository[po.DmActivity, int64](logger, dbs.RwDb),
		log.NewHelper(log.With(logger, "x_module", "dao/DmSiyouActivityRepo")),
	}
}

func (r *DmActivityRepo) SelectActivity() []po.DmActivity {
	list := []po.DmActivity{}
	return list
}
