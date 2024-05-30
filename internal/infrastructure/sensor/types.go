package sensor

/**
* File name types.go
*
* @author  Li Chuang
* @date    2024-01-04 17:18:45
**/

type EventType string

const (
	EventTypeOrderPayed EventType = "paid" //订单支付成功
)

func (t EventType) String() string {
	return string(t)
}

type PropertiesType interface {
	Encode() map[string]interface{}
}

type Tracking struct {
	DistinctId string
	Event      EventType
	Properties PropertiesType
	IsLoginId  bool
}

type PropertiesOrderPayed struct {
	OrderId string   `json:"order_id"`
	SpuId   []string `json:"spu_id"`
	SkuId   []string `json:"sku_id"`
	TalId   string   `json:"tal_id"`
}

func (t *PropertiesOrderPayed) Encode() map[string]interface{} {
	//此处不要直接使用json.Marshal和json.Unmarshal反解，只有不包含[]string字段类型的struct才能使用json.Marshal和json.Unmarshal
	//因为json.Unmarshal之后，spu_id和sku_id的字段类型会由[]string变为[]interface{}，而神策不支持此字段类型
	return map[string]interface{}{
		"order_id": t.OrderId,
		"spu_id":   t.SpuId,
		"sku_id":   t.SkuId,
		"tal_id":   t.TalId,
	}
}
