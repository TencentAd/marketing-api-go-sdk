/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// android 默认落地页内容
type AndroidAppPageSpec struct {
	AndroidAppId       *string `json:"android_app_id,omitempty"`
	WechatCanvasPageId *int64  `json:"wechat_canvas_page_id,omitempty"`
	AndroidChannelId   *string `json:"android_channel_id,omitempty"`
}
