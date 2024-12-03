/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 目标创意形式信息
type TargetCreativeTemplateDimension struct {
	RatioWidth          *int64   `json:"ratio_width,omitempty"`
	MinWidth            *int64   `json:"min_width,omitempty"`
	MinHeight           *int64   `json:"min_height,omitempty"`
	CreativeTemplateIds *[]int64 `json:"creative_template_ids,omitempty"`
}
