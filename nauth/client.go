package nauth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	"net/http"
	"net/url"

	"github.com/netkitcloud/sdk-go/nauth/dto"

	"github.com/valyala/fastjson"
)

type AuthenticationClient struct {
	options     *AuthenticationClientOptions
	ClientUser  *dto.User
	AccessToken string
}

const (
	ClientSecretPost  = "client_secret_post"
	ClientSecretBasic = "client_secret_basic"
	None              = "none"
	CurrentUserToken  = "UserToken"
)

func NewClient(options *AuthenticationClientOptions) (*AuthenticationClient, error) {
	if options.Host == "" {
		options.Host = CoreAuthApiHost
	}

	if options.AppId == "" {
		return nil, errors.New("AppId is required")
	}

	if options.Tenant == "" {
		return nil, errors.New("TenantId is required")
	}

	if options.AppSecret == "" {
		return nil, errors.New("AppSecret is required")
	}

	return &AuthenticationClient{
		options: options,
	}, nil
}

func (c *AuthenticationClient) SetToken(token string) error {
	if token == "" {
		return errors.New("token cannot be empty")
	}
	c.AccessToken = token
	return nil
}

func (c *AuthenticationClient) SetCurrentUser(user *dto.User) (*dto.User, error) {
	c.ClientUser = user
	if len(user.AccessToken) > 0 {
		c.SetToken(user.AccessToken)
	}
	return user, nil
}

func (c *AuthenticationClient) responseGetUser(b []byte) (*dto.User, error) {
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

	byteUser := v.GetObject("data").MarshalTo(nil)
	resultUser := dto.User{}
	err = json.Unmarshal(byteUser, &resultUser)
	if err != nil {
		return nil, err
	}

	c.SetCurrentUser(&resultUser)
	return &resultUser, nil
}

func (c *AuthenticationClient) responseError(body []byte) error {
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil
}

func (c *AuthenticationClient) SendHttpRequest(requestUrl string, method string, reqDto interface{}) ([]byte, error) {
	data, _ := json.Marshal(&reqDto)

	req, err := http.NewRequest(method, c.options.Host+requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	if len(c.options.Tenant) > 0 {
		req.Header.Add("x-nauth-tenant", c.options.Tenant)
	}
	if len(c.options.AppId) > 0 {
		req.Header.Add("x-nauth-app", c.options.AppId)
	}

	if len(c.AccessToken) > 0 {
		req.Header.Add("Authorization", "Bearer "+c.AccessToken)
	}

	for key, value := range commonHeaders {
		req.Header.Add(key, value)
	}

	if method != http.MethodGet && reqDto != nil {
		reqData, err := json.Marshal(reqDto)
		if err != nil {
			return nil, err
		}
		req.Body = io.NopCloser(bytes.NewReader(reqData))
	} else if method == http.MethodGet || reqDto == nil {
		variables := make(map[string]interface{})
		json.Unmarshal(data, &variables)

		querys := url.Values{}
		if len(variables) > 0 {
			for key, value := range variables {
				querys.Set(key, fmt.Sprintf("%v", value))
			}
			req.URL.RawQuery = querys.Encode()
		}
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("http status: " + strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	return body, nil
}

// TO USE: cli.NewClientWithToken(token).AddAccessKey
// 根据token新建一个客户端示例
func (c *AuthenticationClient) NewClientWithToken(token string) *AuthenticationClient {
	if token == "" {
		return nil
	}
	options := *c.options
	newClinet, err := NewClient(&options)
	if err != nil {
		return nil
	}
	err = newClinet.SetToken(token)
	if err != nil {
		return nil
	}
	return newClinet
}
