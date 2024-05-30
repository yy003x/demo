package nacosx

import (
	"be_demo/internal/conf"

	knacos "github.com/go-kratos/kratos/contrib/config/nacos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func GetNacosClient(bc *conf.Bootstrap) []config.Source {
	nconf := bc.GetPublic().GetNacos()
	// Nacos 服务器配置
	sc := []constant.ServerConfig{}
	for _, v := range nconf.GetServerConfigs() {
		sc = append(sc,
			constant.ServerConfig{
				IpAddr: v.GetIp(),
				Port:   v.GetPort(),
			},
		)
	}
	cc := constant.ClientConfig{
		NamespaceId:         nconf.GetClientConfig().GetNamespaceId(),
		TimeoutMs:           nconf.GetClientConfig().GetTimeoutMs(),
		ListenInterval:      nconf.GetClientConfig().GetListenInterval(),
		BeatInterval:        nconf.GetClientConfig().GetBeatInterval(),
		NotLoadCacheAtStart: nconf.GetClientConfig().GetNotCache(),
		LogDir:              nconf.GetClientConfig().GetLogDir(),
		CacheDir:            nconf.GetClientConfig().GetCacheDir(),
		LogLevel:            nconf.GetClientConfig().GetLogLevel(),
	}

	// 创建 Nacos 客户端
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}
	clientParam := nconf.GetClientParam()
	sourceArr := make([]config.Source, 0, len(clientParam))
	for _, v := range clientParam {
		sourceArr = append(sourceArr,
			knacos.NewConfigSource(
				client,
				knacos.WithGroup(v.GetGroup()),
				knacos.WithDataID(v.GetDataId()),
			))
	}
	return sourceArr
}
