# TencentAds\ElementAppealReviewApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ElementAppealReviewAdd**](ElementAppealReviewApi.md#ElementAppealReviewAdd) | **Post** /element_appeal_review/add | 发起元素申诉复审
[**ElementAppealReviewGet**](ElementAppealReviewApi.md#ElementAppealReviewGet) | **Get** /element_appeal_review/get | 获取元素申诉复审结果


# **ElementAppealReviewAdd**
> ElementAppealReviewAddResponse ElementAppealReviewAdd(ctx, data)
发起元素申诉复审

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ElementAppealReviewAddRequest**](ElementAppealReviewAddRequest.md)|  | 

### Return type

[**ElementAppealReviewAddResponse**](ElementAppealReviewAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ElementAppealReviewGet**
> ElementAppealReviewGetResponse ElementAppealReviewGet(ctx, accountId, dynamicCreativeId, componentId, elementId, elementFingerPrint, optional)
获取元素申诉复审结果

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dynamicCreativeId** | **int64**|  | 
  **componentId** | **int64**|  | 
  **elementId** | **int64**|  | 
  **elementFingerPrint** | **string**|  | 
 **optional** | ***ElementAppealReviewApiElementAppealReviewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ElementAppealReviewApiElementAppealReviewGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ElementAppealReviewGetResponse**](ElementAppealReviewGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

