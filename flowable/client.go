package flowable

import openapi "github.com/caliahub/easyops-client-go"

/**
 * @Description
 * @Author 梁泽华Calia
 * @Date 2023/5/24 16:47
 **/

type Client struct {
	openapi.Config
}

func NewClient(config *openapi.Config) *Client {
	return &Client{Config: *config}
}
