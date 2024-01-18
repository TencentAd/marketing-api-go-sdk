# TencentAds\MaterialAuditApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MaterialAuditList**](MaterialAuditApi.md#MaterialAuditList) | **Post** /material_audit/list | 获取预审结果
[**MaterialAuditSubmit**](MaterialAuditApi.md#MaterialAuditSubmit) | **Post** /material_audit/submit | 素材批量送审


# **MaterialAuditList**
> MaterialAuditListResponse MaterialAuditList(ctx, data)
获取预审结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialAuditListRequest**](MaterialAuditListRequest.md)|  | 

### Return type

[**MaterialAuditListResponse**](MaterialAuditListResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MaterialAuditSubmit**
> MaterialAuditSubmitResponse MaterialAuditSubmit(ctx, data)
素材批量送审

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MaterialAuditSubmitRequest**](MaterialAuditSubmitRequest.md)|  | 

### Return type

[**MaterialAuditSubmitResponse**](MaterialAuditSubmitResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

