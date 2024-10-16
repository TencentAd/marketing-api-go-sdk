/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type AdgroupsUpdateRequest struct {
	AccountId                         *int64                      `json:"account_id,omitempty"`
	AdgroupId                         *int64                      `json:"adgroup_id,omitempty"`
	AdgroupName                       *string                     `json:"adgroup_name,omitempty"`
	BeginDate                         *string                     `json:"begin_date,omitempty"`
	EndDate                           *string                     `json:"end_date,omitempty"`
	FirstDayBeginTime                 *string                     `json:"first_day_begin_time,omitempty"`
	BidAmount                         *int64                      `json:"bid_amount,omitempty"`
	OptimizationGoal                  OptimizationGoal            `json:"optimization_goal,omitempty"`
	TimeSeries                        *string                     `json:"time_series,omitempty"`
	DailyBudget                       *int64                      `json:"daily_budget,omitempty"`
	Targeting                         *WriteTargetingSetting      `json:"targeting,omitempty"`
	SceneSpec                         *SceneTargetingForWrite     `json:"scene_spec,omitempty"`
	UserActionSets                    *[]UserActionSetStruct      `json:"user_action_sets,omitempty"`
	DeepConversionSpec                *DeepConversionSpec         `json:"deep_conversion_spec,omitempty"`
	ConversionId                      *int64                      `json:"conversion_id,omitempty"`
	DeepConversionBehaviorBid         *int64                      `json:"deep_conversion_behavior_bid,omitempty"`
	DeepConversionWorthRate           *float64                    `json:"deep_conversion_worth_rate,omitempty"`
	DeepConversionWorthAdvancedRate   *float64                    `json:"deep_conversion_worth_advanced_rate,omitempty"`
	DeepConversionBehaviorAdvancedBid *int64                      `json:"deep_conversion_behavior_advanced_bid,omitempty"`
	BidMode                           BidMode                     `json:"bid_mode,omitempty"`
	AutoAcquisitionEnabled            *bool                       `json:"auto_acquisition_enabled,omitempty"`
	AutoAcquisitionBudget             *int64                      `json:"auto_acquisition_budget,omitempty"`
	AutoDerivedLandingPageSwitch      *bool                       `json:"auto_derived_landing_page_switch,omitempty"`
	AutoDerivedCreativeEnabled        *bool                       `json:"auto_derived_creative_enabled,omitempty"`
	ConfiguredStatus                  ConfiguredStatus            `json:"configured_status,omitempty"`
	FlowOptimizationEnabled           *bool                       `json:"flow_optimization_enabled,omitempty"`
	PoiList                           *[]string                   `json:"poi_list,omitempty"`
	EcomPkamSwitch                    EcomPkamSwitch              `json:"ecom_pkam_switch,omitempty"`
	RtaId                             *int64                      `json:"rta_id,omitempty"`
	RtaTargetId                       *string                     `json:"rta_target_id,omitempty"`
	CostConstraintScene               CostConstraintScene         `json:"cost_constraint_scene,omitempty"`
	CustomCostCap                     *int64                      `json:"custom_cost_cap,omitempty"`
	FeedbackId                        *int64                      `json:"feedback_id,omitempty"`
	SearchExpandTargetingSwitch       SearchExpandTargetingSwitch `json:"search_expand_targeting_switch,omitempty"`
	CloudUnionSpec                    *CloudUnionSpec             `json:"cloud_union_spec,omitempty"`
}
