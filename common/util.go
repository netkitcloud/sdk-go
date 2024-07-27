package common

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fastjson"
)

// 根据传入的byte解析成所需的结构体, 并进行code校验
func ParserDto[T any](body []byte, dto *T) (err error) {
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return
	}
	// fmt.Println("body:", string(body))

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		code := v.GetInt("code")
		err = errors.New(string(msg))
		fmt.Printf("request status is false, error code is %d, error msg is %s.\n", code, msg)
		return
	}

	return
}
