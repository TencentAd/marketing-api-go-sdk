/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 关注视频号信息元素
type VideoChannelSpec struct {
	Title          *string `json:"title,omitempty"`
	FinderNickname *string `json:"finder_nickname,omitempty"`
	FastFollow     *int64  `json:"fast_follow,omitempty"`
}
