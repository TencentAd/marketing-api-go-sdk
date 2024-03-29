/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// EcommerceOrderStatus : 订单状态
type EcommerceOrderStatus string

// List of EcommerceOrderStatus
const (
	EcommerceOrderStatus_AWAITING_ORDER EcommerceOrderStatus = "AWAITING_ORDER"
	EcommerceOrderStatus_SHIPPING_SOON  EcommerceOrderStatus = "SHIPPING_SOON"
	EcommerceOrderStatus_SHIPPED        EcommerceOrderStatus = "SHIPPED"
	EcommerceOrderStatus_DELIVERED      EcommerceOrderStatus = "DELIVERED"
	EcommerceOrderStatus_RETURNED       EcommerceOrderStatus = "RETURNED"
)
