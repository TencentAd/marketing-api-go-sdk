# TencentAds\OperationLogListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OperationLogListGet**](OperationLogListApi.md#OperationLogListGet) | **Get** /operation_log_list/get | 获取操作日志列表


# **OperationLogListGet**
> OperationLogListGetResponse OperationLogListGet(ctx, accountId, operationObjectType, startDate, endDate, page, pageSize, optional)
获取操作日志列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **operationObjectType** | **string**|  | 
  **startDate** | **string**|  | 
  **endDate** | **string**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***OperationLogListApiOperationLogListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OperationLogListApiOperationLogListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **objectId** | **optional.Int64**|  | 
 **operatorPlatformList** | [**optional.Interface of []string**](string.md)|  | 
 **operationActionList** | [**optional.Interface of []string**](string.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**OperationLogListGetResponse**](OperationLogListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

