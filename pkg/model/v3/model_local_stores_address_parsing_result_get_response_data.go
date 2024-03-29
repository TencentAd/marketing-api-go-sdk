/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type LocalStoresAddressParsingResultGetResponseData struct {
	SuccessList   *[]SuccessPoiInfoStruct `json:"success_list,omitempty"`
	FailList      *[]string               `json:"fail_list,omitempty"`
	SensitiveList *[]string               `json:"sensitive_list,omitempty"`
}
