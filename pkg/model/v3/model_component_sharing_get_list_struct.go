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
type ComponentSharingGetListStruct struct {
	SharedAccountId   *int64            `json:"shared_account_id,omitempty"`
	SharedAccountType SharedAccountType `json:"shared_account_type,omitempty"`
}