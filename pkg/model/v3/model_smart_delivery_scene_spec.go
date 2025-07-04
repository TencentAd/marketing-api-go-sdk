/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 场景化投放信息
type SmartDeliverySceneSpec struct {
	SmartDeliveryGoal       SmartDeliveryGoal        `json:"smart_delivery_goal,omitempty"`
	SmartDeliveryGoalSpec   *SmartDeliveryGoalSpec   `json:"smart_delivery_goal_spec,omitempty"`
	ConversionIdList        *[]int64                 `json:"conversion_id_list,omitempty"`
	MarketingAssetOuterSpec *MarketingAssetOuterSpec `json:"marketing_asset_outer_spec,omitempty"`
	MarketingCarrierType    MarketingCarrierType     `json:"marketing_carrier_type,omitempty"`
	MarketingCarrierDetail  *MarketingCarrierDetail  `json:"marketing_carrier_detail,omitempty"`
}
