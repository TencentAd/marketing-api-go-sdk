/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

type OfficialLandingPageDetailGetResponseData struct {
	PageId        *int64            `json:"page_id,omitempty"`
	LandingPageId *int64            `json:"landing_page_id,omitempty"`
	PageConfig    *PageConfigDetail `json:"page_config,omitempty"`
	PageElements  *[]PageElement    `json:"page_elements,omitempty"`
}
