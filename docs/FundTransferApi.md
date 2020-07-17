# TencentAds\FundTransferApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FundTransferAdd**](FundTransferApi.md#FundTransferAdd) | **Post** /fund_transfer/add | 发起代理商与子客户之间转账


# **FundTransferAdd**
> FundTransferAddResponse FundTransferAdd(ctx, data)
发起代理商与子客户之间转账

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**FundTransferAddRequest**](FundTransferAddRequest.md)|  | 

### Return type

[**FundTransferAddResponse**](FundTransferAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

