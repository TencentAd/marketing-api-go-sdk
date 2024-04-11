# TencentAds\BarrageApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BarrageAdd**](BarrageApi.md#BarrageAdd) | **Post** /barrage/add | 添加弹幕
[**BarrageGet**](BarrageApi.md#BarrageGet) | **Get** /barrage/get | 获取弹幕


# **BarrageAdd**
> BarrageAddResponse BarrageAdd(ctx, data)
添加弹幕

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BarrageAddRequest**](BarrageAddRequest.md)|  | 

### Return type

[**BarrageAddResponse**](BarrageAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BarrageGet**
> BarrageGetResponse BarrageGet(ctx, accountId, optional)
获取弹幕

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***BarrageApiBarrageGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BarrageApiBarrageGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idList** | [**optional.Interface of []int64**](int64.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**BarrageGetResponse**](BarrageGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

