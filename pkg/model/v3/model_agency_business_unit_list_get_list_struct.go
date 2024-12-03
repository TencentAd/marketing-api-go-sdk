/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 服务商业务单元信息
type AgencyBusinessUnitListGetListStruct struct {
	OrganizationId    *int64   `json:"organization_id,omitempty"`
	BusinessUnitName  *string  `json:"business_unit_name,omitempty"`
	ManagerUserIdList *[]int64 `json:"manager_user_id_list,omitempty"`
}
