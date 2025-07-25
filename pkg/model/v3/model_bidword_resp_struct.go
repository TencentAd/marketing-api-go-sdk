/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回的关键词信息
type BidwordRespStruct struct {
	Index             *int64           `json:"index,omitempty"`
	BidwordId         *int64           `json:"bidword_id,omitempty"`
	Bidword           *string          `json:"bidword,omitempty"`
	BidPrice          *int64           `json:"bid_price,omitempty"`
	MatchType         BidwordMatchType `json:"match_type,omitempty"`
	ConfiguredStatus  BidwordPauseType `json:"configured_status,omitempty"`
	ErrorMsg          *string          `json:"error_msg,omitempty"`
	ApprovalStatus    *int64           `json:"approval_status,omitempty"`
	DynamicCreativeId *int64           `json:"dynamic_creative_id,omitempty"`
}
