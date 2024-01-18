# TencentAds\LocalStoresCategoriesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalStoresCategoriesGet**](LocalStoresCategoriesApi.md#LocalStoresCategoriesGet) | **Get** /local_stores_categories/get | 查询门店类目


# **LocalStoresCategoriesGet**
> LocalStoresCategoriesGetResponse LocalStoresCategoriesGet(ctx, optional)
查询门店类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LocalStoresCategoriesApiLocalStoresCategoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStoresCategoriesApiLocalStoresCategoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verticalId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LocalStoresCategoriesGetResponse**](LocalStoresCategoriesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

