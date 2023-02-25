## 自主整理的抖店回调消息SDK，非官方

### 调用方法

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