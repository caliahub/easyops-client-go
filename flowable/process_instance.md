## 流程实例接口

- [获取流程实例详情](#获取流程实例详情)
- [发起流程实例](#发起流程实例)

### 获取流程实例详情
```go
package main

import (
	"fmt"
	openapi "github.com/caliahub/easyops-client-go"
	"github.com/caliahub/easyops-client-go/flowable"
)

func main() {
	// config
	config := &openapi.Config{
		AccessKey: openapi.String("xxx"),
		SecretKey: openapi.String("xxx"),
		Host:      openapi.String("xxx"),
	}
	
	// create client 
	client := flowable.NewClient(config)
	
	// request
	resp, err := client.GetProcessInstance("xxx")
	if err != nil {
		return
	}
	
	fmt.Println(resp)
}
```

### 发起流程实例
```go
package main

import (
	"fmt"
	openapi "github.com/caliahub/easyops-client-go"
	"github.com/caliahub/easyops-client-go/flowable"
	"github.com/caliahub/easyops-client-go/flowable/model"
)

func main() {
	// config
	config := &openapi.Config{
		AccessKey: openapi.String("xxx"),
		SecretKey: openapi.String("xxx"),
		Host:      openapi.String("xxx"),
	}
	
	// create client 
	client := flowable.NewClient(config)

	// params
	handlers := []string{"xxx"}
	m := &model.StartProcessInstanceRequestBody{
		ITSCInfluenceScope:  "",
		ITSCUrgency:         "",
		AssigneeGroups:      nil,
		AssigneeUsers:       handlers,
		DepartmentId:        "",
		FormData:            "",
		HandleWay:           "",
		IsSupervision:       false,
		Name:                "发起流程实例",
		RelevanceInstanceId: "",
		ServiceId:           "5fb6254c64822",
		Source:              "",
		SubsequentAssignee:  nil,
		SupervisorUserList:  nil,
		VariableName:        "",
		VariableValue:       "",
		VisibleRange:        "operator",
	}

	// request
	resp, err := client.CreateProcessInstance(m)
	if err != nil {
		return
	}
	
	fmt.Println(resp)
}
```
