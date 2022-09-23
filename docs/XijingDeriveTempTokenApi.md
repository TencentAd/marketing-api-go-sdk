# TencentAds\XijingDeriveTempTokenApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingDeriveTempTokenGet**](XijingDeriveTempTokenApi.md#XijingDeriveTempTokenGet) | **Post** /xijing_derive_temp_token/get | 生成预览token


# **XijingDeriveTempTokenGet**
> XijingDeriveTempTokenGetResponse XijingDeriveTempTokenGet(ctx, data)
生成预览token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**XijingDeriveTempTokenGetRequest**](XijingDeriveTempTokenGetRequest.md)|  | 

### Return type

[**XijingDeriveTempTokenGetResponse**](XijingDeriveTempTokenGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

