# TencentAds\WatermarksApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatermarksAdd**](WatermarksApi.md#WatermarksAdd) | **Post** /watermarks/add | 批量添加警示语
[**WatermarksGet**](WatermarksApi.md#WatermarksGet) | **Get** /watermarks/get | 查询警示语添加记录


# **WatermarksAdd**
> WatermarksAddResponse WatermarksAdd(ctx, data)
批量添加警示语

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WatermarksAddRequest**](WatermarksAddRequest.md)|  | 

### Return type

[**WatermarksAddResponse**](WatermarksAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatermarksGet**
> WatermarksGetResponse WatermarksGet(ctx, optional)
查询警示语添加记录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***WatermarksApiWatermarksGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatermarksApiWatermarksGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **optional.Int64**|  | 
 **organizationId** | **optional.Int64**|  | 
 **filtering** | [**optional.Interface of []QueryConditionStruct**](QueryConditionStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WatermarksGetResponse**](WatermarksGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

