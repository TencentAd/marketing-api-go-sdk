# TencentAds\AsyncTasksApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncTasksAdd**](AsyncTasksApi.md#AsyncTasksAdd) | **Post** /async_tasks/add | 创建异步任务
[**AsyncTasksGet**](AsyncTasksApi.md#AsyncTasksGet) | **Get** /async_tasks/get | 获取异步任务


# **AsyncTasksAdd**
> AsyncTasksAddResponse AsyncTasksAdd(ctx, data)
创建异步任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AsyncTasksAddRequest**](AsyncTasksAddRequest.md)|  | 

### Return type

[**AsyncTasksAddResponse**](AsyncTasksAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AsyncTasksGet**
> AsyncTasksGetResponse AsyncTasksGet(ctx, accountId, optional)
获取异步任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AsyncTasksApiAsyncTasksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsyncTasksApiAsyncTasksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []AsyncTaskFilteringStruct**](AsyncTaskFilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AsyncTasksGetResponse**](AsyncTasksGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

