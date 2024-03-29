/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 微信场景定向
type WechatScene struct {
	OfficialAccountMediaCategory *[]int64 `json:"official_account_media_category,omitempty"`
	MiniProgramAndMiniGame       *[]int64 `json:"mini_program_and_mini_game,omitempty"`
	PayScene                     *[]int64 `json:"pay_scene,omitempty"`
}
