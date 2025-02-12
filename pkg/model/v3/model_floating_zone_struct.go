/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 浮层卡片结构
type FloatingZoneStruct struct {
	FloatingZoneSwitch                *bool                    `json:"floating_zone_switch,omitempty"`
	FloatingZoneImageId               *string                  `json:"floating_zone_image_id,omitempty"`
	FloatingZoneName                  *string                  `json:"floating_zone_name,omitempty"`
	FloatingZoneDesc                  *string                  `json:"floating_zone_desc,omitempty"`
	FloatingZoneButtonText            *string                  `json:"floating_zone_button_text,omitempty"`
	FloatingZoneShowAppPropertySwitch *bool                    `json:"floating_zone_show_app_property_switch,omitempty"`
	FloatingZoneType                  CreativeFloatingZoneType `json:"floating_zone_type,omitempty"`
	FloatingZoneSingleImageId         *string                  `json:"floating_zone_single_image_id,omitempty"`
	ButtonBaseText                    *string                  `json:"button_base_text,omitempty"`
	JumpInfo                          *JumpinfoStruct          `json:"jump_info,omitempty"`
	FloatingZoneInfoType              FloatingZoneInfoType     `json:"floating_zone_info_type,omitempty"`
}
