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
type PreviewItemStruct struct {
	MaterialPreviewId               *int64                           `json:"material_preview_id,omitempty"`
	TargetCreativeTemplateDimension *TargetCreativeTemplateDimension `json:"target_creative_template_dimension,omitempty"`
	TemplateId                      *int64                           `json:"template_id,omitempty"`
	Type_                           *string                          `json:"type,omitempty"`
	TemplatePreviewUrl              *string                          `json:"template_preview_url,omitempty"`
	RecommendTemplates              *[]int64                         `json:"recommend_templates,omitempty"`
}