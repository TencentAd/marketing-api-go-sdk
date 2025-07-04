/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 侧边悬浮按钮组件配置
type RightFixedButtonConfig struct {
	BackgroundColor *string                        `json:"background_color,omitempty"`
	Type_           *string                        `json:"type,omitempty"`
	Content         *[]CompRightFixedButtonContent `json:"content,omitempty"`
}
