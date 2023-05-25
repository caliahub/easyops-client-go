package flowable

import (
	openapi "github.com/caliahub/easyops-client-go"
	"github.com/caliahub/easyops-client-go/flowable/model"
	"testing"
)

/**
 * @Description
 * @Author 梁泽华Calia
 * @Date 2023/5/24 16:24
 **/

func TestGetProcessInstance(t *testing.T) {
	config := &openapi.Config{
		AccessKey: openapi.String("xxx"),
		SecretKey: openapi.String("xxx"),
		Host:      openapi.String("xxx"),
	}
	client := NewClient(config)
	client.GetProcessInstance("xxx")
}

func TestCreateProcessInstance(t *testing.T) {
	config := &openapi.Config{
		AccessKey: openapi.String("xxx"),
		SecretKey: openapi.String("xxx"),
		Host:      openapi.String("xxx"),
	}
	client := NewClient(config)

	handlers := []string{"zehua.liang"}
	m := &model.StartProcessInstanceRequestBody{
		ITSCInfluenceScope:  "",
		ITSCUrgency:         "",
		AssigneeGroups:      nil,
		AssigneeUsers:       handlers,
		DepartmentId:        "",
		FormData:            "",
		HandleWay:           "",
		IsSupervision:       false,
		Name:                "xxx",
		RelevanceInstanceId: "",
		ServiceId:           "xxx",
		Source:              "",
		SubsequentAssignee:  nil,
		SupervisorUserList:  nil,
		VariableName:        "",
		VariableValue:       "",
		VisibleRange:        "operator",
	}

	client.CreateProcessInstance(m)
}
