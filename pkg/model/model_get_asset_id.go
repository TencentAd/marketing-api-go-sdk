/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 资产详情
type GetAssetId struct {
	DataSetId       *int64 `json:"data_set_id,omitempty"`
	EnvType         *int64 `json:"env_type,omitempty"`
	DispathchStatus *int64 `json:"dispathch_status,omitempty"`
}