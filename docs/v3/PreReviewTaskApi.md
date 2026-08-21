# TencentAds\PreReviewTaskApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreReviewTaskAdd**](PreReviewTaskApi.md#PreReviewTaskAdd) | **Post** /pre_review_task/add | 异步预审任务提交
[**PreReviewTaskGet**](PreReviewTaskApi.md#PreReviewTaskGet) | **Get** /pre_review_task/get | 异步预审结果获取


# **PreReviewTaskAdd**
> PreReviewTaskAddResponse PreReviewTaskAdd(ctx, data)
异步预审任务提交

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**PreReviewTaskAddRequest**](PreReviewTaskAddRequest.md)|  | 

### Return type

[**PreReviewTaskAddResponse**](PreReviewTaskAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PreReviewTaskGet**
> PreReviewTaskGetResponse PreReviewTaskGet(ctx, taskId, optional)
异步预审结果获取

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***PreReviewTaskApiPreReviewTaskGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PreReviewTaskApiPreReviewTaskGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**PreReviewTaskGetResponse**](PreReviewTaskGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

