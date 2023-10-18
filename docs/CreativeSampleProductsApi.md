# TencentAds\CreativeSampleProductsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeSampleProductsGet**](CreativeSampleProductsApi.md#CreativeSampleProductsGet) | **Get** /creative_sample_products/get | 获取创意示例商品列表


# **CreativeSampleProductsGet**
> CreativeSampleProductsGetResponse CreativeSampleProductsGet(ctx, accountId, productCatalogId, optional)
获取创意示例商品列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
 **optional** | ***CreativeSampleProductsApiCreativeSampleProductsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativeSampleProductsApiCreativeSampleProductsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **productOuterIds** | [**optional.Interface of []string**](string.md)|  | 
 **productSeriesId** | **optional.Int64**|  | 
 **templateId** | **optional.Int64**|  | 
 **templateType** | **optional.String**|  | 
 **imageId** | **optional.String**|  | 
 **videoId** | **optional.String**|  | 
 **productFields** | [**optional.Interface of []string**](string.md)|  | 
 **limit** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CreativeSampleProductsGetResponse**](CreativeSampleProductsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

