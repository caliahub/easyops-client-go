package flowable

import (
	"encoding/json"
	"errors"
	openapi "github.com/caliahub/easyops-client-go"
	"github.com/caliahub/easyops-client-go/flowable/model"
	"net/http"
)

/**
 * @Description 流程实例
 * @Author 梁泽华Calia
 * @Date 2023/5/24 16:07
 **/

const (
	baseV1Url = "/flowable_service/api/flowable_service/v1/process_instance"
)

// GetProcessInstance 查询流程实例详情
func (client *Client) GetProcessInstance(instanceId string) (*model.GetProcessInstanceResponseBody, error) {
	url := baseV1Url + "/" + instanceId
	resp, err := client.Do(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	getProcessInstanceResponseBody := &model.GetProcessInstanceResponseBody{}
	response := &openapi.Response{
		Data: getProcessInstanceResponseBody,
	}

	err = json.Unmarshal([]byte(resp), response)
	if err != nil {
		return nil, err
	}

	if response.Code != nil && *response.Code != 0 {
		return nil, errors.New(*response.Error)
	}

	return getProcessInstanceResponseBody, err
}

// CreateProcessInstance 发起流程实例
func (client *Client) CreateProcessInstance(req *model.StartProcessInstanceRequestBody) (*model.StartProcessInstanceResponseBody, error) {
	byteReq, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{}, 0)
	if err = json.Unmarshal(byteReq, &data); err != nil {
		return nil, err
	}

	resp, err := client.Do(http.MethodPost, baseV1Url, data)

	if err != nil {
		return nil, err
	}

	startProcessInstanceResponseBody := &model.StartProcessInstanceResponseBody{}
	response := &openapi.Response{
		Data: startProcessInstanceResponseBody,
	}
	err = json.Unmarshal([]byte(resp), response)
	if err != nil {
		return nil, err
	}

	if response.Code != nil && *response.Code != 0 {
		return nil, errors.New(*response.Error)
	}

	return startProcessInstanceResponseBody, nil
}
