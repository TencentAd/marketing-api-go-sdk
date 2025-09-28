# TencentAds\SearchAdgroupsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAdgroupsAdd**](SearchAdgroupsApi.md#SearchAdgroupsAdd) | **Post** /search_adgroups/add | 创建搜索广告
[**SearchAdgroupsUpdate**](SearchAdgroupsApi.md#SearchAdgroupsUpdate) | **Post** /search_adgroups/update | 更新搜索广告


# **SearchAdgroupsAdd**
> SearchAdgroupsAddResponse SearchAdgroupsAdd(ctx, data)
创建搜索广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SearchAdgroupsAddRequest**](SearchAdgroupsAddRequest.md)|  | 

### Return type

[**SearchAdgroupsAddResponse**](SearchAdgroupsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAdgroupsUpdate**
> SearchAdgroupsUpdateResponse SearchAdgroupsUpdate(ctx, data)
更新搜索广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SearchAdgroupsUpdateRequest**](SearchAdgroupsUpdateRequest.md)|  | 

### Return type

[**SearchAdgroupsUpdateResponse**](SearchAdgroupsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

