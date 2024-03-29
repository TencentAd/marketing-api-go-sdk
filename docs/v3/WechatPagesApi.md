# TencentAds\WechatPagesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatPagesAdd**](WechatPagesApi.md#WechatPagesAdd) | **Post** /wechat_pages/add | 基于模板创建微信原生页
[**WechatPagesDelete**](WechatPagesApi.md#WechatPagesDelete) | **Post** /wechat_pages/delete | 删除微信落地页
[**WechatPagesGet**](WechatPagesApi.md#WechatPagesGet) | **Get** /wechat_pages/get | 获取微信落地页列表


# **WechatPagesAdd**
> WechatPagesAddResponse WechatPagesAdd(ctx, data)
基于模板创建微信原生页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatPagesAddRequest**](WechatPagesAddRequest.md)|  | 

### Return type

[**WechatPagesAddResponse**](WechatPagesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatPagesDelete**
> WechatPagesDeleteResponse WechatPagesDelete(ctx, data)
删除微信落地页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatPagesDeleteRequest**](WechatPagesDeleteRequest.md)|  | 

### Return type

[**WechatPagesDeleteResponse**](WechatPagesDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatPagesGet**
> WechatPagesGetResponse WechatPagesGet(ctx, accountId, optional)
获取微信落地页列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatPagesApiWechatPagesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatPagesApiWechatPagesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ownerUid** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatPagesGetResponse**](WechatPagesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

