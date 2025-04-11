# TencentAds\WalletCreateApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletCreateAdd**](WalletCreateApi.md#WalletCreateAdd) | **Post** /wallet_create/add | 新建共享钱包


# **WalletCreateAdd**
> WalletCreateAddResponse WalletCreateAdd(ctx, data)
新建共享钱包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WalletCreateAddRequest**](WalletCreateAddRequest.md)|  | 

### Return type

[**WalletCreateAddResponse**](WalletCreateAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

