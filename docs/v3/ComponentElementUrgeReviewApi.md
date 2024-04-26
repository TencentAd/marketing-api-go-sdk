# TencentAds\ComponentElementUrgeReviewApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComponentElementUrgeReviewAdd**](ComponentElementUrgeReviewApi.md#ComponentElementUrgeReviewAdd) | **Post** /component_element_urge_review/add | 创意组件元素催审
[**ComponentElementUrgeReviewGet**](ComponentElementUrgeReviewApi.md#ComponentElementUrgeReviewGet) | **Get** /component_element_urge_review/get | 获取创意组件元素催审状态


# **ComponentElementUrgeReviewAdd**
> ComponentElementUrgeReviewAddResponse ComponentElementUrgeReviewAdd(ctx, data)
创意组件元素催审

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ComponentElementUrgeReviewAddRequest**](ComponentElementUrgeReviewAddRequest.md)|  | 

### Return type

[**ComponentElementUrgeReviewAddResponse**](ComponentElementUrgeReviewAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ComponentElementUrgeReviewGet**
> ComponentElementUrgeReviewGetResponse ComponentElementUrgeReviewGet(ctx, accountId, dynamicCreativeId, optional)
获取创意组件元素催审状态

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dynamicCreativeId** | **int64**|  | 
 **optional** | ***ComponentElementUrgeReviewApiComponentElementUrgeReviewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ComponentElementUrgeReviewApiComponentElementUrgeReviewGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **componentIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **elementFingerprintList** | [**optional.Interface of []string**](string.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ComponentElementUrgeReviewGetResponse**](ComponentElementUrgeReviewGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

