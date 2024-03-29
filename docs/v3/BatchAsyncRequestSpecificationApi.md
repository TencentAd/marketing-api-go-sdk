# TencentAds\BatchAsyncRequestSpecificationApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchAsyncRequestSpecificationGet**](BatchAsyncRequestSpecificationApi.md#BatchAsyncRequestSpecificationGet) | **Get** /batch_async_request_specification/get | 获取批量异步请求任务详情


# **BatchAsyncRequestSpecificationGet**
> BatchAsyncRequestSpecificationGetResponse BatchAsyncRequestSpecificationGet(ctx, accountId, taskId, optional)
获取批量异步请求任务详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **taskId** | **int64**|  | 
 **optional** | ***BatchAsyncRequestSpecificationApiBatchAsyncRequestSpecificationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BatchAsyncRequestSpecificationApiBatchAsyncRequestSpecificationGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BatchAsyncRequestSpecificationGetResponse**](BatchAsyncRequestSpecificationGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

