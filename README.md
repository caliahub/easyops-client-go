## easyops-go-sdk

[![Go Report Card](https://pkg.go.dev/badge/github.com/prometheus/client_golang.svg)](https://github.com/caliahub/easyops-client-go)

优维DevOps平台Go SDK，简单易用，方便Go开发人员通过openapi快速接入Easyops

### 快速开始
#### 引入依赖包
```bash
go get github.com/caliahub/easyops-client-go@v0.0.1
```

#### 使用（示例）
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

### 接口文档
- [流程实例接口](flowable/process_instance.md)
