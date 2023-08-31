package nauth

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
)

type AuthenticationAdmin struct {
	options     *AuthenticationClientOptions
	ClientUser  *dto.User
}

func NewAdmin(options *AuthenticationClientOptions) (*AuthenticationAdmin, error) {
	if options.Host == "" {
		options.Host = CoreAuthApiHost
	}

	if options.Tenant == "" {
		return nil, errors.New("tenantId is required")
	}

	if options.Secret == "" {
		return nil, errors.New("secret is required")
	}

	return &AuthenticationAdmin{
		options: options,
	}, nil
}

func (c *AuthenticationAdmin) SendHttpRequest(requestUrl string, method string, reqDto interface{}) ([]byte, error) {
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

	for key, value := range commonHeaders {
		req.Header.Add(key, value)
	}

	req.Header.Add("x-n-ts", strconv.FormatInt(time.Now().Unix(), 10))
	rand.Seed(time.Now().UnixNano())
	req.Header.Add("x-n-nonce", strconv.FormatInt(int64(rand.Intn(100000000)), 16))

	if method == http.MethodGet || reqDto == nil {
		variables := make(map[string]interface{})
		json.Unmarshal(data, &variables)

		querys := url.Values{}
		if len(variables) > 0 {
			for key, value := range variables {
				querys.Set(key, fmt.Sprintf("%v", value))
			}
			req.URL.RawQuery = querys.Encode()
		}

		sig := common.SignatureRequestGet(req, querys, []byte(c.options.Secret))
		req.Header.Add("x-n-signature", sig)

	} else if method != http.MethodGet && reqDto != nil {
		reqData, err := json.Marshal(reqDto)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(reqData))
		sig := common.SignatureRequestBody(req, reqDto, []byte(c.options.Secret))
		req.Header.Add("x-n-signature", sig)
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
