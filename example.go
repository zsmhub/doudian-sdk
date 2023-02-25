package example

import (
	"doudian.com/open/sdk_golang/core"
	demo_spi_request "doudian.com/open/sdk_golang/spi/demo_spi/request"
	demo_spi_response "doudian.com/open/sdk_golang/spi/demo_spi/response"
	"fmt"
)

func ApiExample() {
	//doudian_sdk.GlobalConfig.AppKey = "xxxxx"
	//doudian_sdk.GlobalConfig.AppSecret = "xxxxx"
	//
	//var shopId string = "xx"
	//var subId string = "xxx"
	//
	//accessToken, err := doudian_sdk.BuildAccessToken(&doudian_sdk.BuildAccessTokenParam{AuthId: &shopId, AuthSubjectType: &subId})
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Printf("%+v", accessToken)

}

func SpiExample() {

	//query是服务端调用http请求时的请求参数，假如spi实现地址是http://www.xxx.yyy/abc
	//服务端实际调用的时候会拼接上请求参数：http://www.xxx.yyy/abc?app_key=7043675397937989128&param_json={}&timestamp=11112223333&sign=xxxxxxx&sign_method=md5&sign_v2=xxxxx
	//所以query等于app_key=7043675397937989128&param_json={}&timestamp=11112223333&sign=xxxxxxx&sign_method=md5&sign_v2=xxxxx

	query := "app_key=7043675397937989128&param_json={}&timestamp=11112223333&sign=xxxxxxx&sign_method=md5&sign_v2=xxxxx"
	resp, err := doudian_sdk.ConfigSpiWithUrlQuery(demo_spi_request.New(), BizHandler, query, nil).ResponseJson()
	if err != nil {
		panic(err)
	}
	//将resp作为http请求的body返回给服务器
	fmt.Println(resp)
}

// BizHandler 定义BizHandler，BizHandler用来处理接收到服务器请求之后，
//需要处理的业务逻辑，SDK会将服务器下发的业务参数（param_json）
//封装成对应的Param类实例，并放入DoudianOpSpiContext中，开发
//者可以通过context获取封装好的数据。SDK也会将需要返回的数据封
//装成Data类实例，开发者可以通过context.getData()获取该实例，
//并设置其中的字段，最终返回给开放平台服务器。
//
//在本示例中（demo/spi），Param对象为 DemoSpiParam，Data对象
//为 DemoSpiData
func BizHandler(context *doudian_sdk.DoudianOpSpiContext) {
	paramJson := context.GetParamObject().(*demo_spi_request.DemoSpiParam)
	fmt.Printf("%+v\n", paramJson)
	data := context.GetData().(*demo_spi_response.DemoSpiData)
	data.Data1 = "xxx"
	context.WrapSuccess()
}
