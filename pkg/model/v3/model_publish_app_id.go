/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 已发布版本落地页 App信息
type PublishAppId struct {
	AndroidAppId *int64 `json:"android_app_id,omitempty"`
	IosAppId     *int64 `json:"ios_app_id,omitempty"`
}