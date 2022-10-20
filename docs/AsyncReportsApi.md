# TencentAds\AsyncReportsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncReportsAdd**](AsyncReportsApi.md#AsyncReportsAdd) | **Post** /async_reports/add | 创建异步报表任务
[**AsyncReportsGet**](AsyncReportsApi.md#AsyncReportsGet) | **Get** /async_reports/get | 获取异步报表任务


# **AsyncReportsAdd**
> AsyncReportsAddResponse AsyncReportsAdd(ctx, data)
创建异步报表任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AsyncReportsAddRequest**](AsyncReportsAddRequest.md)|  | 

### Return type

[**AsyncReportsAddResponse**](AsyncReportsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AsyncReportsGet**
> AsyncReportsGetResponse AsyncReportsGet(ctx, accountId, optional)
获取异步报表任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AsyncReportsApiAsyncReportsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsyncReportsApiAsyncReportsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **weixinOfficialAccountsUpgradeEnabled** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AsyncReportsGetResponse**](AsyncReportsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

