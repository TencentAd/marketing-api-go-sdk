/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// oCPX 深度辅助配置
type DeepConversionBehaviorAdvancedSpec struct {
	Goal      OptimizationGoal `json:"goal,omitempty"`
	BidAmount *int64           `json:"bid_amount,omitempty"`
}