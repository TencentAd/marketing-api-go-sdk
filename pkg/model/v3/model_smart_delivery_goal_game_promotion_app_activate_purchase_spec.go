/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 智投投放目标激活和付费
type SmartDeliveryGoalGamePromotionAppActivatePurchaseSpec struct {
	AppActivateCost *int64 `json:"app_activate_cost,omitempty"`
	AppPurchaseCost *int64 `json:"app_purchase_cost,omitempty"`
}
