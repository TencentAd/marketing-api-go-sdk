# TencentAds\UserPropertiesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserPropertiesAdd**](UserPropertiesApi.md#UserPropertiesAdd) | **Post** /user_properties/add | 上传用户属性数据


# **UserPropertiesAdd**
> UserPropertiesAddResponse UserPropertiesAdd(ctx, data)
上传用户属性数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserPropertiesAddRequest**](UserPropertiesAddRequest.md)|  | 

### Return type

[**UserPropertiesAddResponse**](UserPropertiesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

