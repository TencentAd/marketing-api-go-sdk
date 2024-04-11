# TencentAds\ProductItemsVerticalsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProductItemsVerticalsGet**](ProductItemsVerticalsApi.md#ProductItemsVerticalsGet) | **Get** /product_items_verticals/get | 行业列表


# **ProductItemsVerticalsGet**
> ProductItemsVerticalsGetResponse ProductItemsVerticalsGet(ctx, optional)
行业列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductItemsVerticalsApiProductItemsVerticalsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductItemsVerticalsApiProductItemsVerticalsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProductItemsVerticalsGetResponse**](ProductItemsVerticalsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

