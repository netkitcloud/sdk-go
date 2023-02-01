package nauth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"

	"net/http"
	"net/url"

	"github.com/valyala/fastjson"
)

type AuthenticationClient struct {
	options *AuthenticationClientOptions
	ClientUser *User
	ClientToken string
}

func NewClient(options *AuthenticationClientOptions) (*AuthenticationClient, error) {
	if options.Host == "" {
		options.Host = CoreAuthApiHost
	}

	return &AuthenticationClient{
		options: options,
	}, nil
}

func (c *AuthenticationClient) SetCurrentUser(user *User) (*User, error) {
	c.ClientUser = user
	if len(user.Token) > 0 {
		c.ClientToken = user.Token
	}
	return user, nil
}

func (c *AuthenticationClient) GetUserByToken(token string) (*User, error) {
	c.ClientToken = token
	body, err := c.SendHttpRequest("/user", http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	return c.responseGetUser(body)
}

func (c *AuthenticationClient) responseGetUser(b []byte) (*User, error) {
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
	resultUser := User{}
	err = json.Unmarshal(byteUser, &resultUser)
	if err != nil {
		return nil, err
	}
	c.SetCurrentUser(&resultUser)
	return &resultUser, nil
}

func (c *AuthenticationClient) SendHttpRequest(requestUrl string, method string, reqDto interface{}) ([]byte, error) {
	data, _ := json.Marshal(&reqDto)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)

	if method == http.MethodGet && variables != nil && len(variables) > 0 {
		params := url.Values{}
		urlTemp, _ := url.Parse(requestUrl)
		for key, value := range variables {
			params.Set(key, fmt.Sprintf("%v", value))
		}
		urlTemp.RawQuery = params.Encode()
		requestUrl = urlTemp.String()
	}

	uri := c.options.Host + requestUrl

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("x-nauth-sdk-version", Version)
	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	if len(c.options.Tenant) > 0 {
		req.Header.Add("x-nauth-tenant", c.options.Tenant)
	}
	if len(c.ClientToken) > 0 {
		req.Header.Add("Authorization", "Bearer "+c.ClientToken)
	}

	for key, value := range commonHeaders {
		req.Header.Add(key, value)
	}

	if method == http.MethodPost {
		reqData, err := json.Marshal(reqDto)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(reqData))
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
		return nil, errors.New("http status: "+strconv.FormatInt(int64(resp.StatusCode), 10))
	}

	return body, nil
}
