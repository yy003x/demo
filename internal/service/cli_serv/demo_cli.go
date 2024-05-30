package cli_serv

import (
	"be_demo/internal/biz/cli_biz"
	"be_demo/internal/infrastructure/command"
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cobra"
)

var _ command.ICliService = (*DemoCmdService)(nil)

//DemoCmdService  ./bin/command --conf=./configs/test --log=./logs test
type DemoCmdService struct {
	log          *log.Helper
	cmd          string
	desc         string
	demoCliLogic *cli_biz.DemoCliLogic
}

// NewGreeterService new a greeter service.
func NewDemoCmdService(
	logger log.Logger,
	demoCliLogic *cli_biz.DemoCliLogic,
) *DemoCmdService {
	return &DemoCmdService{
		log:          log.NewHelper(logger),
		cmd:          "test",
		desc:         "这是一个测试",
		demoCliLogic: demoCliLogic,
	}
}

func (c *DemoCmdService) GetCmd() string {
	return c.cmd
}

func (c *DemoCmdService) GetDesc() string {
	return c.desc
}

func (c *DemoCmdService) StringVar(cmdObj *cobra.Command) {
	// cmdObj.Flags().StringVarP(&repoURL, "repo-url", "r", "默认值", "描述")
}

func (c *DemoCmdService) CmdService(ctx context.Context, value []string) error {
	c.demoCliLogic.Test(ctx, "test")
	return nil
}
