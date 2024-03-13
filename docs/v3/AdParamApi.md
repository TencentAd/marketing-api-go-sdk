# TencentAds\AdParamApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdParamGet**](AdParamApi.md#AdParamGet) | **Get** /ad_param/get | 获取词包


# **AdParamGet**
> AdParamGetResponse AdParamGet(ctx, accountId, marketingGoal, creativeTemplateId, siteSet, optional)
获取词包

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingGoal** | **string**|  | 
  **creativeTemplateId** | **int64**|  | 
  **siteSet** | [**[]string**](string.md)|  | 
 **optional** | ***AdParamApiAdParamGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdParamApiAdParamGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **marketingSubGoal** | **optional.String**|  | 
 **marketingCarrierType** | **optional.String**|  | 
 **marketingTargetType** | **optional.String**|  | 
 **productCatalogId** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdParamGetResponse**](AdParamGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

