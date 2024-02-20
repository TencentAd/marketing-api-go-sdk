# TencentAds\AdcreativePreviewsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativePreviewsAdd**](AdcreativePreviewsApi.md#AdcreativePreviewsAdd) | **Post** /adcreative_previews/add | 绑定广告预览受众
[**AdcreativePreviewsGet**](AdcreativePreviewsApi.md#AdcreativePreviewsGet) | **Get** /adcreative_previews/get | 获取绑定的广告预览受众列表


# **AdcreativePreviewsAdd**
> AdcreativePreviewsAddResponse AdcreativePreviewsAdd(ctx, data)
绑定广告预览受众

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdcreativePreviewsAddRequest**](AdcreativePreviewsAddRequest.md)|  | 

### Return type

[**AdcreativePreviewsAddResponse**](AdcreativePreviewsAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdcreativePreviewsGet**
> AdcreativePreviewsGetResponse AdcreativePreviewsGet(ctx, accountId, filtering, optional)
获取绑定的广告预览受众列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **filtering** | [**[]FilteringStruct**](FilteringStruct.md)|  | 
 **optional** | ***AdcreativePreviewsApiAdcreativePreviewsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativePreviewsApiAdcreativePreviewsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativePreviewsGetResponse**](AdcreativePreviewsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

