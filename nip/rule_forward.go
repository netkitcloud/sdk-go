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

func (c *NIPClient) GetRuleForward(ruleKey string) (*dto.RuleForward, error) {
	if ruleKey == "" {
		return nil, errors.New("ruleKey is required")
	}

	uri := fmt.Sprintf(apiModifyRuleForward, ruleKey)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseRule(body)
}

func (c *NIPClient) CreateRuleForward(dto *dto.CreateRuleForwardDto) (*common.BaseResponse, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	body, err := c.SendHttpRequest(apiRuleForward, http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *NIPClient) DeleteRuleForward(ruleKey string) (*common.BaseResponse, error) {
	if ruleKey == "" {
		return nil, errors.New("ruleKey is required")
	}

	uri := fmt.Sprintf(apiModifyRuleForward, ruleKey)
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

func (c *NIPClient) UpdateRuleForward(dto *dto.UpdateRuleForwardDto) (*common.BaseResponse, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(apiModifyRuleForward, dto.ID)
	body, err := c.SendHttpRequest(uri, http.MethodPut, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 分页获取规则
func (c *NIPClient) ListRuleForward(pagination common.PaginationParams) (ruleListResp dto.ListRuleForwardDto, err error) {
	body, err := c.SendHttpRequest(apiRuleForward, http.MethodGet, pagination)
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

	err = json.Unmarshal(body, &ruleListResp)
	if err != nil {
		return
	}

	if !ruleListResp.Status {
		err = fmt.Errorf("code : %d", ruleListResp.Code)
	}

	return
}

func (c *NIPClient) responseRule(b []byte) (*dto.RuleForward, error) {
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

	byteRule := v.GetObject("data").MarshalTo(nil)
	resultRule := dto.RuleForward{}
	err = json.Unmarshal(byteRule, &resultRule)
	if err != nil {
		return nil, err
	}

	return &resultRule, nil
}
