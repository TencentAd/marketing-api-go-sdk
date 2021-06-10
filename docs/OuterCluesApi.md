# TencentAds\OuterCluesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OuterCluesAdd**](OuterCluesApi.md#OuterCluesAdd) | **Post** /outer_clues/add | 第三方线索导入
[**OuterCluesUpdate**](OuterCluesApi.md#OuterCluesUpdate) | **Post** /outer_clues/update | 更新线索状态


# **OuterCluesAdd**
> OuterCluesAddResponse OuterCluesAdd(ctx, data)
第三方线索导入

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesAddRequest**](OuterCluesAddRequest.md)|  | 

### Return type

[**OuterCluesAddResponse**](OuterCluesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OuterCluesUpdate**
> OuterCluesUpdateResponse OuterCluesUpdate(ctx, data)
更新线索状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesUpdateRequest**](OuterCluesUpdateRequest.md)|  | 

### Return type

[**OuterCluesUpdateResponse**](OuterCluesUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

