/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组一方跑量
type UpdateAdgroupEcomPkamItem struct {
	AdgroupId      *int64         `json:"adgroup_id,omitempty"`
	EcomPkamSwitch EcomPkamSwitch `json:"ecom_pkam_switch,omitempty"`
}
