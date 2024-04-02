# TencentAds\ProductsSystemStatusApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductsSystemStatusGet**](ProductsSystemStatusApi.md#ProductsSystemStatusGet) | **Get** /products_system_status/get | 获取审核失败的商品


# **ProductsSystemStatusGet**
> ProductsSystemStatusGetResponse ProductsSystemStatusGet(ctx, accountId, productCatalogId, optional)
获取审核失败的商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
 **optional** | ***ProductsSystemStatusApiProductsSystemStatusGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductsSystemStatusApiProductsSystemStatusGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **feedId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductsSystemStatusGetResponse**](ProductsSystemStatusGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

