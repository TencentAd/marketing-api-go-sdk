/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 阶段出价建议
type StageBidInfoItem struct {
	StageName  *string  `json:"stage_name,omitempty"`
	LowerBound *int64   `json:"lower_bound,omitempty"`
	UpperBound *int64   `json:"upper_bound,omitempty"`
	DefaultBid *int64   `json:"default_bid,omitempty"`
	Ratio      *float64 `json:"ratio,omitempty"`
}
