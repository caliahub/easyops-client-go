package easyops_client_go

import (
	"net/http"
	"testing"
	"time"
)

/**
 * @Description
 * @Author 梁泽华Calia
 * @Date 2023/5/24 15:58
 **/

func TestSignature(t *testing.T) {
	config := &Config{
		AccessKey: String("xxx"),
		SecretKey: String("xxx"),
		Host:      String("xxx"),
	}

	data := make(map[string]interface{}, 0)
	data["name"] = "xxxx"
	data["serviceId"] = "xxxx"
	data["visibleRange"] = "xxx"

	expires := time.Now().Unix()

	config.signature(http.MethodGet, "/flowable_service/api/flowable_service/v1/process_instance/xxx", "", expires, data)
}

func TestDo(t *testing.T) {
	config := &Config{
		AccessKey: String("xxx"),
		SecretKey: String("xxx"),
		Host:      String("xxx"),
	}

	data := make(map[string]interface{}, 0)
	data["name"] = "xxx"
	data["serviceId"] = "xxx"
	data["visibleRange"] = "xxx"

	config.Do(http.MethodGet, "/flowable_service/api/flowable_service/v1/process_instance/xxx", data)
}
