# TencentAds\WalletInvoiceApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WalletInvoiceGet**](WalletInvoiceApi.md#WalletInvoiceGet) | **Get** /wallet_invoice/get | 共享钱包流水相关信息查询


# **WalletInvoiceGet**
> WalletInvoiceGetResponse WalletInvoiceGet(ctx, accountId, walletIdList, dateRange, optional)
共享钱包流水相关信息查询

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **walletIdList** | **string**|  | 
  **dateRange** | [**WalletDateRangeTransaction**](WalletDateRangeTransaction.md)|  | 
 **optional** | ***WalletInvoiceApiWalletInvoiceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WalletInvoiceApiWalletInvoiceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fundType** | **optional.String**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**WalletInvoiceGetResponse**](WalletInvoiceGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

