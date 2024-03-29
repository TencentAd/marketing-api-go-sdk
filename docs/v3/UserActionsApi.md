# TencentAds\UserActionsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserActionsAdd**](UserActionsApi.md#UserActionsAdd) | **Post** /user_actions/add | 上传用户行为数据


# **UserActionsAdd**
> UserActionsAddResponse UserActionsAdd(ctx, data)
上传用户行为数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**UserActionsAddRequest**](UserActionsAddRequest.md)|  | 

### Return type

[**UserActionsAddResponse**](UserActionsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

