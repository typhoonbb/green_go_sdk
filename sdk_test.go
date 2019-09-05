package green_go_sdk

import (
	"fmt"
	"github.com/typhoonbb/green_go_sdk/green_client"
	"github.com/typhoonbb/green_go_sdk/udid"
	"testing"
)

const accessKeyId string = "xxx"
const accessKeySecret string = "xxx"


func TestTextScan(t *testing.T) {
	profile := green_client.Profile{AccessKeyId: accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/text/scan"

	clientInfo := green_client.ClientInfo{}

	// 构造请求数据
	bizType := "Green"
	scenes := []string{"antispam"}

	task := green_client.Task{DataId: udid.Rand().Hex(), Content:"sds"}
	tasks := []green_client.Task{task}

	bizData := green_client.BizData{bizType, scenes, tasks}

	var client green_client.IAliYunClient = green_client.DefaultClient{Profile: profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, bizData))

}

