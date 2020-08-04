# TencentAds\ProductItemsListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductItemsListGet**](ProductItemsListApi.md#ProductItemsListGet) | **Get** /product_items_list/get | 商品列表


# **ProductItemsListGet**
> ProductItemsListGetResponse ProductItemsListGet(ctx, accountId, productCatalogId, pageSize, page, optional)
商品列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
  **pageSize** | **int64**|  | 
  **page** | **int64**|  | 
 **optional** | ***ProductItemsListApiProductItemsListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductItemsListApiProductItemsListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **sortByProductLastModTime** | **optional.String**|  | 
 **sortByProductName** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductItemsListGetResponse**](ProductItemsListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

