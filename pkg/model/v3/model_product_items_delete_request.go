/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type ProductItemsDeleteRequest struct {
	AccountId      *int64  `json:"account_id,omitempty"`
	CatalogId      *int64  `json:"catalog_id,omitempty"`
	ProductOuterId *string `json:"product_outer_id,omitempty"`
}