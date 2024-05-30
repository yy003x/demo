package api_biz

import (
	"be_demo/internal/entity/api_vo"
	"be_demo/internal/entity/po"
	"be_demo/internal/infrastructure/utils"
	"be_demo/internal/repository/api_repo"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type ActivityLogic struct {
	log            *log.Helper
	dmActivityRepo *api_repo.DmActivityRepo
}

// NewGreeterUsecase new a Greeter usecase.
func NewActivityLogic(
	logger log.Logger,
	dmActivityRepo *api_repo.DmActivityRepo,
) *ActivityLogic {
	return &ActivityLogic{
		log:            log.NewHelper(logger),
		dmActivityRepo: dmActivityRepo,
	}
}

// AddActivity
func (bl *ActivityLogic) AddActivity(ctx context.Context, in *api_vo.AddActivityRequest) (*api_vo.AddActivityReply, error) {
	data := &po.DmActivity{}
	data.ActID = utils.GenId(1, 0)
	data.ActName = in.ActName
	data.ActType = in.ActType
	data.ActStatus = in.ActStatus
	data.StartTime = time.Unix(in.StartTime, 0)
	data.EndTime = time.Unix(in.EndTime, 0)
	data.Version = in.Version
	id, err := bl.dmActivityRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	out := &api_vo.AddActivityReply{}
	out.ID = id
	return out, nil
}

// EditActivity
func (bl *ActivityLogic) EditActivity(ctx context.Context, in *api_vo.EditActivityRequest) (*api_vo.EditActivityReply, error) {
	data := &po.DmActivity{}
	data.ActID = utils.GenId(1, 0)
	data.ActName = in.ActName
	data.ActType = in.ActType
	data.ActStatus = in.ActStatus
	data.StartTime = time.Unix(in.StartTime, 0)
	data.EndTime = time.Unix(in.EndTime, 0)
	data.Version = in.Version
	id, err := bl.dmActivityRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	out := &api_vo.EditActivityReply{}
	out.ID = id
	return out, nil
}

// AddActivity
func (bl *ActivityLogic) RemoveActivity(ctx context.Context, in *api_vo.AddActivityRequest) (*api_vo.AddActivityReply, error) {
	data := &po.DmActivity{}
	data.ActID = utils.GenId(1, 0)
	data.ActName = in.ActName
	data.ActType = in.ActType
	data.ActStatus = in.ActStatus
	data.StartTime = time.Unix(in.StartTime, 0)
	data.EndTime = time.Unix(in.EndTime, 0)
	data.Version = in.Version
	id, err := bl.dmActivityRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	out := &api_vo.AddActivityReply{}
	out.ID = id
	return out, nil
}

// AddActivity
func (bl *ActivityLogic) DetailActivity(ctx context.Context, in *api_vo.AddActivityRequest) (*api_vo.AddActivityReply, error) {
	data := &po.DmActivity{}
	data.ActID = utils.GenId(1, 0)
	data.ActName = in.ActName
	data.ActType = in.ActType
	data.ActStatus = in.ActStatus
	data.StartTime = time.Unix(in.StartTime, 0)
	data.EndTime = time.Unix(in.EndTime, 0)
	data.Version = in.Version
	id, err := bl.dmActivityRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	out := &api_vo.AddActivityReply{}
	out.ID = id
	return out, nil
}

// AddActivity
func (bl *ActivityLogic) ListActivity(ctx context.Context, in *api_vo.AddActivityRequest) (*api_vo.AddActivityReply, error) {
	data := &po.DmActivity{}
	data.ActID = utils.GenId(1, 0)
	data.ActName = in.ActName
	data.ActType = in.ActType
	data.ActStatus = in.ActStatus
	data.StartTime = time.Unix(in.StartTime, 0)
	data.EndTime = time.Unix(in.EndTime, 0)
	data.Version = in.Version
	id, err := bl.dmActivityRepo.Create(ctx, data)
	if err != nil {
		return nil, err
	}
	out := &api_vo.AddActivityReply{}
	out.ID = id
	return out, nil
}
