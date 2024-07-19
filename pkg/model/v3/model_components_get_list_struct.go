/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ComponentsGetListStruct struct {
	AccountId           *int64           `json:"account_id,omitempty"`
	ComponentId         *int64           `json:"component_id,omitempty"`
	ComponentValue      *ComponentValue  `json:"component_value,omitempty"`
	CreatedTime         *int64           `json:"created_time,omitempty"`
	LastModifiedTime    *int64           `json:"last_modified_time,omitempty"`
	ComponentSubType    ComponentSubType `json:"component_sub_type,omitempty"`
	ComponentCustomName *string          `json:"component_custom_name,omitempty"`
	IsDeleted           *bool            `json:"is_deleted,omitempty"`
}
