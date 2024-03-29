# TencentAds\LeadsStatusApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsStatusUpdate**](LeadsStatusApi.md#LeadsStatusUpdate) | **Post** /leads_status/update | 外部线索状态更新


# **LeadsStatusUpdate**
> LeadsStatusUpdateResponse LeadsStatusUpdate(ctx, data)
外部线索状态更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsStatusUpdateRequest**](LeadsStatusUpdateRequest.md)|  | 

### Return type

[**LeadsStatusUpdateResponse**](LeadsStatusUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

