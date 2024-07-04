/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// IntOptimizationGoal : 优化目标类型，支持的优化目标及对应的深度优化目标可通过“<a target='_blank' href='/docs/api/tools/capabilities/optimization_goal_permissions_get' >查询优化目标权限</a>”接口进行查询
type IntOptimizationGoal string

// List of IntOptimizationGoal
const (
	IntOptimizationGoal_NONE                                   IntOptimizationGoal = "OPTIMIZATIONGOAL_NONE"
	IntOptimizationGoal_BRAND_CONVERSION                       IntOptimizationGoal = "OPTIMIZATIONGOAL_BRAND_CONVERSION"
	IntOptimizationGoal_FOLLOW                                 IntOptimizationGoal = "OPTIMIZATIONGOAL_FOLLOW"
	IntOptimizationGoal_CLICK                                  IntOptimizationGoal = "OPTIMIZATIONGOAL_CLICK"
	IntOptimizationGoal_IMPRESSION                             IntOptimizationGoal = "OPTIMIZATIONGOAL_IMPRESSION"
	IntOptimizationGoal_APP_DOWNLOAD                           IntOptimizationGoal = "OPTIMIZATIONGOAL_APP_DOWNLOAD"
	IntOptimizationGoal_APP_ACTIVATE                           IntOptimizationGoal = "OPTIMIZATIONGOAL_APP_ACTIVATE"
	IntOptimizationGoal_APP_REGISTER                           IntOptimizationGoal = "OPTIMIZATIONGOAL_APP_REGISTER"
	IntOptimizationGoal_ONE_DAY_RETENTION                      IntOptimizationGoal = "OPTIMIZATIONGOAL_ONE_DAY_RETENTION"
	IntOptimizationGoal_APP_PURCHASE                           IntOptimizationGoal = "OPTIMIZATIONGOAL_APP_PURCHASE"
	IntOptimizationGoal_ECOMMERCE_ORDER                        IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_ORDER"
	IntOptimizationGoal_ECOMMERCE_CHECKOUT                     IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_CHECKOUT"
	IntOptimizationGoal_LEADS                                  IntOptimizationGoal = "OPTIMIZATIONGOAL_LEADS"
	IntOptimizationGoal_ECOMMERCE_CART                         IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_CART"
	IntOptimizationGoal_PROMOTION_CLICK_KEY_PAGE               IntOptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_CLICK_KEY_PAGE"
	IntOptimizationGoal_VIEW_COMMODITY_PAGE                    IntOptimizationGoal = "OPTIMIZATIONGOAL_VIEW_COMMODITY_PAGE"
	IntOptimizationGoal_ONLINE_CONSULTATION                    IntOptimizationGoal = "OPTIMIZATIONGOAL_ONLINE_CONSULTATION"
	IntOptimizationGoal_TELEPHONE_CONSULTATION                 IntOptimizationGoal = "OPTIMIZATIONGOAL_TELEPHONE_CONSULTATION"
	IntOptimizationGoal_PAGE_RESERVATION                       IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_RESERVATION"
	IntOptimizationGoal_DELIVERY                               IntOptimizationGoal = "OPTIMIZATIONGOAL_DELIVERY"
	IntOptimizationGoal_MESSAGE_AFTER_FOLLOW                   IntOptimizationGoal = "OPTIMIZATIONGOAL_MESSAGE_AFTER_FOLLOW"
	IntOptimizationGoal_CLICK_MENU_AFTER_FOLLOW                IntOptimizationGoal = "OPTIMIZATIONGOAL_CLICK_MENU_AFTER_FOLLOW"
	IntOptimizationGoal_PAGE_EFFECTIVE_ONLINE_CONSULT          IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_EFFECTIVE_ONLINE_CONSULT"
	IntOptimizationGoal_CONFIRM_EFFECTIVE_LEADS_CONSULT        IntOptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_CONSULT"
	IntOptimizationGoal_CONFIRM_EFFECTIVE_LEADS_PHONE          IntOptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_PHONE"
	IntOptimizationGoal_LEADS_COLLECT                          IntOptimizationGoal = "OPTIMIZATIONGOAL_LEADS_COLLECT"
	IntOptimizationGoal_FIRST_PURCHASE                         IntOptimizationGoal = "OPTIMIZATIONGOAL_FIRST_PURCHASE"
	IntOptimizationGoal_APPLY                                  IntOptimizationGoal = "OPTIMIZATIONGOAL_APPLY"
	IntOptimizationGoal_PRE_CREDIT                             IntOptimizationGoal = "OPTIMIZATIONGOAL_PRE_CREDIT"
	IntOptimizationGoal_CREDIT                                 IntOptimizationGoal = "OPTIMIZATIONGOAL_CREDIT"
	IntOptimizationGoal_WITHDRAW_DEPOSITS                      IntOptimizationGoal = "OPTIMIZATIONGOAL_WITHDRAW_DEPOSITS"
	IntOptimizationGoal_PROMOTION_VIEW_KEY_PAGE                IntOptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_VIEW_KEY_PAGE"
	IntOptimizationGoal_MOBILE_APP_CREATE_ROLE                 IntOptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_CREATE_ROLE"
	IntOptimizationGoal_CANVAS_CLICK                           IntOptimizationGoal = "OPTIMIZATIONGOAL_CANVAS_CLICK"
	IntOptimizationGoal_PROMOTION_CLAIM_OFFER                  IntOptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_CLAIM_OFFER"
	IntOptimizationGoal_ECOMMERCE_ADD_TO_WISHLIST              IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_ADD_TO_WISHLIST"
	IntOptimizationGoal_CONFIRM_EFFECTIVE_LEADS_RESERVATION    IntOptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_EFFECTIVE_LEADS_RESERVATION"
	IntOptimizationGoal_PAGE_RECEIPT                           IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_RECEIPT"
	IntOptimizationGoal_PAGE_SCAN_CODE                         IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_SCAN_CODE"
	IntOptimizationGoal_SELECT_COURSE                          IntOptimizationGoal = "OPTIMIZATIONGOAL_SELECT_COURSE"
	IntOptimizationGoal_CONFIRM_POTENTIAL_CUSTOMER_PHONE       IntOptimizationGoal = "OPTIMIZATIONGOAL_CONFIRM_POTENTIAL_CUSTOMER_PHONE"
	IntOptimizationGoal_MOBILE_APP_AD_INCOME                   IntOptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_AD_INCOME"
	IntOptimizationGoal_MOBILE_APP_ACCREDIT                    IntOptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_ACCREDIT"
	IntOptimizationGoal_PURCHASE_MEMBER_CARD                   IntOptimizationGoal = "OPTIMIZATIONGOAL_PURCHASE_MEMBER_CARD"
	IntOptimizationGoal_PAGE_CONFIRM_EFFECTIVE_LEADS           IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_CONFIRM_EFFECTIVE_LEADS"
	IntOptimizationGoal_ADD_DESKTOP                            IntOptimizationGoal = "OPTIMIZATIONGOAL_ADD_DESKTOP"
	IntOptimizationGoal_RESERVATION                            IntOptimizationGoal = "OPTIMIZATIONGOAL_RESERVATION"
	IntOptimizationGoal_FIRST_ECOMMERCE_ORDER                  IntOptimizationGoal = "OPTIMIZATIONGOAL_FIRST_ECOMMERCE_ORDER"
	IntOptimizationGoal_FIRST_TWENTY_FOUR_HOUR_ECOMMERCE_ORDER IntOptimizationGoal = "OPTIMIZATIONGOAL_FIRST_TWENTY_FOUR_HOUR_ECOMMERCE_ORDER"
	IntOptimizationGoal_ECOMMERCE_SCANCODE_WX                  IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_SCANCODE_WX"
	IntOptimizationGoal_CLASS_PARTICIPATED                     IntOptimizationGoal = "OPTIMIZATIONGOAL_CLASS_PARTICIPATED"
	IntOptimizationGoal_INSURANCE_PURCHASE                     IntOptimizationGoal = "OPTIMIZATIONGOAL_INSURANCE_PURCHASE"
	IntOptimizationGoal_MOBILE_APP_SEVEN_DAYS_RETENTION        IntOptimizationGoal = "OPTIMIZATIONGOAL_MOBILE_APP_SEVEN_DAYS_RETENTION"
	IntOptimizationGoal_LIKE                                   IntOptimizationGoal = "OPTIMIZATIONGOAL_LIKE"
	IntOptimizationGoal_EXTERNAL_LINK_CLICK                    IntOptimizationGoal = "OPTIMIZATIONGOAL_EXTERNAL_LINK_CLICK"
	IntOptimizationGoal_BUY_COUPONS                            IntOptimizationGoal = "OPTIMIZATIONGOAL_BUY_COUPONS"
	IntOptimizationGoal_LEAVE_INFORMATION                      IntOptimizationGoal = "OPTIMIZATIONGOAL_LEAVE_INFORMATION"
	IntOptimizationGoal_CORE_ACTION                            IntOptimizationGoal = "OPTIMIZATIONGOAL_CORE_ACTION"
	IntOptimizationGoal_ONE_DAY_RETENTION_RATIO                IntOptimizationGoal = "OPTIMIZATIONGOAL_ONE_DAY_RETENTION_RATIO"
	IntOptimizationGoal_PROMOTION_READ_ARTICLE                 IntOptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_READ_ARTICLE"
	IntOptimizationGoal_RESERVATION_CHECK                      IntOptimizationGoal = "OPTIMIZATIONGOAL_RESERVATION_CHECK"
	IntOptimizationGoal_OPEN_ACCOUNT                           IntOptimizationGoal = "OPTIMIZATIONGOAL_OPEN_ACCOUNT"
	IntOptimizationGoal_SEVEN_DAY_ECOMMERCE_ORDER              IntOptimizationGoal = "OPTIMIZATIONGOAL_SEVEN_DAY_ECOMMERCE_ORDER"
	IntOptimizationGoal_ADD_WECHAT                             IntOptimizationGoal = "OPTIMIZATIONGOAL_ADD_WECHAT"
	IntOptimizationGoal_WECOM_CONSULT                          IntOptimizationGoal = "OPTIMIZATIONGOAL_WECOM_CONSULT"
	IntOptimizationGoal_ADD_GROUP                              IntOptimizationGoal = "OPTIMIZATIONGOAL_ADD_GROUP"
	IntOptimizationGoal_QUICK_ORDER                            IntOptimizationGoal = "OPTIMIZATIONGOAL_QUICK_ORDER"
	IntOptimizationGoal_PRE_PAY                                IntOptimizationGoal = "OPTIMIZATIONGOAL_PRE_PAY"
	IntOptimizationGoal_PAGE_ONLINE_CONSULT_ACTIVE_ONE_MSG     IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_ONLINE_CONSULT_ACTIVE_ONE_MSG"
	IntOptimizationGoal_CALL_DURATION_THIRTY_SECONDS           IntOptimizationGoal = "OPTIMIZATIONGOAL_CALL_DURATION_THIRTY_SECONDS"
	IntOptimizationGoal_CLAIM_COURSE                           IntOptimizationGoal = "OPTIMIZATIONGOAL_CLAIM_COURSE"
	IntOptimizationGoal_QUIT_GROUP                             IntOptimizationGoal = "OPTIMIZATIONGOAL_QUIT_GROUP"
	IntOptimizationGoal_VIEW_ACQUISITION_CONTENT               IntOptimizationGoal = "OPTIMIZATIONGOAL_VIEW_ACQUISITION_CONTENT"
	IntOptimizationGoal_BACK_FLOW                              IntOptimizationGoal = "OPTIMIZATIONGOAL_BACK_FLOW"
	IntOptimizationGoal_PAGE_ONLINE_CONSULT_THREE_MSG          IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_ONLINE_CONSULT_THREE_MSG"
	IntOptimizationGoal_RENEWAL                                IntOptimizationGoal = "OPTIMIZATIONGOAL_RENEWAL"
	IntOptimizationGoal_LOW_PRICE_COURSE                       IntOptimizationGoal = "OPTIMIZATIONGOAL_LOW_PRICE_COURSE"
	IntOptimizationGoal_CONSULT_INTENTION                      IntOptimizationGoal = "OPTIMIZATIONGOAL_CONSULT_INTENTION"
	IntOptimizationGoal_EVERY_DAY_RETENTION                    IntOptimizationGoal = "OPTIMIZATIONGOAL_EVERY_DAY_RETENTION"
	IntOptimizationGoal_PROMOTION_VIEW_KEY_PAGE_UV             IntOptimizationGoal = "OPTIMIZATIONGOAL_PROMOTION_VIEW_KEY_PAGE_UV"
	IntOptimizationGoal_LIVE_STREAM_DURATION_1_MIN             IntOptimizationGoal = "OPTIMIZATIONGOAL_LIVE_STREAM_DURATION_1MIN"
	IntOptimizationGoal_LIVE_STREAM_INTERACTION                IntOptimizationGoal = "OPTIMIZATIONGOAL_LIVE_STREAM_INTERACTION"
	IntOptimizationGoal_ECOMMERCE_CANCEL_ORDER                 IntOptimizationGoal = "OPTIMIZATIONGOAL_ECOMMERCE_CANCEL_ORDER"
	IntOptimizationGoal_CLICK_LEADS_COMPONENT                  IntOptimizationGoal = "OPTIMIZATIONGOAL_CLICK_LEADS_COMPONENT"
	IntOptimizationGoal_REGULAR_PRICE_COURSE                   IntOptimizationGoal = "OPTIMIZATIONGOAL_REGULAR_PRICE_COURSE"
	IntOptimizationGoal_VISIT_STROE                            IntOptimizationGoal = "OPTIMIZATIONGOAL_VISIT_STROE"
	IntOptimizationGoal_EFFECTIVE_ENTRY                        IntOptimizationGoal = "OPTIMIZATIONGOAL_EFFECTIVE_ENTRY"
	IntOptimizationGoal_CREDIT_RATIO                           IntOptimizationGoal = "OPTIMIZATIONGOAL_CREDIT_RATIO"
	IntOptimizationGoal_QYT_LIVE_STREAM_DEAL                   IntOptimizationGoal = "OPTIMIZATIONGOAL_QYT_LIVE_STREAM_DEAL"
	IntOptimizationGoal_QYT_LIVE_STREAM_PRODUCT_CLICK          IntOptimizationGoal = "OPTIMIZATIONGOAL_QYT_LIVE_STREAM_PRODUCT_CLICK"
	IntOptimizationGoal_QYT_LIVE_STREAM_AUDIENCE               IntOptimizationGoal = "OPTIMIZATIONGOAL_QYT_LIVE_STREAM_AUDIENCE"
	IntOptimizationGoal_QYT_LIVE_STREAM_COMMENT                IntOptimizationGoal = "OPTIMIZATIONGOAL_QYT_LIVE_STREAM_COMMENT"
	IntOptimizationGoal_QYT_LIVE_STREAM_FANS                   IntOptimizationGoal = "OPTIMIZATIONGOAL_QYT_LIVE_STREAM_FANS"
	IntOptimizationGoal_24HFIRSTPAY                            IntOptimizationGoal = "OPTIMIZATIONGOAL_24H_FIRSTPAY"
	IntOptimizationGoal_STORE_STAY                             IntOptimizationGoal = "OPTIMIZATIONGOAL_STORE_STAY"
	IntOptimizationGoal_R3                                     IntOptimizationGoal = "OPTIMIZATIONGOAL_R3"
	IntOptimizationGoal_LOW_PRICE_INSURANCE_PAYMENT            IntOptimizationGoal = "OPTIMIZATIONGOAL_LOW_PRICE_INSURANCE_PAYMENT"
	IntOptimizationGoal_UNDERWRITING                           IntOptimizationGoal = "OPTIMIZATIONGOAL_UNDERWRITING"
	IntOptimizationGoal_FIRST_WITHDRAW                         IntOptimizationGoal = "OPTIMIZATIONGOAL_FIRST_WITHDRAW"
	IntOptimizationGoal_PAGE_EFFECTIVE_PHONE_CALL              IntOptimizationGoal = "OPTIMIZATIONGOAL_PAGE_EFFECTIVE_PHONE_CALL"
)
