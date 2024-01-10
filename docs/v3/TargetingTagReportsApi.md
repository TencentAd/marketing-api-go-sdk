# TencentAds\TargetingTagReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetingTagReportsGet**](TargetingTagReportsApi.md#TargetingTagReportsGet) | **Get** /targeting_tag_reports/get | 获取定向标签报表


# **TargetingTagReportsGet**
> TargetingTagReportsGetResponse TargetingTagReportsGet(ctx, accountId, type_, level, dateRange, groupBy, fields, optional)
获取定向标签报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **type_** | **string**|  | 
  **level** | **string**|  | 
  **dateRange** | [**ReportDateRange**](ReportDateRange.md)|  | 
  **groupBy** | [**[]string**](string.md)|  | 
  **fields** | [**[]string**](string.md)|  | 
 **optional** | ***TargetingTagReportsApiTargetingTagReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingTagReportsApiTargetingTagReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **filtering** | [**optional.Interface of []TargetingFilteringStruct**](TargetingFilteringStruct.md)|  | 
 **orderBy** | [**optional.Interface of []OrderByStruct**](OrderByStruct.md)|  | 
 **timeLine** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 

### Return type

[**TargetingTagReportsGetResponse**](TargetingTagReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

