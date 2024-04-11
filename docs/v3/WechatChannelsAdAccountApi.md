# TencentAds\WechatChannelsAdAccountApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatChannelsAdAccountAdd**](WechatChannelsAdAccountApi.md#WechatChannelsAdAccountAdd) | **Post** /wechat_channels_ad_account/add | 视频号开户
[**WechatChannelsAdAccountDelete**](WechatChannelsAdAccountApi.md#WechatChannelsAdAccountDelete) | **Post** /wechat_channels_ad_account/delete | 删除开户记录
[**WechatChannelsAdAccountGet**](WechatChannelsAdAccountApi.md#WechatChannelsAdAccountGet) | **Get** /wechat_channels_ad_account/get | 查询视频号开户记录
[**WechatChannelsAdAccountUpdate**](WechatChannelsAdAccountApi.md#WechatChannelsAdAccountUpdate) | **Post** /wechat_channels_ad_account/update | 视频号开户修改


# **WechatChannelsAdAccountAdd**
> WechatChannelsAdAccountAddResponse WechatChannelsAdAccountAdd(ctx, data)
视频号开户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAdAccountAddRequest**](WechatChannelsAdAccountAddRequest.md)|  | 

### Return type

[**WechatChannelsAdAccountAddResponse**](WechatChannelsAdAccountAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAdAccountDelete**
> WechatChannelsAdAccountDeleteResponse WechatChannelsAdAccountDelete(ctx, data)
删除开户记录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAdAccountDeleteRequest**](WechatChannelsAdAccountDeleteRequest.md)|  | 

### Return type

[**WechatChannelsAdAccountDeleteResponse**](WechatChannelsAdAccountDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAdAccountGet**
> WechatChannelsAdAccountGetResponse WechatChannelsAdAccountGet(ctx, accountId, optional)
查询视频号开户记录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***WechatChannelsAdAccountApiWechatChannelsAdAccountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatChannelsAdAccountApiWechatChannelsAdAccountGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatChannelsAdAccountGetResponse**](WechatChannelsAdAccountGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WechatChannelsAdAccountUpdate**
> WechatChannelsAdAccountUpdateResponse WechatChannelsAdAccountUpdate(ctx, data)
视频号开户修改

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatChannelsAdAccountUpdateRequest**](WechatChannelsAdAccountUpdateRequest.md)|  | 

### Return type

[**WechatChannelsAdAccountUpdateResponse**](WechatChannelsAdAccountUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

