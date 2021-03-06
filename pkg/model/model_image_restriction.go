/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 图片需满足的限制条件，仅当element_type是ELEMENT_TYPE_IMAGE_ARRAY或ELEMENT_TYPE_IMAGE时返回
type ImageRestriction struct {
	Width      *int64    `json:"width,omitempty"`
	Height     *int64    `json:"height,omitempty"`
	FileSize   *int64    `json:"file_size,omitempty"`
	FileFormat *[]string `json:"file_format,omitempty"`
}
