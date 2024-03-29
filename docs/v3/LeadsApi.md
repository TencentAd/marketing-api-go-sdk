# TencentAds\LeadsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsAdd**](LeadsApi.md#LeadsAdd) | **Post** /leads/add | 外部线索数据导入
[**LeadsUpdate**](LeadsApi.md#LeadsUpdate) | **Post** /leads/update | 更新线索基本信息


# **LeadsAdd**
> LeadsAddResponse LeadsAdd(ctx, data)
外部线索数据导入

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsAddRequest**](LeadsAddRequest.md)|  | 

### Return type

[**LeadsAddResponse**](LeadsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LeadsUpdate**
> LeadsUpdateResponse LeadsUpdate(ctx, data)
更新线索基本信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsUpdateRequest**](LeadsUpdateRequest.md)|  | 

### Return type

[**LeadsUpdateResponse**](LeadsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

