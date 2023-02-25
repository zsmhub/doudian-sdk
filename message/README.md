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
	msgJson, _ := json.Marshal(msgs)
	fmt.Println(msgJson)
	return nil
}
```