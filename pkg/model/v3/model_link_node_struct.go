/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 链路节点结构信息
type LinkNodeStruct struct {
	ConversionLinkNodeId     *int64    `json:"conversion_link_node_id,omitempty"`
	ConversionLinkNodeName   *string   `json:"conversion_link_node_name,omitempty"`
	ConversionLinkNodeIndex  *int64    `json:"conversion_link_node_index,omitempty"`
	ConversionLinkActionType *[]string `json:"conversion_link_action_type,omitempty"`
	OgId                     *[]int64  `json:"og_id,omitempty"`
	RoiOgId                  *[]int64  `json:"roi_og_id,omitempty"`
	CarrierId                *int64    `json:"carrier_id,omitempty"`
	CarrierName              *string   `json:"carrier_name,omitempty"`
	DataSource               *string   `json:"data_source,omitempty"`
}
