/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 应用分包信息
type SplitChannelPackageData struct {
	PackageId           *int64                `json:"package_id,omitempty"`
	ChannelName         *string               `json:"channel_name,omitempty"`
	ChannelPackageId    *string               `json:"channel_package_id,omitempty"`
	ChannelId           *string               `json:"channel_id,omitempty"`
	SystemStatus        UnionPackageSysStatus `json:"system_status,omitempty"`
	CreatedTime         *int64                `json:"created_time,omitempty"`
	LastModifiedTime    *int64                `json:"last_modified_time,omitempty"`
	CustomizedChannelId *string               `json:"customized_channel_id,omitempty"`
}