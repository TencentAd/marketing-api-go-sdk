# TencentAds\FundStatementsDetailedApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FundStatementsDetailedGet**](FundStatementsDetailedApi.md#FundStatementsDetailedGet) | **Get** /fund_statements_detailed/get | 获取资金流水


# **FundStatementsDetailedGet**
> FundStatementsDetailedGetResponse FundStatementsDetailedGet(ctx, accountId, fundType, dateRange, optional)
获取资金流水

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **fundType** | **string**|  | 
  **dateRange** | [**DateRangeTransaction**](DateRangeTransaction.md)|  | 
 **optional** | ***FundStatementsDetailedApiFundStatementsDetailedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FundStatementsDetailedApiFundStatementsDetailedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **primaryKey** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**FundStatementsDetailedGetResponse**](FundStatementsDetailedGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

