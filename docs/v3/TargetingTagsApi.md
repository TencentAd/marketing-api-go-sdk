# TencentAds\TargetingTagsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetingTagsGet**](TargetingTagsApi.md#TargetingTagsGet) | **Get** /targeting_tags/get | 获取定向标签


# **TargetingTagsGet**
> TargetingTagsGetResponse TargetingTagsGet(ctx, accountId, type_, optional)
获取定向标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **type_** | **string**|  | 
 **optional** | ***TargetingTagsApiTargetingTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingTagsApiTargetingTagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagSpec** | [**optional.Interface of TagSpec**](TagSpec.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**TargetingTagsGetResponse**](TargetingTagsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

