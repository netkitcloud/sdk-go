package seanet

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
)

// 获取指定产品
func (c *SeanetClient) GetProduct(productKey string) (resp dto.ProductDto, err error) {
	if productKey == "" {
		err = errors.New("productKey is required")
		return
	}

	uri := fmt.Sprintf(apiModifyProduct, productKey)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取产品列表
func (c *SeanetClient) ListProduct(param common.PaginationParams) (resp dto.ListProductDto, err error) {
	if err = c.validate.Struct(param); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiProduct, http.MethodGet, param)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
