# TencentAds\WechatPagesCsgrouplistApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatPagesCsgrouplistAdd**](WechatPagesCsgrouplistApi.md#WechatPagesCsgrouplistAdd) | **Post** /wechat_pages_csgrouplist/add | 增加企业微信组件客服组
[**WechatPagesCsgrouplistGet**](WechatPagesCsgrouplistApi.md#WechatPagesCsgrouplistGet) | **Get** /wechat_pages_csgrouplist/get | 获取企业微信客服组列表
[**WechatPagesCsgrouplistUpdate**](WechatPagesCsgrouplistApi.md#WechatPagesCsgrouplistUpdate) | **Post** /wechat_pages_csgrouplist/update | 更新企业微信组件客服组


# **WechatPagesCsgrouplistAdd**
> WechatPagesCsgrouplistAddResponse WechatPagesCsgrouplistAdd(ctx, data)
增加企业微信组件客服组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatPagesCsgrouplistAddRequest**](WechatPagesCsgrouplistAddRequest.md)|  | 

### Return type

[**WechatPagesCsgrouplistAddResponse**](WechatPagesCsgrouplistAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatPagesCsgrouplistGet**
> WechatPagesCsgrouplistGetResponse WechatPagesCsgrouplistGet(ctx, accountId, optional)
获取企业微信客服组列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatPagesCsgrouplistApiWechatPagesCsgrouplistGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatPagesCsgrouplistApiWechatPagesCsgrouplistGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **corpId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatPagesCsgrouplistGetResponse**](WechatPagesCsgrouplistGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatPagesCsgrouplistUpdate**
> WechatPagesCsgrouplistUpdateResponse WechatPagesCsgrouplistUpdate(ctx, data)
更新企业微信组件客服组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatPagesCsgrouplistUpdateRequest**](WechatPagesCsgrouplistUpdateRequest.md)|  | 

### Return type

[**WechatPagesCsgrouplistUpdateResponse**](WechatPagesCsgrouplistUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

