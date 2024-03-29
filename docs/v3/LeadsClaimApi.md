# TencentAds\LeadsClaimApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsClaimUpdate**](LeadsClaimApi.md#LeadsClaimUpdate) | **Post** /leads_claim/update | 更新线索归因信息


# **LeadsClaimUpdate**
> LeadsClaimUpdateResponse LeadsClaimUpdate(ctx, data)
更新线索归因信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsClaimUpdateRequest**](LeadsClaimUpdateRequest.md)|  | 

### Return type

[**LeadsClaimUpdateResponse**](LeadsClaimUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

