/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 活动促销组件结构
type PromotionSublinkComponentValueStruct struct {
	Title       *string                 `json:"title,omitempty"`
	Description *string                 `json:"description,omitempty"`
	ButtonText  *string                 `json:"button_text,omitempty"`
	JumpInfo    *[]LandingPageStructure `json:"jump_info,omitempty"`
}