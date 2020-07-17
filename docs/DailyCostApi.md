# TencentAds\DailyCostApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DailyCostGet**](DailyCostApi.md#DailyCostGet) | **Get** /daily_cost/get | 【待废弃，请使用新接口wechat_daily_cost/get】获取微信账户实时消耗


# **DailyCostGet**
> DailyCostGetResponse DailyCostGet(ctx, dateRange, optional)
【待废弃，请使用新接口wechat_daily_cost/get】获取微信账户实时消耗

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dateRange** | [**ReportDateRange**](ReportDateRange.md)|  | 
 **optional** | ***DailyCostApiDailyCostGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DailyCostApiDailyCostGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountId** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DailyCostGetResponse**](DailyCostGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

