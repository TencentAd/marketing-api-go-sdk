/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 含有特殊字符导致失败的否定词列表
type HasSpecialNegativeWordStruct struct {
	PhraseNegativeWords *[]string `json:"phrase_negative_words,omitempty"`
	ExactNegativeWords  *[]string `json:"exact_negative_words,omitempty"`
}
