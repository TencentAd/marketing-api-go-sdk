/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config/v3"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model/v3"
)

type OperationLogListGetExample struct {
	TAds                    *ads.SDKClient
	AccessToken             string
	AccountId               int64
	OperationObjectType     string
	StartDate               string
	EndDate                 string
	Page                    int64
	PageSize                int64
	OperationLogListGetOpts *api.OperationLogListGetOpts
}

func (e *OperationLogListGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.AccountId = 789
	e.OperationObjectType = "operationObjectType_example"
	e.StartDate = "startDate_example"
	e.EndDate = "endDate_example"
	e.Page = 789
	e.PageSize = 789
	e.OperationLogListGetOpts = &api.OperationLogListGetOpts{}
}

func (e *OperationLogListGetExample) RunExample() (model.OperationLogListGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.OperationLogList().Get(ctx, e.AccountId, e.OperationObjectType, e.StartDate, e.EndDate, e.Page, e.PageSize, e.OperationLogListGetOpts)
}

func main() {
	e := &OperationLogListGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
