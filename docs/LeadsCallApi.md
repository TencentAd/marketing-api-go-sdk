# TencentAds\LeadsCallApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallCreate**](LeadsCallApi.md#LeadsCallCreate) | **Post** /leads_call/create | 网络电话呼叫


# **LeadsCallCreate**
> LeadsCallCreateResponse LeadsCallCreate(ctx, data)
网络电话呼叫

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsCallCreateRequest**](LeadsCallCreateRequest.md)|  | 

### Return type

[**LeadsCallCreateResponse**](LeadsCallCreateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

