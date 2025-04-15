# TencentAds\WechatStoreCatalogsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatStoreCatalogsGet**](WechatStoreCatalogsApi.md#WechatStoreCatalogsGet) | **Get** /wechat_store_catalogs/get | 获取微信小店商品库


# **WechatStoreCatalogsGet**
> WechatStoreCatalogsGetResponse WechatStoreCatalogsGet(ctx, accountId, optional)
获取微信小店商品库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatStoreCatalogsApiWechatStoreCatalogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatStoreCatalogsApiWechatStoreCatalogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeIds** | [**optional.Interface of []string**](string.md)|  | 
 **catalogIds** | [**optional.Interface of []int64**](int64.md)|  | 
 **catalogName** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatStoreCatalogsGetResponse**](WechatStoreCatalogsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

