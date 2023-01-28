# TencentAds\FileDispatchApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileDispatchUpdate**](FileDispatchApi.md#FileDispatchUpdate) | **Post** /file_dispatch/update | 更新文件分发关系


# **FileDispatchUpdate**
> FileDispatchUpdateResponse FileDispatchUpdate(ctx, data)
更新文件分发关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FileDispatchUpdateRequest**](FileDispatchUpdateRequest.md)|  | 

### Return type

[**FileDispatchUpdateResponse**](FileDispatchUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

