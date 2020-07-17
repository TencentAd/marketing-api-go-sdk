# TencentAds\FundStatementsDailyApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FundStatementsDailyGet**](FundStatementsDailyApi.md#FundStatementsDailyGet) | **Get** /fund_statements_daily/get | 获取资金账户日结明细


# **FundStatementsDailyGet**
> FundStatementsDailyGetResponse FundStatementsDailyGet(ctx, accountId, fundType, dateRange, optional)
获取资金账户日结明细

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **fundType** | **string**|  | 
  **dateRange** | [**DateRange**](DateRange.md)|  | 
 **optional** | ***FundStatementsDailyApiFundStatementsDailyGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundStatementsDailyApiFundStatementsDailyGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **tradeType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**FundStatementsDailyGetResponse**](FundStatementsDailyGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

