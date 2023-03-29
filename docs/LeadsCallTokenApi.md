# TencentAds\LeadsCallTokenApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallTokenGet**](LeadsCallTokenApi.md#LeadsCallTokenGet) | **Post** /leads_call_token/get | 获取网络电话token


# **LeadsCallTokenGet**
> LeadsCallTokenGetResponse LeadsCallTokenGet(ctx, data)
获取网络电话token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsCallTokenGetRequest**](LeadsCallTokenGetRequest.md)|  | 

### Return type

[**LeadsCallTokenGetResponse**](LeadsCallTokenGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

