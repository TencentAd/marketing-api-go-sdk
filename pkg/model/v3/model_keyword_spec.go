/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// Keyword人群信息
type KeywordSpec struct {
	IncludeKeyword *[]string `json:"include_keyword,omitempty"`
	ExcludeKeyword *[]string `json:"exclude_keyword,omitempty"`
}