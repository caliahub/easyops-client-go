package easyops_client_go

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	netUrl "net/url"
	"strconv"
	"time"
)

/**
 * @Description easyops client
 * @Author 梁泽华Calia
 * @Date 2023/5/24 15:57
 **/

type Config struct {
	AccessKey *string
	SecretKey *string
	Host      *string
}

type Response struct {
	Code        *int        `json:"code"`
	CodeExplain *string     `json:"codeExplain"`
	Error       *string     `json:"error"`
	Data        interface{} `json:"data"`
}

func String(s string) *string {
	return &s
}

func StringValue(a *string) string {
	if a == nil {
		return ""
	}
	return *a
}

// 签名
func (config *Config) signature(method, url, contentType string, expires int64, data map[string]interface{}) (string, error) {
	var params, contentMd5 string

	if contentType == "" {
		contentType = "application/json"
	}

	if method == http.MethodPost || method == http.MethodPut {
		ds, err := json.Marshal(data)
		if err != nil {
			return "", err
		}

		hash := md5.Sum([]byte(ds))
		contentMd5 = hex.EncodeToString(hash[:])
	}

	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%d\n%s", method, url, params, contentType, contentMd5, expires, StringValue(config.AccessKey))

	// 生成签名
	h := hmac.New(sha1.New, []byte(StringValue(config.SecretKey)))
	h.Write([]byte(signStr))
	hashed := h.Sum(nil)
	signature := hex.EncodeToString(hashed)

	return signature, nil
}

// Do 请求
func (config *Config) Do(method, url string, data map[string]interface{}) (string, error) {
	// 签名
	expires := time.Now().Unix()
	signature, err := config.signature(method, url, "", expires, data)
	if err != nil {
		return "", err
	}

	// URL
	if StringValue(config.Host) == "" {
		return "", errors.New("param err: host is nil")
	}
	Url, err := netUrl.Parse(fmt.Sprintf("https://%s%s", StringValue(config.Host), url))
	if err != nil {
		return "", err
	}
	params := netUrl.Values{}
	params.Add("accesskey", StringValue(config.AccessKey))
	params.Add("signature", signature)
	params.Add("expires", strconv.FormatInt(expires, 10))
	Url.RawQuery = params.Encode()
	url = Url.String()

	bytesData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, url, bytes.NewReader(bytesData))
	if err != nil {
		return "", err
	}

	// 请求头
	req.Host = "openapi.easyops-only.com"
	req.Header.Add("Content-Type", "application/json")

	// 响应
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}

	res, err := io.ReadAll(resp.Body)

	return string(res), err
}
