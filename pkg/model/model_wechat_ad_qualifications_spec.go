/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信广告资质信息，当且仅当qualification_type=AD_QUALIFICATION_WECHAT时可填且必填
type WechatAdQualificationsSpec struct {
	QualificationName *string `json:"qualification_name,omitempty"`
	ImageId           *string `json:"image_id,omitempty"`
}
