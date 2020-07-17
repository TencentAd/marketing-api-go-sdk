# TencentAds\AdcreativeTemplateDetailApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativeTemplateDetailGet**](AdcreativeTemplateDetailApi.md#AdcreativeTemplateDetailGet) | **Get** /adcreative_template_detail/get | 获取创意规格信息


# **AdcreativeTemplateDetailGet**
> AdcreativeTemplateDetailGetResponse AdcreativeTemplateDetailGet(ctx, adcreativeTemplateId, promotedObjectType, optional)
获取创意规格信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adcreativeTemplateId** | **int64**|  | 
  **promotedObjectType** | **string**|  | 
 **optional** | ***AdcreativeTemplateDetailApiAdcreativeTemplateDetailGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativeTemplateDetailApiAdcreativeTemplateDetailGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **accountId** | **optional.Int64**|  | 
 **automaticSiteEnabled** | **optional.Bool**|  | 
 **siteSet** | [**optional.Interface of []string**](string.md)|  | 
 **isDynamicCreativeAd** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativeTemplateDetailGetResponse**](AdcreativeTemplateDetailGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

