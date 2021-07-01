# TencentAds\LeadCluesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadCluesGet**](LeadCluesApi.md#LeadCluesGet) | **Post** /lead_clues/get | 获取线索列表
[**LeadCluesUpdate**](LeadCluesApi.md#LeadCluesUpdate) | **Post** /lead_clues/update | 回传线索状态


# **LeadCluesGet**
> LeadCluesGetResponse LeadCluesGet(ctx, data)
获取线索列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadCluesGetRequest**](LeadCluesGetRequest.md)|  | 

### Return type

[**LeadCluesGetResponse**](LeadCluesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LeadCluesUpdate**
> LeadCluesUpdateResponse LeadCluesUpdate(ctx, data)
回传线索状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadCluesUpdateRequest**](LeadCluesUpdateRequest.md)|  | 

### Return type

[**LeadCluesUpdateResponse**](LeadCluesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

