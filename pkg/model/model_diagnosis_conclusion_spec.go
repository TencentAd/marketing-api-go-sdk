/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// 整体曝光评估结果
type DiagnosisConclusionSpec struct {
	TimeRange        *DiagnosisConclusionSpecTimeRange `json:"time_range,omitempty"`
	DiagnosisScore   *int64                            `json:"diagnosis_score,omitempty"`
	SameIndustryRank *int64                            `json:"same_industry_rank,omitempty"`
}
