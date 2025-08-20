# TencentAds\IllegalComplaintApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IllegalComplaintAdd**](IllegalComplaintApi.md#IllegalComplaintAdd) | **Post** /illegal_complaint/add | 新增广告主违规申述
[**IllegalComplaintGet**](IllegalComplaintApi.md#IllegalComplaintGet) | **Post** /illegal_complaint/get | 获取直客广告主违规申述列表


# **IllegalComplaintAdd**
> IllegalComplaintAddResponse IllegalComplaintAdd(ctx, accountId, illegalOrderId, complaintReason, file)
新增广告主违规申述

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **illegalOrderId** | **string**|  | 
  **complaintReason** | **string**|  | 
  **file** | ***os.File**|  | 

### Return type

[**IllegalComplaintAddResponse**](IllegalComplaintAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **IllegalComplaintGet**
> IllegalComplaintGetResponse IllegalComplaintGet(ctx, data)
获取直客广告主违规申述列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**IllegalComplaintGetRequest**](IllegalComplaintGetRequest.md)|  | 

### Return type

[**IllegalComplaintGetResponse**](IllegalComplaintGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

