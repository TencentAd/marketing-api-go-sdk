# TencentAds\WalletTransferApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletTransferAdd**](WalletTransferApi.md#WalletTransferAdd) | **Post** /wallet_transfer/add | 发起代理商与钱包之间转账


# **WalletTransferAdd**
> WalletTransferAddResponse WalletTransferAdd(ctx, data)
发起代理商与钱包之间转账

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WalletTransferAddRequest**](WalletTransferAddRequest.md)|  | 

### Return type

[**WalletTransferAddResponse**](WalletTransferAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

