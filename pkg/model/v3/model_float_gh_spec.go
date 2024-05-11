/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 关注公众号组件元素
type FloatGhSpec struct {
	FastFollow      *int64  `json:"fast_follow,omitempty"`
	BtnTitle        *string `json:"btn_title,omitempty"`
	FontColor       *string `json:"font_color,omitempty"`
	BtnBgColorTheme *string `json:"btn_bg_color_theme,omitempty"`
	BtnFontType     *int64  `json:"btn_font_type,omitempty"`
}