/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type MaterialAuditListListStruct struct {
	AccountId    *int64  `json:"account_id,omitempty"`
	MaterialId   *string `json:"material_id,omitempty"`
	MaterialType *int64  `json:"material_type,omitempty"`
}
