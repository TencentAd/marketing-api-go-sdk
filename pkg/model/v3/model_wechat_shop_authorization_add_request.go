/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type WechatShopAuthorizationAddRequest struct {
	AccountId              *int64  `json:"account_id,omitempty"`
	WechatChannelsShopName *string `json:"wechat_channels_shop_name,omitempty"`
	AuthorizationTtl       *int64  `json:"authorization_ttl,omitempty"`
}
