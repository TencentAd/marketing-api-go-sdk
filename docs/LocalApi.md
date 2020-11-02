# TencentAds\LocalApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalAdd**](LocalApi.md#LocalAdd) | **Post** /local/add | 创建广告
[**LocalDelete**](LocalApi.md#LocalDelete) | **Post** /local/delete | 删除广告
[**LocalGet**](LocalApi.md#LocalGet) | **Post** /local/get | 拉取广告列表
[**LocalUpdate**](LocalApi.md#LocalUpdate) | **Post** /local/update | 更新广告


# **LocalAdd**
> LocalAddResponse LocalAdd(ctx, data)
创建广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalAddRequest**](LocalAddRequest.md)|  | 

### Return type

[**LocalAddResponse**](LocalAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalDelete**
> LocalDeleteResponse LocalDelete(ctx, data)
删除广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalDeleteRequest**](LocalDeleteRequest.md)|  | 

### Return type

[**LocalDeleteResponse**](LocalDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalGet**
> LocalGetResponse LocalGet(ctx, data)
拉取广告列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalGetRequest**](LocalGetRequest.md)|  | 

### Return type

[**LocalGetResponse**](LocalGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LocalUpdate**
> LocalUpdateResponse LocalUpdate(ctx, data)
更新广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalUpdateRequest**](LocalUpdateRequest.md)|  | 

### Return type

[**LocalUpdateResponse**](LocalUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

