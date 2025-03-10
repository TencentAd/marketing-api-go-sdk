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
type WechatChannelsAdAccountGetListStruct struct {
	WechatChannelsAdAccountId *int64                                  `json:"wechat_channels_ad_account_id,omitempty"`
	ExportUsername            *string                                 `json:"export_username,omitempty"`
	WechatBindAuthToken       *string                                 `json:"wechat_bind_auth_token,omitempty"`
	WechatBindStatus          WechatChannelsAdAccountBindWechatStatus `json:"wechat_bind_status,omitempty"`
	Nickname                  *string                                 `json:"nickname,omitempty"`
	HeadImageId               *string                                 `json:"head_image_id,omitempty"`
	HeadImageUrl              *string                                 `json:"head_image_url,omitempty"`
	AuditMsg                  *string                                 `json:"audit_msg,omitempty"`
	Status                    WechatChannelsAdAccountStatus           `json:"status,omitempty"`
	CertificationList         *[]AdAccountCertification               `json:"certification_list,omitempty"`
	CreatedTime               *int64                                  `json:"created_time,omitempty"`
	LastModifiedTime          *int64                                  `json:"last_modified_time,omitempty"`
	FinderFailMsg             *string                                 `json:"finder_fail_msg,omitempty"`
	FinderFailRet             *int64                                  `json:"finder_fail_ret,omitempty"`
	WechatChannelsAccountId   *string                                 `json:"wechat_channels_account_id,omitempty"`
	LogoutTimeSecond          *int64                                  `json:"logout_time_second,omitempty"`
}
