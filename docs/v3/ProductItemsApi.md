# TencentAds\ProductItemsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductItemsAdd**](ProductItemsApi.md#ProductItemsAdd) | **Post** /product_items/add | 添加商品
[**ProductItemsBatchUpdate**](ProductItemsApi.md#ProductItemsBatchUpdate) | **Post** /product_items/batch_update | 批量更新商品信息
[**ProductItemsUpdate**](ProductItemsApi.md#ProductItemsUpdate) | **Post** /product_items/update | 更新商品信息


# **ProductItemsAdd**
> ProductItemsAddResponse ProductItemsAdd(ctx, data)
添加商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductItemsAddRequest**](ProductItemsAddRequest.md)|  | 

### Return type

[**ProductItemsAddResponse**](ProductItemsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductItemsBatchUpdate**
> ProductItemsBatchUpdateResponse ProductItemsBatchUpdate(ctx, data)
批量更新商品信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductItemsBatchUpdateRequest**](ProductItemsBatchUpdateRequest.md)|  | 

### Return type

[**ProductItemsBatchUpdateResponse**](ProductItemsBatchUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProductItemsUpdate**
> ProductItemsUpdateResponse ProductItemsUpdate(ctx, data)
更新商品信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProductItemsUpdateRequest**](ProductItemsUpdateRequest.md)|  | 

### Return type

[**ProductItemsUpdateResponse**](ProductItemsUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

