/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatChannelsAdAccountWechatBindingAddResponseData struct {
	WechatBindAuthToken        *string `json:"wechat_bind_auth_token,omitempty"`
	WechatBindAuthTokenExpTime *int64  `json:"wechat_bind_auth_token_exp_time,omitempty"`
	WechatBindScanUrl          *string `json:"wechat_bind_scan_url,omitempty"`
}