# TencentAds\WechatShopAuthorizationValidationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatShopAuthorizationValidationGet**](WechatShopAuthorizationValidationApi.md#WechatShopAuthorizationValidationGet) | **Get** /wechat_shop_authorization_validation/get | 微信小店授权校验


# **WechatShopAuthorizationValidationGet**
> WechatShopAuthorizationValidationGetResponse WechatShopAuthorizationValidationGet(ctx, accountId, optional)
微信小店授权校验

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatShopAuthorizationValidationApiWechatShopAuthorizationValidationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatShopAuthorizationValidationApiWechatShopAuthorizationValidationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wechatChannelsShopId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatShopAuthorizationValidationGetResponse**](WechatShopAuthorizationValidationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

