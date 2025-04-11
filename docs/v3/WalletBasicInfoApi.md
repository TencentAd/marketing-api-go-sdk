# TencentAds\WalletBasicInfoApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletBasicInfoGet**](WalletBasicInfoApi.md#WalletBasicInfoGet) | **Get** /wallet_basic_info/get | 通过钱包id去查询共享钱包基础信息


# **WalletBasicInfoGet**
> WalletBasicInfoGetResponse WalletBasicInfoGet(ctx, accountId, walletId, optional)
通过钱包id去查询共享钱包基础信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **walletId** | **int64**|  | 
 **optional** | ***WalletBasicInfoApiWalletBasicInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletBasicInfoApiWalletBasicInfoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WalletBasicInfoGetResponse**](WalletBasicInfoGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

