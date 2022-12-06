/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 创意组件详情
type CreativeComponentSpecStruct struct {
	LiveImageComponent *LiveImageComponentSpecStruct `json:"live_image_component,omitempty"`
	LiveConvComponent  *LiveConvComponentSpecStruct  `json:"live_conv_component,omitempty"`
}