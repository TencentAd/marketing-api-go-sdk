/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件字段结构
type ComponentMetadataValueField struct {
	Name      *string                                      `json:"name,omitempty"`
	Type_     ComponentMetadataFieldType                   `json:"type,omitempty"`
	Structure *[]ComponentMetadataValueFieldStructureField `json:"structure,omitempty"`
	IsArray   *bool                                        `json:"is_array,omitempty"`
	Valid     *ComponentMetadataValueValid                 `json:"valid,omitempty"`
}
