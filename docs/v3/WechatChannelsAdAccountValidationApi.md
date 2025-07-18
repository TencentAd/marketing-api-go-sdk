# TencentAds\WechatChannelsAdAccountValidationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatChannelsAdAccountValidationGet**](WechatChannelsAdAccountValidationApi.md#WechatChannelsAdAccountValidationGet) | **Get** /wechat_channels_ad_account_validation/get | 视频号一键开户校验


# **WechatChannelsAdAccountValidationGet**
> WechatChannelsAdAccountValidationGetResponse WechatChannelsAdAccountValidationGet(ctx, accountId, optional)
视频号一键开户校验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatChannelsAdAccountValidationApiWechatChannelsAdAccountValidationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatChannelsAdAccountValidationApiWechatChannelsAdAccountValidationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nickname** | **optional.String**|  | 
 **headImageId** | **optional.String**|  | 
 **wechatChannelsAccountId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatChannelsAdAccountValidationGetResponse**](WechatChannelsAdAccountValidationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

