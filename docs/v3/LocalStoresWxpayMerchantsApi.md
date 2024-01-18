# TencentAds\LocalStoresWxpayMerchantsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalStoresWxpayMerchantsGet**](LocalStoresWxpayMerchantsApi.md#LocalStoresWxpayMerchantsGet) | **Get** /local_stores_wxpay_merchants/get | 查询微信支付商户号


# **LocalStoresWxpayMerchantsGet**
> LocalStoresWxpayMerchantsGetResponse LocalStoresWxpayMerchantsGet(ctx, accountId, optional)
查询微信支付商户号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***LocalStoresWxpayMerchantsApiLocalStoresWxpayMerchantsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocalStoresWxpayMerchantsApiLocalStoresWxpayMerchantsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LocalStoresWxpayMerchantsGetResponse**](LocalStoresWxpayMerchantsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

