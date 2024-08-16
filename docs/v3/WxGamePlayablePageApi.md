# TencentAds\WxGamePlayablePageApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WxGamePlayablePageGet**](WxGamePlayablePageApi.md#WxGamePlayablePageGet) | **Get** /wx_game_playable_page/get | 获取微信小游戏试玩页


# **WxGamePlayablePageGet**
> WxGamePlayablePageGetResponse WxGamePlayablePageGet(ctx, accountId, appId, optional)
获取微信小游戏试玩页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **appId** | **string**|  | 
 **optional** | ***WxGamePlayablePageApiWxGamePlayablePageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WxGamePlayablePageApiWxGamePlayablePageGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WxGamePlayablePageGetResponse**](WxGamePlayablePageGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

