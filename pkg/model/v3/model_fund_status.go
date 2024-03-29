/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// FundStatus : 资金状态
type FundStatus string

// List of FundStatus
const (
	FundStatus_NORMAL     FundStatus = "FUND_STATUS_NORMAL"
	FundStatus_NOT_ENOUGH FundStatus = "FUND_STATUS_NOT_ENOUGH"
	FundStatus_CLOSED     FundStatus = "FUND_STATUS_CLOSED"
	FundStatus_FROZEN     FundStatus = "FUND_STATUS_FROZEN"
)
