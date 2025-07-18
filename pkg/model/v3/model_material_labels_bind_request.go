/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type MaterialLabelsBindRequest struct {
	AccountId        *int64           `json:"account_id,omitempty"`
	OrganizationId   *int64           `json:"organization_id,omitempty"`
	ImageIdList      *[]string        `json:"image_id_list,omitempty"`
	MediaIdList      *[]string        `json:"media_id_list,omitempty"`
	LabelIdList      *[]int64         `json:"label_id_list,omitempty"`
	BindingType      BindingType      `json:"binding_type,omitempty"`
	BusinessScenario BusinessScenario `json:"business_scenario,omitempty"`
}
