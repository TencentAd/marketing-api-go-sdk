# TencentAds\CreativetoolsTextApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativetoolsTextGet**](CreativetoolsTextApi.md#CreativetoolsTextGet) | **Get** /creativetools_text/get | 获取广告文案


# **CreativetoolsTextGet**
> CreativetoolsTextGetResponse CreativetoolsTextGet(ctx, accountId, maxTextLength, optional)
获取广告文案

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **maxTextLength** | **int64**|  | 
 **optional** | ***CreativetoolsTextApiCreativetoolsTextGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativetoolsTextApiCreativetoolsTextGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **keyword** | **optional.String**|  | 
 **filtering** | [**optional.Interface of []int64**](int64.md)|  | 
 **number** | **optional.Int64**|  | 
 **categoryFirstLevel** | **optional.Int64**|  | 
 **categorySecondLevel** | **optional.Int64**|  | 
 **productCatalogId** | **optional.Int64**|  | 
 **productOuterIds** | [**optional.Interface of []string**](string.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CreativetoolsTextGetResponse**](CreativetoolsTextGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

