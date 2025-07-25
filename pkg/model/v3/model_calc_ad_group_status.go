/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// CalcAdGroupStatus : 广告在系统中的状态
type CalcAdGroupStatus string

// List of CalcAdGroupStatus
const (
	CalcAdGroupStatus_ADGROUP_STATUS_FROZEN                                           CalcAdGroupStatus = "ADGROUP_STATUS_FROZEN"
	CalcAdGroupStatus_ADGROUP_STATUS_SUSPEND                                          CalcAdGroupStatus = "ADGROUP_STATUS_SUSPEND"
	CalcAdGroupStatus_ADGROUP_STATUS_NOT_IN_DELIVERY_TIME                             CalcAdGroupStatus = "ADGROUP_STATUS_NOT_IN_DELIVERY_TIME"
	CalcAdGroupStatus_ADGROUP_STATUS_ACTIVE                                           CalcAdGroupStatus = "ADGROUP_STATUS_ACTIVE"
	CalcAdGroupStatus_ADGROUP_STATUS_DELETED                                          CalcAdGroupStatus = "ADGROUP_STATUS_DELETED"
	CalcAdGroupStatus_ADGROUP_STATUS_ACCOUNT_BALANCE_NOT_ENOUGH                       CalcAdGroupStatus = "ADGROUP_STATUS_ACCOUNT_BALANCE_NOT_ENOUGH"
	CalcAdGroupStatus_ADGROUP_STATUS_DAILY_BUDGET_REACHED                             CalcAdGroupStatus = "ADGROUP_STATUS_DAILY_BUDGET_REACHED"
	CalcAdGroupStatus_ADGROUP_STATUS_PARTIAL_ACTIVE                                   CalcAdGroupStatus = "ADGROUP_STATUS_PARTIAL_ACTIVE"
	CalcAdGroupStatus_ADGROUP_STATUS_WECHAT_CHANNELS_STOP                             CalcAdGroupStatus = "ADGROUP_STATUS_WECHAT_CHANNELS_STOP"
	CalcAdGroupStatus_ADGROUP_STATUS_CREATIVE_STATUS_PENDING                          CalcAdGroupStatus = "ADGROUP_STATUS_CREATIVE_STATUS_PENDING"
	CalcAdGroupStatus_ADGROUP_STATUS_CREATIVE_EMPTY                                   CalcAdGroupStatus = "ADGROUP_STATUS_CREATIVE_EMPTY"
	CalcAdGroupStatus_ADGROUP_STATUS_JOINT_BUDGET_REACHED                             CalcAdGroupStatus = "ADGROUP_STATUS_JOINT_BUDGET_REACHED"
	CalcAdGroupStatus_ADGROUP_STATUS_TOTAL_BUDGET_REACHED                             CalcAdGroupStatus = "ADGROUP_STATUS_TOTAL_BUDGET_REACHED"
	CalcAdGroupStatus_ADGROUP_STATUS_PRE_LOCK                                         CalcAdGroupStatus = "ADGROUP_STATUS_PRE_LOCK"
	CalcAdGroupStatus_ADGROUP_STATUS_UNLOCKING                                        CalcAdGroupStatus = "ADGROUP_STATUS_UNLOCKING"
	CalcAdGroupStatus_ADGROUP_STATUS_STOP                                             CalcAdGroupStatus = "ADGROUP_STATUS_STOP"
	CalcAdGroupStatus_ADGROUP_STATUS_LIVE_NOT_ACTIVE                                  CalcAdGroupStatus = "ADGROUP_STATUS_LIVE_NOT_ACTIVE"
	CalcAdGroupStatus_ADGROUP_STATUS_NOT_ACTIVE_PRODUCT_AUDIT_FAIL                    CalcAdGroupStatus = "ADGROUP_STATUS_NOT_ACTIVE_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_ADGROUP_STATUS_LIVE_VIOLATION                                   CalcAdGroupStatus = "ADGROUP_STATUS_LIVE_VIOLATION"
	CalcAdGroupStatus_ADGROUP_STATUS_WECHAT_STORE_PENDING                             CalcAdGroupStatus = "ADGROUP_STATUS_WECHAT_STORE_PENDING"
	CalcAdGroupStatus_ADGROUP_STATUS_WECHAT_STORE_CLOSING                             CalcAdGroupStatus = "ADGROUP_STATUS_WECHAT_STORE_CLOSING"
	CalcAdGroupStatus_ADGROUP_STATUS_WECHAT_STORE_CLOSED                              CalcAdGroupStatus = "ADGROUP_STATUS_WECHAT_STORE_CLOSED"
	CalcAdGroupStatus_ADGROUP_STATUS_NOT_ACTIVE_PRODUCT_REMOVED                       CalcAdGroupStatus = "ADGROUP_STATUS_NOT_ACTIVE_PRODUCT_REMOVED"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_DELETED                                    CalcAdGroupStatus = "SMART_ADGROUP_STATUS_DELETED"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_SUSPEND                                    CalcAdGroupStatus = "SMART_ADGROUP_STATUS_SUSPEND"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_JOINT_BUDGET_REACHED                       CalcAdGroupStatus = "SMART_ADGROUP_STATUS_JOINT_BUDGET_REACHED"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_ACCOUNT_BALANCE_NOT_ENOUGH                 CalcAdGroupStatus = "SMART_ADGROUP_STATUS_ACCOUNT_BALANCE_NOT_ENOUGH"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_DAILY_BUDGET_REACHED                       CalcAdGroupStatus = "SMART_ADGROUP_STATUS_DAILY_BUDGET_REACHED"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_ADGROUP_STATUS_STOP                        CalcAdGroupStatus = "SMART_ADGROUP_STATUS_ADGROUP_STATUS_STOP"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_NOT_IN_DELIVERY_TIME                       CalcAdGroupStatus = "SMART_ADGROUP_STATUS_NOT_IN_DELIVERY_TIME"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_ACTIVE                                     CalcAdGroupStatus = "SMART_ADGROUP_STATUS_ACTIVE"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_USING                                      CalcAdGroupStatus = "SMART_ADGROUP_STATUS_USING"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_LIVE_NOT_ACTIVE                    CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_LIVE_NOT_ACTIVE"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_LIVE_VIOLATION                     CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_LIVE_VIOLATION"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_PRODUCT_AUDIT_FAIL                 CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_CHANNELS_STOP                      CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_CHANNELS_STOP"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_LIVE_NOT_ACTIVE_PRODUCT_AUDIT_FAIL CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_LIVE_NOT_ACTIVE_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_LIVE_VIOLATION_PRODUCT_AUDIT_FAIL  CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_LIVE_VIOLATION_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_INVALID_CHANNELS_STOP_PRODUCT_AUDIT_FAIL   CalcAdGroupStatus = "SMART_ADGROUP_STATUS_INVALID_CHANNELS_STOP_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_USING_LIVE_NOT_ACTIVE                      CalcAdGroupStatus = "SMART_ADGROUP_STATUS_USING_LIVE_NOT_ACTIVE"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_USING_LIVE_VIOLATION                       CalcAdGroupStatus = "SMART_ADGROUP_STATUS_USING_LIVE_VIOLATION"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_USING_PRODUCT_AUDIT_FAIL                   CalcAdGroupStatus = "SMART_ADGROUP_STATUS_USING_PRODUCT_AUDIT_FAIL"
	CalcAdGroupStatus_SMART_ADGROUP_STATUS_USING_CHANNELS_STOP                        CalcAdGroupStatus = "SMART_ADGROUP_STATUS_USING_CHANNELS_STOP"
)
