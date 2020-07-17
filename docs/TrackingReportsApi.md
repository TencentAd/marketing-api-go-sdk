# TencentAds\TrackingReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TrackingReportsGet**](TrackingReportsApi.md#TrackingReportsGet) | **Get** /tracking_reports/get | 点击追踪报表


# **TrackingReportsGet**
> TrackingReportsGetResponse TrackingReportsGet(ctx, accountId, dateRange, optional)
点击追踪报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dateRange** | [**DateRange**](DateRange.md)|  | 
 **optional** | ***TrackingReportsApiTrackingReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TrackingReportsApiTrackingReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeGranularity** | **optional.String**|  | 
 **promotedObjectType** | **optional.String**|  | 
 **promotedObjectId** | **optional.String**|  | 
 **feedbackUrl** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**TrackingReportsGetResponse**](TrackingReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

