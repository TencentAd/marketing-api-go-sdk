/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 返回结构
type WechatStoreProductItemsGetListStruct struct {
	ProductOuterId                       *string           `json:"product_outer_id,omitempty"`
	ProductName                          *string           `json:"product_name,omitempty"`
	Description                          *string           `json:"description,omitempty"`
	ProductPrimaryImageUrls              *[]string         `json:"product_primary_image_urls,omitempty"`
	ProductDetailImageUrls               *[]string         `json:"product_detail_image_urls,omitempty"`
	Price                                *float64          `json:"price,omitempty"`
	FirstCategoryId                      *int64            `json:"first_category_id,omitempty"`
	FirstCategoryName                    *string           `json:"first_category_name,omitempty"`
	SecondCategoryId                     *int64            `json:"second_category_id,omitempty"`
	SecondCategoryName                   *string           `json:"second_category_name,omitempty"`
	ThirdCategoryId                      *int64            `json:"third_category_id,omitempty"`
	ThirdCategoryName                    *string           `json:"third_category_name,omitempty"`
	FourthCategoryId                     *int64            `json:"fourth_category_id,omitempty"`
	FourthCategoryName                   *string           `json:"fourth_category_name,omitempty"`
	BrandName                            *string           `json:"brand_name,omitempty"`
	SalesVolume                          *int64            `json:"sales_volume,omitempty"`
	CommentNum                           *int64            `json:"comment_num,omitempty"`
	SpuLinkStatus                        SpuLinkStatus     `json:"spu_link_status,omitempty"`
	LinkSpuId                            *string           `json:"link_spu_id,omitempty"`
	LinkSpuName                          *string           `json:"link_spu_name,omitempty"`
	WechatStoreProductStatus             *int64            `json:"wechat_store_product_status,omitempty"`
	WechatStoreProductAdStatus           *int64            `json:"wechat_store_product_ad_status,omitempty"`
	WechatStoreProductAdUnavailableCause *[]string         `json:"wechat_store_product_ad_unavailable_cause,omitempty"`
	StorePriceItemList                   *[]StorePriceItem `json:"store_price_item_list,omitempty"`
	StoreId                              *string           `json:"store_id,omitempty"`
	WxStoreProductOperateType            *int64            `json:"wx_store_product_operate_type,omitempty"`
	SkuStockStatus                       *int64            `json:"sku_stock_status,omitempty"`
}
