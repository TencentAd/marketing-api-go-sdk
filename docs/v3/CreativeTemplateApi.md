# TencentAds\CreativeTemplateApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeTemplateGet**](CreativeTemplateApi.md#CreativeTemplateGet) | **Get** /creative_template/get | 获取创意形式详情


# **CreativeTemplateGet**
> CreativeTemplateGetResponse CreativeTemplateGet(ctx, accountId, marketingGoal, marketingTargetType, marketingCarrierType, deliveryMode, optional)
获取创意形式详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingGoal** | **string**|  | 
  **marketingTargetType** | **string**|  | 
  **marketingCarrierType** | **string**|  | 
  **deliveryMode** | **string**|  | 
 **optional** | ***CreativeTemplateApiCreativeTemplateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativeTemplateApiCreativeTemplateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **marketingSubGoal** | **optional.String**|  | 
 **automaticSiteEnabled** | **optional.Bool**|  | 
 **siteSet** | [**optional.Interface of []string**](string.md)|  | 
 **dynamicCreativeType** | **optional.String**|  | 
 **creativeTemplateId** | **optional.Int64**|  | 
 **useNewVersion** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CreativeTemplateGetResponse**](CreativeTemplateGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

