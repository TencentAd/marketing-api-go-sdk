# TencentAds\OuterCluesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OuterCluesActionTypeReport**](OuterCluesApi.md#OuterCluesActionTypeReport) | **Post** /outer_clues/action_type_report | 线索上报DMP平台
[**OuterCluesAdd**](OuterCluesApi.md#OuterCluesAdd) | **Post** /outer_clues/add | 外部线索数据导入
[**OuterCluesUpdate**](OuterCluesApi.md#OuterCluesUpdate) | **Post** /outer_clues/update | 外部线索状态更新


# **OuterCluesActionTypeReport**
> OuterCluesActionTypeReportResponse OuterCluesActionTypeReport(ctx, data)
线索上报DMP平台

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesActionTypeReportRequest**](OuterCluesActionTypeReportRequest.md)|  | 

### Return type

[**OuterCluesActionTypeReportResponse**](OuterCluesActionTypeReportResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OuterCluesAdd**
> OuterCluesAddResponse OuterCluesAdd(ctx, data)
外部线索数据导入

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesAddRequest**](OuterCluesAddRequest.md)|  | 

### Return type

[**OuterCluesAddResponse**](OuterCluesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OuterCluesUpdate**
> OuterCluesUpdateResponse OuterCluesUpdate(ctx, data)
外部线索状态更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesUpdateRequest**](OuterCluesUpdateRequest.md)|  | 

### Return type

[**OuterCluesUpdateResponse**](OuterCluesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

