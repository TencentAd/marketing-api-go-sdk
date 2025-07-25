/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 单图组件
type ImageComponent struct {
	ComponentId *int64       `json:"component_id,omitempty"`
	Value       *ImageStruct `json:"value,omitempty"`
	IsDeleted   *bool        `json:"is_deleted,omitempty"`
}
