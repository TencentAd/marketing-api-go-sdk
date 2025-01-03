/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 渠道包信息
type WechatOfficialAccountStruct struct {
	WechatOfficialAccountId   *string `json:"wechat_official_account_id,omitempty"`
	WechatOfficialAccountName *string `json:"wechat_official_account_name,omitempty"`
	CreatedTime               *int64  `json:"created_time,omitempty"`
	LastModifiedTime          *int64  `json:"last_modified_time,omitempty"`
	WechatOfficialAccountIcon *string `json:"wechat_official_account_icon,omitempty"`
}
