package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type Product struct {
	Id                     uint                     `json:"id"`
	Name                   string                   `json:"name"`
	ProductKey             string                   `json:"product_key"`
	IotPlatformId          uint                     `json:"iot_platform_id"`
	DeviceConfig           map[string]interface{}   `json:"device_config"`
	Property               map[string]interface{}   `json:"property"`
	CfgTemplate            map[string]interface{}   `json:"cfg_template"`
	CfgFeature             []map[string]interface{} `json:"feature"`
	ConfigPostSize         uint                     `json:"config_post_size"`
	PropertyInterval       uint                     `json:"property_interval"`
	MqttInterval           uint                     `json:"mqtt_interval"`
	ConfigVersionDifferent bool                     `json:"config_version_different"`
	ProductCode            uint                     `json:"product_code"`
	Createdat              time.Time                `json:"createdat"`
	Updatedat              time.Time                `json:"updatedat"`
	Img                    string                   `json:"img"`
	ProtocolType           uint8                    `json:"protocol_type"`
	AllowUserCreate        bool                     `json:"allow_user_create"`
}

type ProductDto struct {
	Data Product
	common.BaseResponse
}

type ListProductDto struct {
	Data []Product
	common.BaseListResponse
}
