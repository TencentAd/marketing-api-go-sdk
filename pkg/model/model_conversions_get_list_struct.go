/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type ConversionsGetListStruct struct {
	ConversionId                     *int64                     `json:"conversion_id,omitempty"`
	ConversionName                   *string                    `json:"conversion_name,omitempty"`
	AccessType                       AccessType                 `json:"access_type,omitempty"`
	ClaimType                        ClaimType                  `json:"claim_type,omitempty"`
	FeedbackUrl                      *string                    `json:"feedback_url,omitempty"`
	SelfAttributed                   *bool                      `json:"self_attributed,omitempty"`
	OptimizationGoal                 IntOptimizationGoal        `json:"optimization_goal,omitempty"`
	DeepBehaviorOptimizationGoal     IntOptimizationGoal        `json:"deep_behavior_optimization_goal,omitempty"`
	DeepWorthOptimizationGoal        ConversionOptimizationGoal `json:"deep_worth_optimization_goal,omitempty"`
	UserActionSetId                  *int64                     `json:"user_action_set_id,omitempty"`
	UserActionSetKey                 *string                    `json:"user_action_set_key,omitempty"`
	SiteSetEnable                    *bool                      `json:"site_set_enable,omitempty"`
	IsDeleted                        *bool                      `json:"is_deleted,omitempty"`
	AccessStatus                     AccessStatus               `json:"access_status,omitempty"`
	CreateSourceType                 CreateSourceType           `json:"create_source_type,omitempty"`
	AppAndroidChannelPackageId       *string                    `json:"app_android_channel_package_id,omitempty"`
	PromotedObjectId                 *string                    `json:"promoted_object_id,omitempty"`
	ConversionScene                  ConversionScene            `json:"conversion_scene,omitempty"`
	OwnerId                          *int64                     `json:"owner_id,omitempty"`
	DeepWorthAdvancedGoal            ConversionOptimizationGoal `json:"deep_worth_advanced_goal,omitempty"`
	ConversionLinkId                 *int64                     `json:"conversion_link_id,omitempty"`
	ImpressionFeedbackUrl            *string                    `json:"impression_feedback_url,omitempty"`
	AttributionWindow                *int64                     `json:"attribution_window,omitempty"`
	DeepBehaviorAdvancedGoal         IntOptimizationGoal        `json:"deep_behavior_advanced_goal,omitempty"`
	DeepBehaviorAdvancedGoalMinPrice *int64                     `json:"deep_behavior_advanced_goal_min_price,omitempty"`
	DeepBehaviorAdvancedGoalMaxPrice *int64                     `json:"deep_behavior_advanced_goal_max_price,omitempty"`
	DeepOptimizationGoalType         DeepOptimizationGoalType   `json:"deep_optimization_goal_type,omitempty"`
	ForwardLinkAssist                IntOptimizationGoal        `json:"forward_link_assist,omitempty"`
	IncubationOptimizationGoal       IntOptimizationGoal        `json:"incubation_optimization_goal,omitempty"`
	DisableMessage                   *string                    `json:"disable_message,omitempty"`
	InspectionFreeSwitch             *bool                      `json:"inspection_free_switch,omitempty"`
}
