package sensor

/**
* File name client.go
*
* @author  Li Chuang
* @date    2024-01-04 15:32:52
**/
import (
	"be_infra_tpl/internal/conf"
	"os"
	"path/filepath"

	sdk "github.com/sensorsdata/sa-sdk-go"
)

type SensorAgent struct {
	XPadStoreAnalytics        *sdk.SensorsAnalytics
	XPadStoreLoggingAnalytics *sdk.SensorsAnalytics
}

func NewSensorAgent(conf *conf.Bootstrap) *SensorAgent {
	sa := new(SensorAgent)
	sa.loadXPadStoreAnalytics(conf)
	return sa
}

func (c *SensorAgent) mkdir(filePath string) (err error) {
	dir := filepath.Dir(filePath)
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, 0755)
	}
	return
}

func (c *SensorAgent) loadXPadStoreAnalytics(conf *conf.Bootstrap) {
	// if c.XPadStoreAnalytics != nil {
	// 	return
	// }
	// //判断是否开启神策
	// if len(conf.Common.Sensor.XpadstoreFilePath) == 0 {
	// 	return
	// }
	// //判断文件是否存在
	// c.mkdir(conf.Common.Sensor.XpadstoreFilePath)
	// // xPadStoreConsumer, err := sdk.InitConcurrentLoggingConsumer(conf.Common.Sensor.XpadstoreFilePath, false)
	// // xPadStoreConsumer, err := sdk.InitDebugConsumer(conf.Common.Sensor.ServerUrl, true, 5000)
	// xPadStoreConsumer, err := sdk.InitDefaultConsumer(conf.Common.Sensor.ServerUrl, 5000)
	// if err != nil {
	// 	panic(err)
	// }
	// sa := sdk.InitSensorsAnalytics(xPadStoreConsumer, conf.Common.Sensor.ProjectName, false)
	// c.XPadStoreAnalytics = &sa
}

func (c *SensorAgent) Close() {
	if c.XPadStoreAnalytics != nil {
		c.XPadStoreAnalytics.Close()
	}
	if c.XPadStoreLoggingAnalytics != nil {
		c.XPadStoreLoggingAnalytics.Close()
	}
}

func (c *SensorAgent) TrackXPadStore(tracking Tracking) (err error) {
	if c.XPadStoreAnalytics == nil {
		return
	}
	pp := tracking.Properties.Encode()
	err = c.XPadStoreAnalytics.Track(tracking.DistinctId, tracking.Event.String(), pp, tracking.IsLoginId)

	//记录日志
	if c.XPadStoreLoggingAnalytics != nil {
		err = c.XPadStoreLoggingAnalytics.Track(tracking.DistinctId, tracking.Event.String(), pp, tracking.IsLoginId)
		c.XPadStoreLoggingAnalytics.Flush()
	}

	return
}
