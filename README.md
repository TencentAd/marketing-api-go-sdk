# Marketing Go SDK
## 概述
腾讯广告 Marketing API(以下简称API) SDK 提供了Token获取、请求封装、响应解释等功能，以本地化方式轻松完成API的调用和结果的获取，旨在帮助开发者快速搭建投放管理系统。
未来还会基于常用的广告投放场景，提供场景化的接口组合及调用封装。

## 使用条件
1. 使用SDK需要首先注册成为腾讯广告开发者，请参考[开发者快速入门文档](https://developers.e.qq.com/docs/start)
2. 使用SDK需要先拥有API的访问权限，所有SDK的使用与应用拥有的权限组相关联
3. Go SDK 需要依赖 1.11 版本及以上

## 如何安装
推荐使用 go mod 的方式获取
### 使用 go mod 获取
其中的 ${VERSION} 需要替换成要使用的版本，如v1.0.0，具体版本号可查看相应的tag
```shell
go mod edit -require="github.com/tencentad/marketing-api-go-sdk@${VERSION}"
go mod download
```

### 使用 go get 获取
```shell
go get github.com/tencentad/marketing-api-go-sdk
```

## 如何使用
SDK数组参数调用的方法名与API接口一一对应，如campaigns/get接口就对应tads.Campaigns().Get()方法

注意:model中的所有基本数据类型均为指针类型, 例如：*string, *bool, *int64, *float64

### 获取Access Token
###### 注：本示例适用于授权时通过Authorization Code获取Access Token和Refresh Token，如需更新Access Token请参考 ./examples/Authentication/RefreshAccessToken.go 示例
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

func main() {
	tads := ads.Init(&config.SDKConfig{})
	// your client id
	clientId := int64(0)
	clientSecret := "your client secret"
	grantType := "authorization_code"
	oauthTokenOpts := &api.OauthTokenOpts{
		AuthorizationCode: optional.NewString("your authorization code"),
		RedirectUri: optional.NewString("your authorization code"),
	}
	ctx := *tads.Ctx
	// oauth/token接口即对应Oauth().Token()方法
	response, _, err := tads.Oauth().Token(ctx, clientId, clientSecret, grantType, oauthTokenOpts)

	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			// TODO for api error
			fmt.Println("Response error:", string(errStr))
		} else {
			// TODO for other error
			fmt.Println("Error:", err)
		}
	}
	tads.SetAccessToken(response.AccessToken)
}
```
### 设置调用环境、Access Token
```go
package main

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
)

func main() {
	accessToken := "YOUR ACCESS TOKEN"
	tads := ads.Init(&config.SDKConfig{
		AccessToken: accessToken,
	})
	// 默认访问沙箱环境，如访问正式环境，请调用tads.UseProduction()
	tads.UseSandbox()
}
```

### 调用API接口
```go
package main

import (
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

func main() {
	accessToken := "YOUR ACCESS TOKEN"
	tads := ads.Init(&config.SDKConfig{
		AccessToken: accessToken,
	})
	// your account id
	accountId := int64(0)
    field := "promoted_object_type"
    operator := "EQUALS"
	campaignsGetOpts := &api.CampaignsGetOpts{
		Filtering: optional.NewInterface([]model.FilteringStruct{model.FilteringStruct{
			Field:&field,
			Operator:&operator,
			Values:[]string{"PROMOTED_OBJECT_TYPE_APP_IOS"},
		}}),
	}
	ctx := *tads.Ctx
	// oauth/token接口即对应Campaigns().Get()方法
	response, _, err := tads.Campaigns().Get(ctx, accountId, campaignsGetOpts)

	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			// When Api returns an error
			errStr, _ := json.Marshal(resErr)
			// TODO for api error
			fmt.Println("Response error:", string(errStr))
		} else {
			// When validation fails or other local issues
			// TODO for other error
			fmt.Println("Error:", err)
		}
	}
	fmt.Println(response)
}
```

### 调试和查看API接口日志
```go
package main

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
)

func main() {
	accessToken := "YOUR ACCESS TOKEN"
	ads.Init(&config.SDKConfig{
		AccessToken: accessToken,
		IsDebug: true,
		DebugFile: "YOUR LOG FILE PATH",
	})
}
```

### 对于返回的Json取严格模式校验
```go
package main

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
)

func main() {
	accessToken := "YOUR ACCESS TOKEN"
	ads.Init(&config.SDKConfig{
		AccessToken: accessToken,
		IsStrictMode: true
	})
}
```
如果返回值中包含不认识的属性，会抛ResponseStrictError


### 关闭SDK上报
###### 目前SDK上报信息为您的服务器版本和 Go 版本信息，为了帮助您更好地定位使用上的问题，建议开启上报，如需关闭请参考如下配置。
```go
package main

import (
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
)

func main() {
	accessToken := "YOUR ACCESS TOKEN"
	ads.Init(&config.SDKConfig{
		AccessToken: accessToken,
		SkipMonitor: true,
	})
}
```

## 问题建议与反馈
如果您在使用SDK过程中有任何问题与建议，请随时登录[开发者官网](https://developers.e.qq.com/)，点击右下角的"咨询"按钮，与我们的客服支持人员联系

## 后续计划
1. 丰富各类场景示例
2. 推出其他语言的SDK
