# TencentAds\DynamicCreativeReviewResultsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicCreativeReviewResultsGet**](DynamicCreativeReviewResultsApi.md#DynamicCreativeReviewResultsGet) | **Get** /dynamic_creative_review_results/get | 查询动态创意审核结果


# **DynamicCreativeReviewResultsGet**
> DynamicCreativeReviewResultsGetResponse DynamicCreativeReviewResultsGet(ctx, accountId, dynamicCreativeIdList, optional)
查询动态创意审核结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dynamicCreativeIdList** | [**[]int64**](int64.md)|  | 
 **optional** | ***DynamicCreativeReviewResultsApiDynamicCreativeReviewResultsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicCreativeReviewResultsApiDynamicCreativeReviewResultsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DynamicCreativeReviewResultsGetResponse**](DynamicCreativeReviewResultsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

