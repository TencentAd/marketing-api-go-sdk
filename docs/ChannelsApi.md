# TencentAds\ChannelsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsAdd**](ChannelsApi.md#ChannelsAdd) | **Post** /channels/add | 创建广告
[**ChannelsDelete**](ChannelsApi.md#ChannelsDelete) | **Post** /channels/delete | 删除广告
[**ChannelsGet**](ChannelsApi.md#ChannelsGet) | **Post** /channels/get | 拉取广告列表
[**ChannelsUpdate**](ChannelsApi.md#ChannelsUpdate) | **Post** /channels/update | 更新广告


# **ChannelsAdd**
> ChannelsAddResponse ChannelsAdd(ctx, data)
创建广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsAddRequest**](ChannelsAddRequest.md)|  | 

### Return type

[**ChannelsAddResponse**](ChannelsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsDelete**
> ChannelsDeleteResponse ChannelsDelete(ctx, data)
删除广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsDeleteRequest**](ChannelsDeleteRequest.md)|  | 

### Return type

[**ChannelsDeleteResponse**](ChannelsDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsGet**
> ChannelsGetResponse ChannelsGet(ctx, data)
拉取广告列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsGetRequest**](ChannelsGetRequest.md)|  | 

### Return type

[**ChannelsGetResponse**](ChannelsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChannelsUpdate**
> ChannelsUpdateResponse ChannelsUpdate(ctx, data)
更新广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ChannelsUpdateRequest**](ChannelsUpdateRequest.md)|  | 

### Return type

[**ChannelsUpdateResponse**](ChannelsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

