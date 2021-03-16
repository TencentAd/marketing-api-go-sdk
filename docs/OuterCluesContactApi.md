# TencentAds\OuterCluesContactApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OuterCluesContactUpdate**](OuterCluesContactApi.md#OuterCluesContactUpdate) | **Post** /outer_clues_contact/update | 更新线索基本信息


# **OuterCluesContactUpdate**
> OuterCluesContactUpdateResponse OuterCluesContactUpdate(ctx, data)
更新线索基本信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**OuterCluesContactUpdateRequest**](OuterCluesContactUpdateRequest.md)|  | 

### Return type

[**OuterCluesContactUpdateResponse**](OuterCluesContactUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

