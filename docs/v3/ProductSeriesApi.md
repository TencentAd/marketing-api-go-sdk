# TencentAds\ProductSeriesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductSeriesAdd**](ProductSeriesApi.md#ProductSeriesAdd) | **Post** /product_series/add | 创建商品系列
[**ProductSeriesGet**](ProductSeriesApi.md#ProductSeriesGet) | **Get** /product_series/get | 获取商品系列


# **ProductSeriesAdd**
> ProductSeriesAddResponse ProductSeriesAdd(ctx, data)
创建商品系列

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductSeriesAddRequest**](ProductSeriesAddRequest.md)|  | 

### Return type

[**ProductSeriesAddResponse**](ProductSeriesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductSeriesGet**
> ProductSeriesGetResponse ProductSeriesGet(ctx, accountId, catalogId, optional)
获取商品系列

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **catalogId** | **int64**|  | 
 **optional** | ***ProductSeriesApiProductSeriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductSeriesApiProductSeriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []ProductSeriesSearchFilteringStruct**](ProductSeriesSearchFilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductSeriesGetResponse**](ProductSeriesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

