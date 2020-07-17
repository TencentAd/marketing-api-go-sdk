# TencentAds\WechatFundStatementsDetailedApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatFundStatementsDetailedGet**](WechatFundStatementsDetailedApi.md#WechatFundStatementsDetailedGet) | **Get** /wechat_fund_statements_detailed/get | 获取微信资金账户流水信息


# **WechatFundStatementsDetailedGet**
> WechatFundStatementsDetailedGetResponse WechatFundStatementsDetailedGet(ctx, tradeType, dateRange, optional)
获取微信资金账户流水信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tradeType** | **string**|  | 
  **dateRange** | [**DateRange**](DateRange.md)|  | 
 **optional** | ***WechatFundStatementsDetailedApiWechatFundStatementsDetailedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatFundStatementsDetailedApiWechatFundStatementsDetailedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatFundStatementsDetailedGetResponse**](WechatFundStatementsDetailedGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

