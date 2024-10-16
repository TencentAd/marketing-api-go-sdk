/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 直播间转化组件详情
type LiveConvComponentSpecStruct struct {
	PageId     *int64          `json:"page_id,omitempty"`
	PageType   DestinationType `json:"page_type,omitempty"`
	Title      *string         `json:"title,omitempty"`
	ImgUrl     *string         `json:"img_url,omitempty"`
	ImgId      *string         `json:"img_id,omitempty"`
	Desc       *string         `json:"desc,omitempty"`
	ButtonText *string         `json:"button_text,omitempty"`
}
