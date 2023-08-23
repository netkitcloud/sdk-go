package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type CreateDeviceDto struct {
	Productkey  string `json:"productkey" validate:"required"`
	Devicekey   string `json:"devicekey" validate:"required"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Latlng      string `json:"latlng"`
	Description string `json:"description"`
}

type GetDeviceDto struct {
	Productkey string `json:"productkey" validate:"required"`
	Devicekey  string `json:"devicekey" validate:"required"`
}

type UpdateDeviceDto struct {
	Productkey  string `json:"productkey" validate:"required"`
	Devicekey   string `json:"devicekey" validate:"required"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Latlng      string `json:"latlng"`
	Description string `json:"description"`
}
type Device struct {
	// product
	ProductName string `json:"productname"`
	Productkey  string `json:"productkey"`
	ProductImg  string `json:"img"`
	Classif     string `json:"classif"`
	Network     string `json:"network"`
	Protocol    string `json:"protocol"`
	Dataformat  string `json:"dataformat"`
	// device
	Name                string     `json:"name"`
	Devicekey           string     `json:"devicekey"`
	Address             string     `json:"address"`
	Secret              string     `json:"secret"`
	Latlng              string     `json:"latlng"`
	Description         string     `json:"description"`
	Online              int        `json:"online"`
	Activate            int        `json:"activate"`
	Lastcommunicationat *time.Time `json:"lastcommunicationat,omitempty"`
	Activatedat         *time.Time `json:"activatedat,omitempty"`
	UpdatedAt           *time.Time `json:"updatedat,omitempty"`
	CreatedAt           *time.Time `json:"createdat,omitempty"`
	NodeType            int        `json:"nodetype,omitempty" copier:"-"`
}

type ListDeviceDto struct {
	Data []Device
	common.BaseResponse
}
