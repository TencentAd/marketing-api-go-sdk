/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ChannelsGetResponseData struct {
	List     *[]AllDataSpec `json:"list,omitempty"`
	PageInfo *DpPageInfo    `json:"page_info,omitempty"`
}
