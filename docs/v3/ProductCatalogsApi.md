# TencentAds\ProductCatalogsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductCatalogsAdd**](ProductCatalogsApi.md#ProductCatalogsAdd) | **Post** /product_catalogs/add | 创建商品库
[**ProductCatalogsGet**](ProductCatalogsApi.md#ProductCatalogsGet) | **Get** /product_catalogs/get | 获取商品库


# **ProductCatalogsAdd**
> ProductCatalogsAddResponse ProductCatalogsAdd(ctx, data)
创建商品库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductCatalogsAddRequest**](ProductCatalogsAddRequest.md)|  | 

### Return type

[**ProductCatalogsAddResponse**](ProductCatalogsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductCatalogsGet**
> ProductCatalogsGetResponse ProductCatalogsGet(ctx, accountId, optional)
获取商品库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ProductCatalogsApiProductCatalogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductCatalogsApiProductCatalogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **catalogId** | **optional.Int64**|  | 
 **catalogName** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductCatalogsGetResponse**](ProductCatalogsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

