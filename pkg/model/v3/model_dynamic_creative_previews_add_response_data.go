/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type DynamicCreativePreviewsAddResponseData struct {
	AccountId          *int64    `json:"account_id,omitempty"`
	TraceId            *string   `json:"trace_id,omitempty"`
	WxFailUserNameList *[]string `json:"wx_fail_user_name_list,omitempty"`
}
