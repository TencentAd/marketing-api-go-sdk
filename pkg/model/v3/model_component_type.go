/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package model

// ComponentType : 创意组件类型
type ComponentType string

// List of ComponentType
const (
	ComponentType_UNKNOWN                    ComponentType = "UNKNOWN"
	ComponentType_TITLE                      ComponentType = "TITLE"
	ComponentType_DESCRIPTION                ComponentType = "DESCRIPTION"
	ComponentType_IMAGE                      ComponentType = "IMAGE"
	ComponentType_IMAGE_LIST                 ComponentType = "IMAGE_LIST"
	ComponentType_JUMP_INFO                  ComponentType = "JUMP_INFO"
	ComponentType_VIDEO                      ComponentType = "VIDEO"
	ComponentType_BRAND                      ComponentType = "BRAND"
	ComponentType_CONSULT                    ComponentType = "CONSULT"
	ComponentType_PHONE                      ComponentType = "PHONE"
	ComponentType_FORM                       ComponentType = "FORM"
	ComponentType_ACTION_BUTTON              ComponentType = "ACTION_BUTTON"
	ComponentType_CHOSEN_BUTTON              ComponentType = "CHOSEN_BUTTON"
	ComponentType_LABEL                      ComponentType = "LABEL"
	ComponentType_SHOW_DATA                  ComponentType = "SHOW_DATA"
	ComponentType_MARKETING_PENDANT          ComponentType = "MARKETING_PENDANT"
	ComponentType_APP_GIFT_PACK_CODE         ComponentType = "APP_GIFT_PACK_CODE"
	ComponentType_SHOP_IMAGE                 ComponentType = "SHOP_IMAGE"
	ComponentType_COUNT_DOWN                 ComponentType = "COUNT_DOWN"
	ComponentType_BARRAGE                    ComponentType = "BARRAGE"
	ComponentType_FLOATING_ZONE              ComponentType = "FLOATING_ZONE"
	ComponentType_TEXT_LINK                  ComponentType = "TEXT_LINK"
	ComponentType_END_PAGE                   ComponentType = "END_PAGE"
	ComponentType_LIVING_DESC                ComponentType = "LIVING_DESC"
	ComponentType_WECHAT_CHANNELS            ComponentType = "WECHAT_CHANNELS"
	ComponentType_SHORT_VIDEO                ComponentType = "SHORT_VIDEO"
	ComponentType_ELEMENT_STORY              ComponentType = "ELEMENT_STORY"
	ComponentType_WXGAME_PLAYABLE_PAGE       ComponentType = "WXGAME_PLAYABLE_PAGE"
	ComponentType_APP_PROMOTION_VIDEO        ComponentType = "APP_PROMOTION_VIDEO"
	ComponentType_VIDEO_SHOWCASE             ComponentType = "VIDEO_SHOWCASE"
	ComponentType_IMAGE_SHOWCASE             ComponentType = "IMAGE_SHOWCASE"
	ComponentType_SOCIAL_SKILL               ComponentType = "SOCIAL_SKILL"
	ComponentType_MINI_CARD_LINK             ComponentType = "MINI_CARD_LINK"
	ComponentType_SUBLINK                    ComponentType = "SUBLINK"
	ComponentType_CONSULT_LINK               ComponentType = "CONSULT_LINK"
	ComponentType_SHOP_PRODUCT_CARD          ComponentType = "SHOP_PRODUCT_CARD"
	ComponentType_DYNAMIC_SHOWCASE           ComponentType = "DYNAMIC_SHOWCASE"
	ComponentType_UNVEIL_CARD                ComponentType = "UNVEIL_CARD"
	ComponentType_HOT_TOPIC                  ComponentType = "HOT_TOPIC"
	ComponentType_VIDEO_CHANNELS_CONTENT     ComponentType = "VIDEO_CHANNELS_CONTENT"
	ComponentType_V2_TITLE                   ComponentType = "V2_TITLE"
	ComponentType_V2_DESCRIPTION             ComponentType = "V2_DESCRIPTION"
	ComponentType_V2_LONG_SUBLINK            ComponentType = "V2_LONG_SUBLINK"
	ComponentType_V2_SHORT_SUBLINK           ComponentType = "V2_SHORT_SUBLINK"
	ComponentType_V2_LONG_SUBLINK_LIST       ComponentType = "V2_LONG_SUBLINK_LIST"
	ComponentType_V2_SHORT_SUBLINK_LIST      ComponentType = "V2_SHORT_SUBLINK_LIST"
	ComponentType_V2_APP_DOWNLOAD            ComponentType = "V2_APP_DOWNLOAD"
	ComponentType_V2_IMAGE_1_X1              ComponentType = "V2_IMAGE_1X1"
	ComponentType_V2_IMAGE_LIST_1_X1         ComponentType = "V2_IMAGE_LIST_1X1"
	ComponentType_V2_IMAGE_TEXT_1_X1         ComponentType = "V2_IMAGE_TEXT_1X1"
	ComponentType_V2_IMAGE_BIG_20_X7         ComponentType = "V2_IMAGE_BIG_20X7"
	ComponentType_V2_VIDEO_16_X9_IMAGE_16_X9 ComponentType = "V2_VIDEO_16X9_IMAGE_16X9"
	ComponentType_V2_VIDEO_16_X9_IMAGE_4_X3  ComponentType = "V2_VIDEO_16X9_IMAGE_4X3"
	ComponentType_V2_VIDEO_16_X9_IMAGE_1_X1  ComponentType = "V2_VIDEO_16X9_IMAGE_1X1"
	ComponentType_V2_QUICK_CONSULT           ComponentType = "V2_QUICK_CONSULT"
	ComponentType_V2_PHONE                   ComponentType = "V2_PHONE"
	ComponentType_V2_FORM                    ComponentType = "V2_FORM"
	ComponentType_V2_BRAND                   ComponentType = "V2_BRAND"
	ComponentType_V2_LANDING_PAGE            ComponentType = "V2_LANDING_PAGE"
	ComponentType_V2_HOLIDAY_LOGO            ComponentType = "V2_HOLIDAY_LOGO"
	ComponentType_V2_ACTION_BUTTON           ComponentType = "V2_ACTION_BUTTON"
	ComponentType_V2_CHOSEN_BUTTON           ComponentType = "V2_CHOSEN_BUTTON"
	ComponentType_V2_VIDEO_9_X16_IMAGE_9_X16 ComponentType = "V2_VIDEO_9X16_IMAGE_9X16"
	ComponentType_V2_IMAGE_16_X9             ComponentType = "V2_IMAGE_16X9"
	ComponentType_V2_LABEL                   ComponentType = "V2_LABEL"
	ComponentType_V2_PROMOTION_SUBLINK       ComponentType = "V2_PROMOTION_SUBLINK"
	ComponentType_V2_IMAGE_LIST_3_X2         ComponentType = "V2_IMAGE_LIST_3X2"
	ComponentType_V2_IMAGE_LIST_9_X16        ComponentType = "V2_IMAGE_LIST_9X16"
	ComponentType_V2_LIST_SUBLINK            ComponentType = "V2_LIST_SUBLINK"
	ComponentType_V2_MDPA_TITLE              ComponentType = "V2_MDPA_TITLE"
	ComponentType_V2_MDPA_DESCRIPTION        ComponentType = "V2_MDPA_DESCRIPTION"
	ComponentType_SEARCH_ALGORITHM_GEN       ComponentType = "SEARCH_ALGORITHM_GEN"
	ComponentType_SEARCH_DERIVATIVE_TITLE    ComponentType = "SEARCH_DERIVATIVE_TITLE"
	ComponentType_SMART_DELIVERY_AIGC        ComponentType = "SMART_DELIVERY_AIGC"
	ComponentType_MARKETING_ASSET_PLAYLET    ComponentType = "MARKETING_ASSET_PLAYLET"
)
