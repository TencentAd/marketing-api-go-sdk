# TencentAds\ComponentReviewResultsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentReviewResultsGet**](ComponentReviewResultsApi.md#ComponentReviewResultsGet) | **Get** /component_review_results/get | 查询组件审核结果


# **ComponentReviewResultsGet**
> ComponentReviewResultsGetResponse ComponentReviewResultsGet(ctx, accountId, componentIdList, optional)
查询组件审核结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **componentIdList** | [**[]int64**](int64.md)|  | 
 **optional** | ***ComponentReviewResultsApiComponentReviewResultsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentReviewResultsApiComponentReviewResultsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ComponentReviewResultsGetResponse**](ComponentReviewResultsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

