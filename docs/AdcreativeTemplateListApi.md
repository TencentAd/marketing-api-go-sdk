# TencentAds\AdcreativeTemplateListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdcreativeTemplateListGet**](AdcreativeTemplateListApi.md#AdcreativeTemplateListGet) | **Get** /adcreative_template_list/get | 获取创意规格列表


# **AdcreativeTemplateListGet**
> AdcreativeTemplateListGetResponse AdcreativeTemplateListGet(ctx, accountId, optional)
获取创意规格列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***AdcreativeTemplateListApiAdcreativeTemplateListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdcreativeTemplateListApiAdcreativeTemplateListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **siteSet** | **optional.String**|  | 
 **campaignType** | **optional.String**|  | 
 **promotedObjectType** | **optional.String**|  | 
 **dynamicAbilityType** | [**optional.Interface of []string**](string.md)|  | 
 **isDynamicCreative** | **optional.Bool**|  | 
 **wechatSceneSpecPosition** | [**optional.Interface of []int64**](int64.md)|  | 
 **adcreativeTemplateId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**AdcreativeTemplateListGetResponse**](AdcreativeTemplateListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

