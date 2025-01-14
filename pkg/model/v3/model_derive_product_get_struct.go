/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 素材信息
type DeriveProductGetStruct struct {
	DeriveComponentId *int64                `json:"derive_component_id,omitempty"`
	OriginComponentId *int64                `json:"origin_component_id,omitempty"`
	TemplateId        *int64                `json:"template_id,omitempty"`
	DeriveMaterial    *MaterialItemStruct   `json:"derive_material,omitempty"`
	OriginMaterial    *[]MaterialItemStruct `json:"origin_material,omitempty"`
}