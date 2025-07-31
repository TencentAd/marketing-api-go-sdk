# TencentAds\WechatShopApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatShopGet**](WechatShopApi.md#WechatShopGet) | **Get** /wechat_shop/get | 查询微信小店信息


# **WechatShopGet**
> WechatShopGetResponse WechatShopGet(ctx, accountId, optional)
查询微信小店信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatShopApiWechatShopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatShopApiWechatShopGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wechatChannelsShopName** | **optional.String**|  | 
 **wechatChannelsShopId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatShopGetResponse**](WechatShopGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

