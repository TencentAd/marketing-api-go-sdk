# TencentAds\CustomDataSaltApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomDataSaltGet**](CustomDataSaltApi.md#CustomDataSaltGet) | **Post** /custom_data_salt/get | 获取盐列表


# **CustomDataSaltGet**
> CustomDataSaltGetResponse CustomDataSaltGet(ctx, data)
获取盐列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CustomDataSaltGetRequest**](CustomDataSaltGetRequest.md)|  | 

### Return type

[**CustomDataSaltGetResponse**](CustomDataSaltGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

