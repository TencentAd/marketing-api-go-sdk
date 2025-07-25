/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告绑定的rta策略
type UpdateAdgroupBindRtaPolicyItem struct {
	AdgroupId           *int64  `json:"adgroup_id,omitempty"`
	OriginRtaPolicyUuid *string `json:"origin_rta_policy_uuid,omitempty"`
	TargetRtaPolicyUuid *string `json:"target_rta_policy_uuid,omitempty"`
	AccountId           *int64  `json:"account_id,omitempty"`
}
