# TencentAds\WechatChannelsAuthorizationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatChannelsAuthorizationAdd**](WechatChannelsAuthorizationApi.md#WechatChannelsAuthorizationAdd) | **Post** /wechat_channels_authorization/add | 创建视频号授权
[**WechatChannelsAuthorizationDelete**](WechatChannelsAuthorizationApi.md#WechatChannelsAuthorizationDelete) | **Post** /wechat_channels_authorization/delete | 删除视频号授权
[**WechatChannelsAuthorizationGet**](WechatChannelsAuthorizationApi.md#WechatChannelsAuthorizationGet) | **Get** /wechat_channels_authorization/get | 获取视频号授权记录列表
[**WechatChannelsAuthorizationUpdate**](WechatChannelsAuthorizationApi.md#WechatChannelsAuthorizationUpdate) | **Post** /wechat_channels_authorization/update | 更新视频号授权


# **WechatChannelsAuthorizationAdd**
> WechatChannelsAuthorizationAddResponse WechatChannelsAuthorizationAdd(ctx, data)
创建视频号授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAuthorizationAddRequest**](WechatChannelsAuthorizationAddRequest.md)|  | 

### Return type

[**WechatChannelsAuthorizationAddResponse**](WechatChannelsAuthorizationAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAuthorizationDelete**
> WechatChannelsAuthorizationDeleteResponse WechatChannelsAuthorizationDelete(ctx, data)
删除视频号授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAuthorizationDeleteRequest**](WechatChannelsAuthorizationDeleteRequest.md)|  | 

### Return type

[**WechatChannelsAuthorizationDeleteResponse**](WechatChannelsAuthorizationDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAuthorizationGet**
> WechatChannelsAuthorizationGetResponse WechatChannelsAuthorizationGet(ctx, accountId, optional)
获取视频号授权记录列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatChannelsAuthorizationApiWechatChannelsAuthorizationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatChannelsAuthorizationApiWechatChannelsAuthorizationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wechatChannelsAccountName** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatChannelsAuthorizationGetResponse**](WechatChannelsAuthorizationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAuthorizationUpdate**
> WechatChannelsAuthorizationUpdateResponse WechatChannelsAuthorizationUpdate(ctx, data)
更新视频号授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAuthorizationUpdateRequest**](WechatChannelsAuthorizationUpdateRequest.md)|  | 

### Return type

[**WechatChannelsAuthorizationUpdateResponse**](WechatChannelsAuthorizationUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

