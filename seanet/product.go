package seanet

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/valyala/fastjson"
)

func (c *SeanetClient) GetProduct(productKey string) (*dto.Product, error) {
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

func (c *SeanetClient) ListProduct(pagination common.PaginationParams) (productListResp dto.ListProductDto, err error) {
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

func (c *SeanetClient) responseProduct(b []byte) (*dto.Product, error) {
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
