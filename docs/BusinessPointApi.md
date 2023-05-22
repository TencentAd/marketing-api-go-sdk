# TencentAds\BusinessPointApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BusinessPointGet**](BusinessPointApi.md#BusinessPointGet) | **Post** /business_point/get | 查询行业业务点信息


# **BusinessPointGet**
> BusinessPointGetResponse BusinessPointGet(ctx, data)
查询行业业务点信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**BusinessPointGetRequest**](BusinessPointGetRequest.md)|  | 

### Return type

[**BusinessPointGetResponse**](BusinessPointGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

