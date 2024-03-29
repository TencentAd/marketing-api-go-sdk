/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 基础轮播图组件元素
type SliderSpec struct {
	ImageIdList *[]string `json:"image_id_list,omitempty"`
	Width       *int64    `json:"width,omitempty"`
	Height      *int64    `json:"height,omitempty"`
	SliderStyle *int64    `json:"slider_style,omitempty"`
	BgColor     *string   `json:"bg_color,omitempty"`
}
