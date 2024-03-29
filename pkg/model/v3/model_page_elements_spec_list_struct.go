/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 页面组件列表
type PageElementsSpecListStruct struct {
	ElementType      CanvasPageElementType      `json:"element_type,omitempty"`
	TopImageSpec     *TopImageSpec              `json:"top_image_spec,omitempty"`
	TopSliderSpec    *TopSliderSpec             `json:"top_slider_spec,omitempty"`
	TopVideoSpec     *TopVideoSpec              `json:"top_video_spec,omitempty"`
	ImageSpec        *ImageSpec                 `json:"image_spec,omitempty"`
	SliderSpec       *SliderSpec                `json:"slider_spec,omitempty"`
	VideoSpec        *VideoSpec                 `json:"video_spec,omitempty"`
	TextSpec         *TextSpec                  `json:"text_spec,omitempty"`
	AppDownloadSpec  *CanvasAppDownloadSpecType `json:"app_download_spec,omitempty"`
	WeappSpec        *WeappSpec                 `json:"weapp_spec,omitempty"`
	GhSpec           *GhSpec                    `json:"gh_spec,omitempty"`
	EnterpriseWxSpec *EnterpriseWxSpec          `json:"enterprise_wx_spec,omitempty"`
	ImageTextSpec    *ImageTextSpec             `json:"image_text_spec,omitempty"`
}
