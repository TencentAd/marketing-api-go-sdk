/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type UnionPositionPackagesGetResponseData struct {
	List     *[]UnionPositionPackagesGetListStruct `json:"list,omitempty"`
	PageInfo *PageInfo                             `json:"page_info,omitempty"`
}
