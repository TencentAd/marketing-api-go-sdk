# TencentAds\LeadsCallRecordApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallRecordGet**](LeadsCallRecordApi.md#LeadsCallRecordGet) | **Get** /leads_call_record/get | 获取通话结果


# **LeadsCallRecordGet**
> LeadsCallRecordGetResponse LeadsCallRecordGet(ctx, accountId, optional)
获取通话结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LeadsCallRecordApiLeadsCallRecordGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeadsCallRecordApiLeadsCallRecordGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **leadsId** | **optional.Int64**|  | 
 **outerLeadsId** | **optional.String**|  | 
 **requestId** | **optional.String**|  | 
 **contactId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LeadsCallRecordGetResponse**](LeadsCallRecordGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

