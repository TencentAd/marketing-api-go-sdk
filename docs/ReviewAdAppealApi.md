# TencentAds\ReviewAdAppealApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReviewAdAppealAdd**](ReviewAdAppealApi.md#ReviewAdAppealAdd) | **Post** /review_ad_appeal/add | 发起广告申诉复审
[**ReviewAdAppealGet**](ReviewAdAppealApi.md#ReviewAdAppealGet) | **Post** /review_ad_appeal/get | 获取广告申诉复审结果


# **ReviewAdAppealAdd**
> ReviewAdAppealAddResponse ReviewAdAppealAdd(ctx, data)
发起广告申诉复审

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReviewAdAppealAddRequest**](ReviewAdAppealAddRequest.md)|  | 

### Return type

[**ReviewAdAppealAddResponse**](ReviewAdAppealAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReviewAdAppealGet**
> ReviewAdAppealGetResponse ReviewAdAppealGet(ctx, data)
获取广告申诉复审结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ReviewAdAppealGetRequest**](ReviewAdAppealGetRequest.md)|  | 

### Return type

[**ReviewAdAppealGetResponse**](ReviewAdAppealGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

