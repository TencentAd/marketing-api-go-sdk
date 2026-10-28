# TencentAds\CreativeReviewApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeReviewAdd**](CreativeReviewApi.md#CreativeReviewAdd) | **Post** /creative_review/add | 发起创意审核


# **CreativeReviewAdd**
> CreativeReviewAddResponse CreativeReviewAdd(ctx, data)
发起创意审核

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeReviewAddRequest**](CreativeReviewAddRequest.md)|  | 

### Return type

[**CreativeReviewAddResponse**](CreativeReviewAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

