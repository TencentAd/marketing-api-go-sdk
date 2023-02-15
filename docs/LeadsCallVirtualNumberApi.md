# TencentAds\LeadsCallVirtualNumberApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallVirtualNumberGet**](LeadsCallVirtualNumberApi.md#LeadsCallVirtualNumberGet) | **Post** /leads_call_virtual_number/get | 获取中间号


# **LeadsCallVirtualNumberGet**
> LeadsCallVirtualNumberGetResponse LeadsCallVirtualNumberGet(ctx, data)
获取中间号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsCallVirtualNumberGetRequest**](LeadsCallVirtualNumberGetRequest.md)|  | 

### Return type

[**LeadsCallVirtualNumberGetResponse**](LeadsCallVirtualNumberGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

