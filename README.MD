## Aliyun green go sdk

### Installation

`go get github.com/typhoonbb/green_go_sdk`

### Useage
````
profile := green_client.Profile{AccessKeyId: "your accessid", AccessKeySecret:"your key"}

	path := "/green/text/scan"

	clientInfo := green_client.ClientInfo{}

	// request data
	bizType := "Green"
	scenes := []string{"antispam"}

	task := green_client.Task{DataId: udid.Rand().Hex(), Content:"your content..."}
	tasks := []green_client.Task{task}

	bizData := green_client.BizData{bizType, scenes, tasks}

	var client green_client.IAliYunClient = green_client.DefaultClient{Profile: profile}

	// your biz code
	fmt.Println(client.GetResponse(path, clientInfo, bizData))


````
