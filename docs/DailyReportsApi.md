# TencentAds\DailyReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DailyReportsGet**](DailyReportsApi.md#DailyReportsGet) | **Get** /daily_reports/get | 获取日报表


# **DailyReportsGet**
> DailyReportsGetResponse DailyReportsGet(ctx, accountId, level, dateRange, optional)
获取日报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **level** | **string**|  | 
  **dateRange** | [**ReportDateRange**](ReportDateRange.md)|  | 
 **optional** | ***DailyReportsApiDailyReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DailyReportsApiDailyReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **groupBy** | [**optional.Interface of []string**](string.md)|  | 
 **orderBy** | [**optional.Interface of []OrderByStruct**](OrderByStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **timeLine** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)|  | 
 **weixinOfficialAccountsUpgradeEnabled** | **optional.Bool**|  | 
 **adqAccountsUpgradeEnabled** | **optional.Bool**|  | 

### Return type

[**DailyReportsGetResponse**](DailyReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

