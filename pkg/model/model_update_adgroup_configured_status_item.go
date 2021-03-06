/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组客户设置的状态
type UpdateAdgroupConfiguredStatusItem struct {
	AdgroupId        *int64   `json:"adgroup_id,omitempty"`
	ConfiguredStatus AdStatus `json:"configured_status,omitempty"`
}
