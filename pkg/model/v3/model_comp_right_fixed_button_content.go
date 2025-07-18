/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 侧边悬浮按钮组件内容
type CompRightFixedButtonContent struct {
	Title                 *string                           `json:"title,omitempty"`
	TitleColor            *string                           `json:"title_color,omitempty"`
	ButtonContent         *string                           `json:"button_content,omitempty"`
	ButtonFontColor       *string                           `json:"button_font_color,omitempty"`
	ButtonBackgroundColor *string                           `json:"button_background_color,omitempty"`
	IconColor             *string                           `json:"icon_color,omitempty"`
	ConvertUnknown        *ConvertInfoConfigUnknown         `json:"convert_unknown,omitempty"`
	ConvertDownload       *ConvertInfoConfigDownload        `json:"convert_download,omitempty"`
	ConvertLink           *ConvertInfoConfigLink            `json:"convert_link,omitempty"`
	ConvertWeapp          *ConvertInfoConfigMiniProgram     `json:"convert_weapp,omitempty"`
	ConvertGh             *ConvertInfoConfigOfficialAccount `json:"convert_gh,omitempty"`
	ConvertFollowVideo    *ConvertInfoConfigFollowVideo     `json:"convert_followVideo,omitempty"`
}
