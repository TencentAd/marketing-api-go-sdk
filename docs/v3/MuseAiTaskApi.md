# TencentAds\MuseAiTaskApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MuseAiTaskAdd**](MuseAiTaskApi.md#MuseAiTaskAdd) | **Post** /muse_ai_task/add | 创建妙思任务接口
[**MuseAiTaskGet**](MuseAiTaskApi.md#MuseAiTaskGet) | **Get** /muse_ai_task/get | 获取妙思任务结果接口


# **MuseAiTaskAdd**
> MuseAiTaskAddResponse MuseAiTaskAdd(ctx, data)
创建妙思任务接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MuseAiTaskAddRequest**](MuseAiTaskAddRequest.md)|  | 

### Return type

[**MuseAiTaskAddResponse**](MuseAiTaskAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MuseAiTaskGet**
> MuseAiTaskGetResponse MuseAiTaskGet(ctx, accountId, taskId, optional)
获取妙思任务结果接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **taskId** | **int64**|  | 
 **optional** | ***MuseAiTaskApiMuseAiTaskGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MuseAiTaskApiMuseAiTaskGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**MuseAiTaskGetResponse**](MuseAiTaskGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

