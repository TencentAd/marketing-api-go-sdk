/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 小店商品落地页详情
type ChannelsShopProductDetail struct {
	ProductId   *int64  `json:"product_id,omitempty"`
	ShopId      *string `json:"shop_id,omitempty"`
	ProductName *string `json:"product_name,omitempty"`
	ShopName    *string `json:"shop_name,omitempty"`
}
