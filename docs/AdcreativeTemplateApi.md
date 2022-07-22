# TencentAds\AdcreativeTemplateApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativeTemplateGet**](AdcreativeTemplateApi.md#AdcreativeTemplateGet) | **Get** /adcreative_template/get | 获取创意规格详情


# **AdcreativeTemplateGet**
> AdcreativeTemplateGetResponse AdcreativeTemplateGet(ctx, accountId, promotedObjectType, optional)
获取创意规格详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **promotedObjectType** | **string**|  | 
 **optional** | ***AdcreativeTemplateApiAdcreativeTemplateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativeTemplateApiAdcreativeTemplateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **siteSet** | [**optional.Interface of []string**](string.md)|  | 
 **automaticSiteEnabled** | **optional.Bool**|  | 
 **isDynamicCreative** | **optional.Bool**|  | 
 **adcreativeTemplateId** | **optional.Int64**|  | 
 **dynamicCreativeType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativeTemplateGetResponse**](AdcreativeTemplateGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

