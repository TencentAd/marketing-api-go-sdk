/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 出价模拟结果
type BidSimulateResultItem struct {
	Bid           *int64 `json:"bid,omitempty"`
	ExposeCnt     *int64 `json:"expose_cnt,omitempty"`
	ClickCnt      *int64 `json:"click_cnt,omitempty"`
	ConversionCnt *int64 `json:"conversion_cnt,omitempty"`
}
