package nip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nip/dto"
	"github.com/valyala/fastjson"
)

func (c *NIPClient) CreateProduct(dto *dto.CreateProductDto) (*dto.Product, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	body, err := c.SendHttpRequest(apiProduct, http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	return c.responseProduct(body)
}

func (c *NIPClient) DeleteProduct(productKey string) (*common.BaseResponse, error) {
	if productKey == "" {
		return nil, errors.New("productKey is required")
	}

	uri := fmt.Sprintf(apiModifyProduct, productKey)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *NIPClient) GetProduct(productKey string) (*dto.Product, error) {
	if productKey == "" {
		return nil, errors.New("productKey is required")
	}

	uri := fmt.Sprintf(apiModifyProduct, productKey)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseProduct(body)
}

func (c *NIPClient) UpdateProduct(dto *dto.UpdateProductDto) (*dto.Product, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(apiModifyProduct, dto.ProductKey)
	body, err := c.SendHttpRequest(uri, http.MethodPut, dto)
	if err != nil {
		return nil, err
	}

	return c.responseProduct(body)
}

func (c *NIPClient) ListProduct(pagination common.PaginationParams) (productListResp dto.ListProductDto, err error) {
	body, err := c.SendHttpRequest(apiProduct, http.MethodGet, pagination)
	if err != nil {
		return
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return
		}
		err = errors.New(string(msg))
		return
	}

	err = json.Unmarshal(body, &productListResp)
	if err != nil {
		return
	}

	if !productListResp.Status {
		err = fmt.Errorf("code : %d", productListResp.Code)
	}

	return
}

func (c *NIPClient) responseProduct(b []byte) (*dto.Product, error) {
	var p fastjson.Parser
	v, err := p.Parse(string(b))
	if err != nil {
		return nil, err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(msg))
	}

	byteProduct := v.GetObject("data").MarshalTo(nil)
	resultProduct := dto.Product{}
	err = json.Unmarshal(byteProduct, &resultProduct)
	if err != nil {
		return nil, err
	}

	return &resultProduct, nil
}
