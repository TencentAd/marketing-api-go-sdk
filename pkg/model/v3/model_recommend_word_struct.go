/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 门店推荐文案
type RecommendWordStruct struct {
	RecommendWord *string   `json:"recommend_word,omitempty"`
	Status        SysStatus `json:"status,omitempty"`
	AuditMsg      *string   `json:"audit_msg,omitempty"`
}
