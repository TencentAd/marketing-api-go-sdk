# TencentAds\CustomAudienceReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAudienceReportsGet**](CustomAudienceReportsApi.md#CustomAudienceReportsGet) | **Get** /custom_audience_reports/get | 人群报表


# **CustomAudienceReportsGet**
> CustomAudienceReportsGetResponse CustomAudienceReportsGet(ctx, accountId, filtering, dateRange, optional)
人群报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **filtering** | [**[]FilteringStruct**](FilteringStruct.md)|  | 
  **dateRange** | [**DateRange**](DateRange.md)|  | 
 **optional** | ***CustomAudienceReportsApiCustomAudienceReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudienceReportsApiCustomAudienceReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **groupBy** | [**optional.Interface of []string**](string.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomAudienceReportsGetResponse**](CustomAudienceReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

