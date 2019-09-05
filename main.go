package main

import (
	"fmt"
	"green_go_sdk/sdk"
)

const accessKeyId string = "xxx"
const accessKeySecret string = "xxx"


func main(){
	profile := sdk.Profile{AccessKeyId: accessKeyId, AccessKeySecret:accessKeySecret}

	path := "/green/text/scan"

	clientInfo := sdk.ClinetInfo{}

	// 构造请求数据
	bizType := "Green"
	scenes := []string{"antispam"}

	task := sdk.Task{DataId: sdk.Rand().Hex(), Content:"sds"}
	tasks := []sdk.Task{task}

	bizData := sdk.BizData{bizType, scenes, tasks}

	var client sdk.IAliYunClient = sdk.DefaultClient{Profile: profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, bizData))

}

