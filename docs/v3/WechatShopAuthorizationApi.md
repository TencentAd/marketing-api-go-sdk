# TencentAds\WechatShopAuthorizationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatShopAuthorizationAdd**](WechatShopAuthorizationApi.md#WechatShopAuthorizationAdd) | **Post** /wechat_shop_authorization/add | 创建微信小店授权
[**WechatShopAuthorizationGet**](WechatShopAuthorizationApi.md#WechatShopAuthorizationGet) | **Get** /wechat_shop_authorization/get | 获取微信小店授权记录列表


# **WechatShopAuthorizationAdd**
> WechatShopAuthorizationAddResponse WechatShopAuthorizationAdd(ctx, data)
创建微信小店授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatShopAuthorizationAddRequest**](WechatShopAuthorizationAddRequest.md)|  | 

### Return type

[**WechatShopAuthorizationAddResponse**](WechatShopAuthorizationAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatShopAuthorizationGet**
> WechatShopAuthorizationGetResponse WechatShopAuthorizationGet(ctx, accountId, optional)
获取微信小店授权记录列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatShopAuthorizationApiWechatShopAuthorizationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatShopAuthorizationApiWechatShopAuthorizationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatShopAuthorizationGetResponse**](WechatShopAuthorizationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

