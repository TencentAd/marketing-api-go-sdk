# TencentAds\TargetingTagsUvApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetingTagsUvGet**](TargetingTagsUvApi.md#TargetingTagsUvGet) | **Get** /targeting_tags_uv/get | 获取行为/兴趣/意向标签覆盖人群数


# **TargetingTagsUvGet**
> TargetingTagsUvGetResponse TargetingTagsUvGet(ctx, accountId, categoryType, optional)
获取行为/兴趣/意向标签覆盖人群数

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **categoryType** | **string**|  | 
 **optional** | ***TargetingTagsUvApiTargetingTagsUvGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingTagsUvApiTargetingTagsUvGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **categoryList** | [**optional.Interface of []int64**](int64.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**TargetingTagsUvGetResponse**](TargetingTagsUvGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

