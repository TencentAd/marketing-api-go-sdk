# TencentAds\PropertyFileSessionsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertyFileSessionsAdd**](PropertyFileSessionsApi.md#PropertyFileSessionsAdd) | **Post** /property_file_sessions/add | 创建属性数据文件上传会话
[**PropertyFileSessionsUpdate**](PropertyFileSessionsApi.md#PropertyFileSessionsUpdate) | **Post** /property_file_sessions/update | 提交属性数据文件上传会话


# **PropertyFileSessionsAdd**
> PropertyFileSessionsAddResponse PropertyFileSessionsAdd(ctx, data)
创建属性数据文件上传会话

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PropertyFileSessionsAddRequest**](PropertyFileSessionsAddRequest.md)|  | 

### Return type

[**PropertyFileSessionsAddResponse**](PropertyFileSessionsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PropertyFileSessionsUpdate**
> PropertyFileSessionsUpdateResponse PropertyFileSessionsUpdate(ctx, data)
提交属性数据文件上传会话

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PropertyFileSessionsUpdateRequest**](PropertyFileSessionsUpdateRequest.md)|  | 

### Return type

[**PropertyFileSessionsUpdateResponse**](PropertyFileSessionsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

