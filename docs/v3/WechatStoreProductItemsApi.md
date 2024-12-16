# TencentAds\WechatStoreProductItemsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatStoreProductItemsGet**](WechatStoreProductItemsApi.md#WechatStoreProductItemsGet) | **Get** /wechat_store_product_items/get | 获取微信小店商品


# **WechatStoreProductItemsGet**
> WechatStoreProductItemsGetResponse WechatStoreProductItemsGet(ctx, accountId, productCatalogId, optional)
获取微信小店商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
 **optional** | ***WechatStoreProductItemsApiWechatStoreProductItemsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WechatStoreProductItemsApiWechatStoreProductItemsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filtering** | [**optional.Interface of []WechatStoreProductFilteringStruct**](WechatStoreProductFilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WechatStoreProductItemsGetResponse**](WechatStoreProductItemsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

