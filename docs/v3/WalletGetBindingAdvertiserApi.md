# TencentAds\WalletGetBindingAdvertiserApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletGetBindingAdvertiserGet**](WalletGetBindingAdvertiserApi.md#WalletGetBindingAdvertiserGet) | **Get** /wallet_get_binding_advertiser/get | 查询单个共享钱包下的关联账户信息


# **WalletGetBindingAdvertiserGet**
> WalletGetBindingAdvertiserGetResponse WalletGetBindingAdvertiserGet(ctx, accountId, walletId, page, pageSize, optional)
查询单个共享钱包下的关联账户信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **walletId** | **int64**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***WalletGetBindingAdvertiserApiWalletGetBindingAdvertiserGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletGetBindingAdvertiserApiWalletGetBindingAdvertiserGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WalletGetBindingAdvertiserGetResponse**](WalletGetBindingAdvertiserGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

