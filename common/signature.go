package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func SignatureString(str string, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	hash := h.Sum(nil)
	result := fmt.Sprintf("%x", hash)
	return result
}

func SignatureSortMap(requestMap map[string]interface{}, key []byte) string {
	keys := make([]string, 0, len(requestMap))
	for k := range requestMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var str []string
	for _, k := range keys {
		str = append(str, fmt.Sprintf("%s=%v", k, requestMap[k]))
	}
	return SignatureString(strings.Join(str, "&"), key)
}

func SignatureStruct(d interface{}, key []byte) string {
	data, _ := json.Marshal(&d)
	variables := make(map[string]interface{})
	json.Unmarshal(data, &variables)
	return SignatureSortMap(variables, key)
}

func SignatureRequestBody(req *http.Request, d interface{}, key []byte) string {
	var str []string

	data, _ := json.Marshal(&d)
	requestMap := make(map[string]interface{})
	json.Unmarshal(data, &requestMap)
	keys := make([]string, 0, len(requestMap))
	for k := range requestMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	str = append(str, strings.Replace(req.URL.Path, "/api", "", 1))
	str = append(str, fmt.Sprint("x-n-ts=", req.Header.Get("x-n-ts")))
	str = append(str, fmt.Sprint("x-n-nonce=", req.Header.Get("x-n-nonce")))

	for _, k := range keys {
		str = append(str, fmt.Sprintf("%s=%v", k, requestMap[k]))
	}
	return SignatureString(strings.Join(str, "&"), key)

}

func SignatureRequestGet(req *http.Request, querys url.Values, key []byte) string {
	var keys []string
	for key, _ := range querys {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var str []string

	str = append(str, strings.Replace(req.URL.Path, "/api", "", 1))
	str = append(str, fmt.Sprint("x-n-ts=", req.Header.Get("x-n-ts")))
	str = append(str, fmt.Sprint("x-n-nonce=", req.Header.Get("x-n-nonce")))

	for _, key := range keys {
		str = append(str, fmt.Sprintf("%s=%v", key, querys.Get(key)))
	}
	return SignatureString(strings.Join(str, "&"), key)
}
