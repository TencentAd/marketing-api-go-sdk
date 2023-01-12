/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 场景分发状态
type SceneStatusRead struct {
	Scene           FileAvailableScene      `json:"scene,omitempty"`
	DispathchStatus DnFileDispatchStatus    `json:"dispathch_status,omitempty"`
	AttachInfo      *DispatchAttachInfoRead `json:"attach_info,omitempty"`
}