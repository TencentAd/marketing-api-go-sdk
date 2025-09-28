# TencentAds\SearchDynamicCreativesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchDynamicCreativesAdd**](SearchDynamicCreativesApi.md#SearchDynamicCreativesAdd) | **Post** /search_dynamic_creatives/add | 创建搜索动态创意
[**SearchDynamicCreativesUpdate**](SearchDynamicCreativesApi.md#SearchDynamicCreativesUpdate) | **Post** /search_dynamic_creatives/update | 更新搜索创意


# **SearchDynamicCreativesAdd**
> SearchDynamicCreativesAddResponse SearchDynamicCreativesAdd(ctx, data)
创建搜索动态创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SearchDynamicCreativesAddRequest**](SearchDynamicCreativesAddRequest.md)|  | 

### Return type

[**SearchDynamicCreativesAddResponse**](SearchDynamicCreativesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchDynamicCreativesUpdate**
> SearchDynamicCreativesUpdateResponse SearchDynamicCreativesUpdate(ctx, data)
更新搜索创意

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SearchDynamicCreativesUpdateRequest**](SearchDynamicCreativesUpdateRequest.md)|  | 

### Return type

[**SearchDynamicCreativesUpdateResponse**](SearchDynamicCreativesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

