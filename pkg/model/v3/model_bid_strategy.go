/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// BidStrategy : 出价策略
type BidStrategy string

// List of BidStrategy
const (
	BidStrategy_UNSUPPORTED       BidStrategy = "BID_STRATEGY_UNSUPPORTED"
	BidStrategy_AVERAGE_COST      BidStrategy = "BID_STRATEGY_AVERAGE_COST"
	BidStrategy_TARGET_COST       BidStrategy = "BID_STRATEGY_TARGET_COST"
	BidStrategy_PRIORITY_LOW_COST BidStrategy = "BID_STRATEGY_PRIORITY_LOW_COST"
	BidStrategy_PRIORITY_CAP_COST BidStrategy = "BID_STRATEGY_PRIORITY_CAP_COST"
)
