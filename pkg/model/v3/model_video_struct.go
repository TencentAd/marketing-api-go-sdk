/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 视频组件
type VideoStruct struct {
	VideoId  *string         `json:"video_id,omitempty"`
	CoverId  *string         `json:"cover_id,omitempty"`
	JumpInfo *JumpinfoStruct `json:"jump_info,omitempty"`
}