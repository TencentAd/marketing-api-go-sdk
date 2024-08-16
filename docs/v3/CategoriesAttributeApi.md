# TencentAds\CategoriesAttributeApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesAttributeGet**](CategoriesAttributeApi.md#CategoriesAttributeGet) | **Get** /categories_attribute/get | 获取商品属性


# **CategoriesAttributeGet**
> CategoriesAttributeGetResponse CategoriesAttributeGet(ctx, verticalId, page, pageSize, optional)
获取商品属性

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **verticalId** | **int64**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***CategoriesAttributeApiCategoriesAttributeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CategoriesAttributeApiCategoriesAttributeGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lastCategoryIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CategoriesAttributeGetResponse**](CategoriesAttributeGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

