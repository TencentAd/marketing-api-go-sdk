/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdParamGetResponseData struct {
	SystemAdParamList  *[]AdParamListStruct `json:"system_ad_param_list,omitempty"`
	ProductAdParamList *[]AdParamListStruct `json:"product_ad_param_list,omitempty"`
}
