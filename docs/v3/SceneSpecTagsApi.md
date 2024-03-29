# TencentAds\SceneSpecTagsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SceneSpecTagsGet**](SceneSpecTagsApi.md#SceneSpecTagsGet) | **Get** /scene_spec_tags/get | 获取场景定向标签


# **SceneSpecTagsGet**
> SceneSpecTagsGetResponse SceneSpecTagsGet(ctx, accountId, type_, optional)
获取场景定向标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **type_** | **string**|  | 
 **optional** | ***SceneSpecTagsApiSceneSpecTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SceneSpecTagsApiSceneSpecTagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**SceneSpecTagsGetResponse**](SceneSpecTagsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

