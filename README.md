## 抖店开放平台GO SDK

在官方SDK的基础上，集成了自己整合的消息推送SDK，方便开发者更快速的对接。

当前项目依赖的官方SDK版本号：doudian-sdk-golang-1.1.0-20230221172709，后续官方SDK版本升级了，我们可直接下载官方最新SDK对本项目进行覆盖升级。

### 目录结构

```sh
.
├── LICENSE
├── README.md
├── api             # 纯官方SDK代码，可覆盖升级
├── core            # 纯官方SDK代码，可覆盖升级
├── errors          # 纯官方SDK代码，可覆盖升级
├── example.go      # 纯官方SDK代码，可覆盖升级
├── example_test.go # 纯官方SDK代码，可覆盖升级
├── go.mod          # 纯官方SDK代码，可覆盖升级
├── message         # 自己整合的消息推送SDK，不可删除，缺失的消息可自行补上
├── spi             # 纯官方SDK代码，可覆盖升级
└── utils           # 纯官方SDK代码，可覆盖升级
```

### 安装教程

本SDK不支持 go get 方式，仅支持本地导入的方式，安装步骤如下：

1. 下载SDK源代码

```sh
git clone https://github.com/zsmhub/doudian-sdk.git
```

2. 将源码文件夹拷贝到自己的项目工程中，路径任意，例如lib文件夹下：

![img](https://zsmhub.github.io/images/2023/WX20230225-094247.png)

3. 打开自己项目工程的go.mod文件：

在required模块添加「doudian.com/open/sdk_golang v1.0.0」

在replace模块添加「doudian.com/open/sdk_golang => ./lib/doudian-sdk」

![img](https://zsmhub.github.io/images/2023/WX20230225-150227.png)

4. 执行go build命令编译，没有报错说明导入sdk包成功

### 获取Acccess Token代码示例

1. 自用型应用获取Access Token

	```go
	//设置appKey和appSecret，全局设置一次
	doudian_sdk.GlobalConfig.AppKey = "xxxxxxxx"
	doudian_sdk.GlobalConfig.AppSecret = "xxxxxxxxxx"
	accessToken, err := doudian_sdk.BuildAccessToken(&doudian_sdk.BuildAccessTokenParam{ShopId: 4463798})
	```

2. 工具型应用获取Access Token

	```go
	//设置appKey和appSecret，全局设置一次
	doudian_sdk.GlobalConfig.AppKey = "xxxxxxxx"
	doudian_sdk.GlobalConfig.AppSecret = "xxxxxxxxxx"
	accessToken, err := doudian_sdk.BuildAccessToken(&doudian_sdk.BuildAccessTokenParam{Code: "xxxxxxxx"})
	```

3. 应用自己管理Access Token

	```go
	//设置appKey和appSecret，全局设置一次
	doudian_sdk.GlobalConfig.AppKey = "xxxxxxxx"
	doudian_sdk.GlobalConfig.AppSecret = "xxxxxxxxxx"
	accessToken, err := doudian_sdk.ParseAccessToken("xxxxxxxxxx")
	```

### 调用Open Api示例

```go
//设置appKey和appSecret，全局设置一次
doudian_sdk.GlobalConfig.AppKey = "xxxxxxxx"
doudian_sdk.GlobalConfig.AppSecret = "xxxxxxxxxx"
//创建Access Token
accessToken, err := doudian_sdk.BuildAccessToken(&doudian_sdk.BuildAccessTokenParam{ShopId: 4463798})
if err != nil {
	panic(err)
}
//构建Open Api请求参数
request := address_list_request.New()
param := request.GetParams()
param.PageNo = 1
param.PageSize = 10
param.ShopId = 123456
//调用Open Api
response, err := request.Execute(accessToken)
```

### 调用消息推送示例

```go
import (
	doudian_message "doudian.com/open/sdk_golang/message"
	"encoding/json"
	"fmt"
)

func init() {
	doudian_message.GlobalConfig.AppKey = "xxx"
	doudian_message.GlobalConfig.AppSecret = "xxx"
}

func HandleDoudianPostRequestController(c echo.Context) error {
	if err := HandleDoudianPostRequestRepo(c); err != nil {
		return c.JSON(http.StatusOK, doudian_message.FailResponse)
	}
	return c.JSON(http.StatusOK, doudian_message.SuccessResponse)
}

func HandleDoudianPostRequestRepo(c echo.Context) error {
	msgs, err := doudian_message.GetMessage(c.Request())
	if err != nil {
		return err
	}
	for _, msg := range msgs {
		if msg.Error != nil {
			fmt.Println(msg.Error)
			continue
		}
		if msg.Message.Tag == new(doudian_message.TradeCreate).Tag() {
			msgData, ok := msg.Message.DataStruct.(doudian_message.TradeCreate)
			if !ok {
				fmt.Println("消息结构体解析异常")
				continue
			}
			fmt.Println("回调的父订单号：", msgData.PId)
		}
	}
}
```

### 抖店官方文档

- [抖店开放平台文档中心](https://op.jinritemai.com/docs/guide-docs/213/14)
- [SDK使用手册](https://op.jinritemai.com/docs/guide-docs/1041/1072)
- [消息推送服务接入指南](https://op.jinritemai.com/docs/guide-docs/215/1977)

### 推荐开源项目

- [企业微信 GO SDK](https://github.com/zsmhub/workweixin)
- [微信视频号 GO SDK](https://github.com/zsmhub/wx-channels-sdk)