/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 匹配规则
type UrlMatcher struct {
	ParamValue *string  `json:"param_value,omitempty"`
	Operator   Operator `json:"operator,omitempty"`
}
