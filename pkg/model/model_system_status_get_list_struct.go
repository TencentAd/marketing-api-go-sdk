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
type SystemStatusGetListStruct struct {
	AdgroupId      *int64         `json:"adgroup_id,omitempty"`
	LearningStatus LearningStatus `json:"learning_status,omitempty"`
}
