package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type CreateProductDto struct {
	Protocol    string `json:"protocol" validate:"required"`
	DataFormat  string `json:"data_format" validate:"required"`
	NodeType    int    `json:"node_type"  validate:"min=0,max=3"`
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Classify    string `json:"product_type"`
	Description string `json:"description"`
	Network     string `json:"network"`
}

type UpdateProductDto struct {
	ProductKey  string `json:"productkey" validate:"required"`
	Name        string `json:"name" validate:"required,min=1,max=30"`
	Protocol    string `json:"protocol" validate:"required"`
	DataFormat  string `json:"data_format" validate:"required"`
	Classify    string `json:"product_type"`
	Description string `json:"description"`
	Network     string `json:"network"`
	NodeType    int    `json:"node_type" validate:"min=0,max=3"`
}

type Product struct {
	ProjectKey  string     `json:"projectkey"`
	ProductKey  string     `json:"productkey"`
	Name        string     `json:"name"`
	Img         string     `json:"img"`
	ClassIf     string     `json:"classif"`
	Template    string     `json:"template"`
	Network     string     `json:"network"`
	Protocol    string     `json:"protocol"`
	DataFormat  string     `json:"dataformat"`
	Description string     `json:"description"`
	UpdatedAt   *time.Time `json:"updatedat,omitempty"`
	CreatedAt   *time.Time `json:"createdat,omitempty"`
	ActivedAt   *time.Time `json:"activedat,omitempty"`
	Secret      string     `json:"secret"`
	DeviceTotal int        `json:"devicetotal"`
	NodeType    int        `json:"nodetype,omitempty" copier:"-"`
}

type ListProductDto struct {
	Data []Product
	common.BaseResponse
}
