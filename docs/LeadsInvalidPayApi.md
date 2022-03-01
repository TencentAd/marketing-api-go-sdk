# TencentAds\LeadsInvalidPayApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsInvalidPayGet**](LeadsInvalidPayApi.md#LeadsInvalidPayGet) | **Post** /leads_invalid_pay/get | 获取无效赔付明细


# **LeadsInvalidPayGet**
> LeadsInvalidPayGetResponse LeadsInvalidPayGet(ctx, data)
获取无效赔付明细

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsInvalidPayGetRequest**](LeadsInvalidPayGetRequest.md)|  | 

### Return type

[**LeadsInvalidPayGetResponse**](LeadsInvalidPayGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

