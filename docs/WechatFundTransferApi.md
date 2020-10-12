# TencentAds\WechatFundTransferApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WechatFundTransferAdd**](WechatFundTransferApi.md#WechatFundTransferAdd) | **Post** /wechat_fund_transfer/add | 微信服务商子客之间转账


# **WechatFundTransferAdd**
> WechatFundTransferAddResponse WechatFundTransferAdd(ctx, data)
微信服务商子客之间转账

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**WechatFundTransferAddRequest**](WechatFundTransferAddRequest.md)|  | 

### Return type

[**WechatFundTransferAddResponse**](WechatFundTransferAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

