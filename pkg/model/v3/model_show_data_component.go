/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 数据外显组件
type ShowDataComponent struct {
	ComponentId *int64          `json:"component_id,omitempty"`
	Value       *ShowDataStruct `json:"value,omitempty"`
	IsDeleted   *bool           `json:"is_deleted,omitempty"`
}
