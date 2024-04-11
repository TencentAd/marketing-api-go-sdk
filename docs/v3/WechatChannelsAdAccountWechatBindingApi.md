# TencentAds\WechatChannelsAdAccountWechatBindingApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatChannelsAdAccountWechatBindingAdd**](WechatChannelsAdAccountWechatBindingApi.md#WechatChannelsAdAccountWechatBindingAdd) | **Post** /wechat_channels_ad_account_wechat_binding/add | 视频号开户绑定微信号
[**WechatChannelsAdAccountWechatBindingGet**](WechatChannelsAdAccountWechatBindingApi.md#WechatChannelsAdAccountWechatBindingGet) | **Get** /wechat_channels_ad_account_wechat_binding/get | 视频号开户绑定微信状态查询


# **WechatChannelsAdAccountWechatBindingAdd**
> WechatChannelsAdAccountWechatBindingAddResponse WechatChannelsAdAccountWechatBindingAdd(ctx, data)
视频号开户绑定微信号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAdAccountWechatBindingAddRequest**](WechatChannelsAdAccountWechatBindingAddRequest.md)|  | 

### Return type

[**WechatChannelsAdAccountWechatBindingAddResponse**](WechatChannelsAdAccountWechatBindingAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAdAccountWechatBindingGet**
> WechatChannelsAdAccountWechatBindingGetResponse WechatChannelsAdAccountWechatBindingGet(ctx, accountId, wechatBindAuthToken, optional)
视频号开户绑定微信状态查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **wechatBindAuthToken** | **string**|  | 
 **optional** | ***WechatChannelsAdAccountWechatBindingApiWechatChannelsAdAccountWechatBindingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatChannelsAdAccountWechatBindingApiWechatChannelsAdAccountWechatBindingGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatChannelsAdAccountWechatBindingGetResponse**](WechatChannelsAdAccountWechatBindingGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

