/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 修改广告组安卓应用渠道包 id
type UpdateAdgroupAppAndroidChannelPackageIdItem struct {
	AdgroupId                  *int64  `json:"adgroup_id,omitempty"`
	AppAndroidChannelPackageId *string `json:"app_android_channel_package_id,omitempty"`
}
