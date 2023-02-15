# TencentAds\LeadsCallRecordApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallRecordGet**](LeadsCallRecordApi.md#LeadsCallRecordGet) | **Post** /leads_call_record/get | 获取通话结果


# **LeadsCallRecordGet**
> LeadsCallRecordGetResponse LeadsCallRecordGet(ctx, data)
获取通话结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsCallRecordGetRequest**](LeadsCallRecordGetRequest.md)|  | 

### Return type

[**LeadsCallRecordGetResponse**](LeadsCallRecordGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

