/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 更新单个广告日限额条件
type AdgroupsUpdateDailyBudgetUpdateDailyBudgetStruct struct {
	AdgroupId   *int64 `json:"adgroup_id,omitempty"`
	DailyBudget *int64 `json:"daily_budget,omitempty"`
}
