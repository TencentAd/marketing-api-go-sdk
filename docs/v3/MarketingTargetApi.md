# TencentAds\MarketingTargetApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketingTargetAdd**](MarketingTargetApi.md#MarketingTargetAdd) | **Post** /marketing_target/add | 营销对象新增
[**MarketingTargetEducationSearch**](MarketingTargetApi.md#MarketingTargetEducationSearch) | **Post** /marketing_target/education_search | 教育营销对象搜索


# **MarketingTargetAdd**
> MarketingTargetAddResponse MarketingTargetAdd(ctx, data)
营销对象新增

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MarketingTargetAddRequest**](MarketingTargetAddRequest.md)|  | 

### Return type

[**MarketingTargetAddResponse**](MarketingTargetAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarketingTargetEducationSearch**
> MarketingTargetEducationSearchResponse MarketingTargetEducationSearch(ctx, data)
教育营销对象搜索

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**MarketingTargetEducationSearchRequest**](MarketingTargetEducationSearchRequest.md)|  | 

### Return type

[**MarketingTargetEducationSearchResponse**](MarketingTargetEducationSearchResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

