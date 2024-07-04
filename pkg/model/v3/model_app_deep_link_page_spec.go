/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用直达数据
type AppDeepLinkPageSpec struct {
	AndroidDeepLinkAppId          *string `json:"android_deep_link_app_id,omitempty"`
	AndroidDeepLinkUrl            *string `json:"android_deep_link_url,omitempty"`
	IosDeepLinkAppId              *string `json:"ios_deep_link_app_id,omitempty"`
	IosDeepLinkUrl                *string `json:"ios_deep_link_url,omitempty"`
	UniversalLinkUrl              *string `json:"universal_link_url,omitempty"`
	MpaAndroidDeepLinkWildcardUrl *string `json:"mpa_android_deep_link_wildcard_url,omitempty"`
	MpaIosDeepLinkWildcardUrl     *string `json:"mpa_ios_deep_link_wildcard_url,omitempty"`
	MpaUniversalLinkWildcardUrl   *string `json:"mpa_universal_link_wildcard_url,omitempty"`
	DeepLinkUrl                   *string `json:"deep_link_url,omitempty"`
}
