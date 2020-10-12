# TencentAds\LocalEndadsmanuallyApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocalEndadsmanuallyAdd**](LocalEndadsmanuallyApi.md#LocalEndadsmanuallyAdd) | **Post** /local_endadsmanually/add | 手动结束广告


# **LocalEndadsmanuallyAdd**
> LocalEndadsmanuallyAddResponse LocalEndadsmanuallyAdd(ctx, data)
手动结束广告

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LocalEndadsmanuallyAddRequest**](LocalEndadsmanuallyAddRequest.md)|  | 

### Return type

[**LocalEndadsmanuallyAddResponse**](LocalEndadsmanuallyAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

