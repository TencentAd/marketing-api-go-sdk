/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 浮层卡片组件
type FloatingZoneComponent struct {
	ComponentId *int64              `json:"component_id,omitempty"`
	Value       *FloatingZoneStruct `json:"value,omitempty"`
	IsDeleted   *bool               `json:"is_deleted,omitempty"`
}
