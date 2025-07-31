# TencentAds\CreativeRecommendApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeRecommendGet**](CreativeRecommendApi.md#CreativeRecommendGet) | **Post** /creative_recommend/get | 获取创意相关建议


# **CreativeRecommendGet**
> CreativeRecommendGetResponse CreativeRecommendGet(ctx, data)
获取创意相关建议

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**CreativeRecommendGetRequest**](CreativeRecommendGetRequest.md)|  | 

### Return type

[**CreativeRecommendGetResponse**](CreativeRecommendGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

