package nip

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"time"

	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
	"github.com/netkitcloud/sdk-go/common"

	"github.com/valyala/fastjson"
)

type NIPClient struct {
	options  *NIPClientOptions
	validate *validator.Validate
}

type NIPClientOptions struct {
	Host         string
	AccessKey    string
	AccessSecret string
	ProjectKey   string
}

const (
	ClientSecretPost  = "client_secret_post"
	ClientSecretBasic = "client_secret_basic"
	None              = "none"
)

func NewClient(options *NIPClientOptions) (*NIPClient, error) {
	if options.Host == "" {
		options.Host = CoreAuthApiHost
	}

	if options.AccessKey == "" {
		return nil, errors.New("accesskey is required")
	}

	if options.AccessSecret == "" {
		return nil, errors.New("accesssecret is required")
	}

	if options.ProjectKey == "" {
		return nil, errors.New("projectkey is required")
	}

	validate := validator.New()
	return &NIPClient{
		options:  options,
		validate: validate,
	}, nil
}

func (c *NIPClient) responseError(body []byte) error {
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

func (c *NIPClient) SendHttpRequest(requestUrl string, method string, reqDto interface{}) ([]byte, error) {
	data, _ := json.Marshal(&reqDto)

	req, err := http.NewRequest(method, c.options.Host+requestUrl, nil)
	if err != nil {
		return nil, err
	}

	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/json;charset=UTF-8")
	}

	if len(c.options.ProjectKey) > 0 {
		req.Header.Add("x-nip-project", c.options.ProjectKey)
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

		sig := common.SignatureRequestGet(req, querys, []byte(c.options.AccessSecret))
		req.Header.Add("x-n-accesskey", c.options.AccessKey)
		req.Header.Add("x-n-signature", sig)

	} else if method != http.MethodGet && reqDto != nil {
		reqData, err := json.Marshal(reqDto)
		if err != nil {
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewReader(reqData))
		sig := common.SignatureRequestBody(req, reqDto, []byte(c.options.AccessSecret))
		req.Header.Add("x-n-accesskey", c.options.AccessKey)
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
