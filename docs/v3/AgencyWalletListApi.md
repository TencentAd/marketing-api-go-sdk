# TencentAds\AgencyWalletListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AgencyWalletListGet**](AgencyWalletListApi.md#AgencyWalletListGet) | **Get** /agency_wallet_list/get | 获取代理商创建的共享钱包信息列表


# **AgencyWalletListGet**
> AgencyWalletListGetResponse AgencyWalletListGet(ctx, accountId, page, pageSize, optional)
获取代理商创建的共享钱包信息列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **page** | **int64**|  | 
  **pageSize** | **int64**|  | 
 **optional** | ***AgencyWalletListApiAgencyWalletListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AgencyWalletListApiAgencyWalletListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mdmId** | **optional.Int64**|  | 
 **walletId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AgencyWalletListGetResponse**](AgencyWalletListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

