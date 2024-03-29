# TencentAds\UserActionSetReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserActionSetReportsGet**](UserActionSetReportsApi.md#UserActionSetReportsGet) | **Get** /user_action_set_reports/get | 获取用户行为数据源报表


# **UserActionSetReportsGet**
> UserActionSetReportsGetResponse UserActionSetReportsGet(ctx, accountId, userActionSetId, dateRange, timeGranularity, optional)
获取用户行为数据源报表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **userActionSetId** | **int64**|  | 
  **dateRange** | [**DateRangeDn**](DateRangeDn.md)|  | 
  **timeGranularity** | **string**|  | 
 **optional** | ***UserActionSetReportsApiUserActionSetReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UserActionSetReportsApiUserActionSetReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **aggregation** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**UserActionSetReportsGetResponse**](UserActionSetReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

