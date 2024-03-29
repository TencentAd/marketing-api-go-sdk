# TencentAds\XijingPageApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingPageAdd**](XijingPageApi.md#XijingPageAdd) | **Post** /xijing_page/add | 蹊径-基于模板创建落地页
[**XijingPageDelete**](XijingPageApi.md#XijingPageDelete) | **Post** /xijing_page/delete | 蹊径-删除落地页
[**XijingPageUpdate**](XijingPageApi.md#XijingPageUpdate) | **Post** /xijing_page/update | 蹊径-送审落地页


# **XijingPageAdd**
> XijingPageAddResponse XijingPageAdd(ctx, data)
蹊径-基于模板创建落地页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**XijingPageAddRequest**](XijingPageAddRequest.md)|  | 

### Return type

[**XijingPageAddResponse**](XijingPageAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **XijingPageDelete**
> XijingPageDeleteResponse XijingPageDelete(ctx, data)
蹊径-删除落地页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**XijingPageDeleteRequest**](XijingPageDeleteRequest.md)|  | 

### Return type

[**XijingPageDeleteResponse**](XijingPageDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **XijingPageUpdate**
> XijingPageUpdateResponse XijingPageUpdate(ctx, data)
蹊径-送审落地页

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**XijingPageUpdateRequest**](XijingPageUpdateRequest.md)|  | 

### Return type

[**XijingPageUpdateResponse**](XijingPageUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

