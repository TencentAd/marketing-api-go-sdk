/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 基础视频组件元素
type VideoSpec struct {
	VideoId  *string `json:"video_id,omitempty"`
	Width    *int64  `json:"width,omitempty"`
	Height   *int64  `json:"height,omitempty"`
	InMiddle *int64  `json:"in_middle,omitempty"`
}
