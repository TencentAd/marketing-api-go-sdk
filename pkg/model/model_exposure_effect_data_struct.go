/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 效果数据（曝光）
type ExposureEffectDataStruct struct {
	Count       *int64   `json:"count,omitempty"`
	CategoryWin *float64 `json:"category_win,omitempty"`
	CategoryAvg *float64 `json:"category_avg,omitempty"`
}
