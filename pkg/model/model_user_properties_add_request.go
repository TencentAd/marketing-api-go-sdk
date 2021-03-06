/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type UserPropertiesAddRequest struct {
	AccountId         *int64                       `json:"account_id,omitempty"`
	UserPropertySetId *int64                       `json:"user_property_set_id,omitempty"`
	WechatAppId       *string                      `json:"wechat_app_id,omitempty"`
	Property          *[]UserPropertiesAddProperty `json:"property,omitempty"`
}
