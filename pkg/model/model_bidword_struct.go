/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 关键词信息
type BidwordStruct struct {
	CampaignId            *int64                    `json:"campaign_id,omitempty"`
	AdgroupId             *int64                    `json:"adgroup_id,omitempty"`
	Bidword               *string                   `json:"bidword,omitempty"`
	BidPrice              *int64                    `json:"bid_price,omitempty"`
	UseGroupPrice         UseGroupPriceType         `json:"use_group_price,omitempty"`
	MatchType             BidwordMatchType          `json:"match_type,omitempty"`
	ConfiguredStatus      BidwordPauseType          `json:"configured_status,omitempty"`
	PcLandingPageInfo     *[]KeywordLandingPageInfo `json:"pc_landing_page_info,omitempty"`
	MobileLandingPageInfo *[]KeywordLandingPageInfo `json:"mobile_landing_page_info,omitempty"`
	DynamicCreativeId     *int64                    `json:"dynamic_creative_id,omitempty"`
}
