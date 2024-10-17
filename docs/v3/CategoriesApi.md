# TencentAds\CategoriesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesGet**](CategoriesApi.md#CategoriesGet) | **Get** /categories/get | 获取腾讯系商品类目


# **CategoriesGet**
> CategoriesGetResponse CategoriesGet(ctx, verticalId, page, pageSize, optional)
获取腾讯系商品类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verticalId** | **int64**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***CategoriesApiCategoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesApiCategoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **level** | **optional.Int64**|  | 
 **categoryId** | **optional.Int64**|  | 
 **categoryName** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CategoriesGetResponse**](CategoriesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

