/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

//
type Tag struct {
	TagId       *int64  `json:"tag_id,omitempty"`
	ParentTagId *int64  `json:"parent_tag_id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	TagCode     *string `json:"tag_code,omitempty"`
	UserCount   *int64  `json:"user_count,omitempty"`
	CreatedTime *string `json:"created_time,omitempty"`
}
