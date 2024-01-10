# TencentAds\MergeFundTypeDailyBalanceReportApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MergeFundTypeDailyBalanceReportGet**](MergeFundTypeDailyBalanceReportApi.md#MergeFundTypeDailyBalanceReportGet) | **Get** /merge_fund_type_daily_balance_report/get | 获取资金合并类型资金账户日结明细


# **MergeFundTypeDailyBalanceReportGet**
> MergeFundTypeDailyBalanceReportGetResponse MergeFundTypeDailyBalanceReportGet(ctx, accountId, dateRange, optional)
获取资金合并类型资金账户日结明细

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dateRange** | [**DateRangeTransaction**](DateRangeTransaction.md)|  | 
 **optional** | ***MergeFundTypeDailyBalanceReportApiMergeFundTypeDailyBalanceReportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MergeFundTypeDailyBalanceReportApiMergeFundTypeDailyBalanceReportGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MergeFundTypeDailyBalanceReportGetResponse**](MergeFundTypeDailyBalanceReportGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

