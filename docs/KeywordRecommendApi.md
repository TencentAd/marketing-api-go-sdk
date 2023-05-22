# TencentAds\KeywordRecommendApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KeywordRecommendGet**](KeywordRecommendApi.md#KeywordRecommendGet) | **Post** /keyword_recommend/get | 获取关键词推荐结果


# **KeywordRecommendGet**
> KeywordRecommendGetResponse KeywordRecommendGet(ctx, data)
获取关键词推荐结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**KeywordRecommendGetRequest**](KeywordRecommendGetRequest.md)|  | 

### Return type

[**KeywordRecommendGetResponse**](KeywordRecommendGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

