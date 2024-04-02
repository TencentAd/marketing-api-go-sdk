# TencentAds\BatchAsyncRequestsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchAsyncRequestsAdd**](BatchAsyncRequestsApi.md#BatchAsyncRequestsAdd) | **Post** /batch_async_requests/add | 创建批量异步请求任务
[**BatchAsyncRequestsGet**](BatchAsyncRequestsApi.md#BatchAsyncRequestsGet) | **Get** /batch_async_requests/get | 获取批量异步请求任务列表


# **BatchAsyncRequestsAdd**
> BatchAsyncRequestsAddResponse BatchAsyncRequestsAdd(ctx, data)
创建批量异步请求任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BatchAsyncRequestsAddRequest**](BatchAsyncRequestsAddRequest.md)|  | 

### Return type

[**BatchAsyncRequestsAddResponse**](BatchAsyncRequestsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchAsyncRequestsGet**
> BatchAsyncRequestsGetResponse BatchAsyncRequestsGet(ctx, accountId, optional)
获取批量异步请求任务列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***BatchAsyncRequestsApiBatchAsyncRequestsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchAsyncRequestsApiBatchAsyncRequestsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []BatchAsyncTaskFilteringStruct**](BatchAsyncTaskFilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BatchAsyncRequestsGetResponse**](BatchAsyncRequestsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

