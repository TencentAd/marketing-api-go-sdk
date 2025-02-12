/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProgrammedTemplateGetRequest struct {
	AccountId         *int64          `json:"account_id,omitempty"`
	MaterialDeriveId  *int64          `json:"material_derive_id,omitempty"`
	MaterialPreviewId *int64          `json:"material_preview_id,omitempty"`
	TemplateIdList    *[]int64        `json:"template_id_list,omitempty"`
	KeyWord           *string         `json:"key_word,omitempty"`
	SortBy            *[]SortByStruct `json:"sort_by,omitempty"`
	PageInfo          *PageInfoStruct `json:"page_info,omitempty"`
}
