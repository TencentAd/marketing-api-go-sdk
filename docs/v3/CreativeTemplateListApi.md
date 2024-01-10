# TencentAds\CreativeTemplateListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreativeTemplateListGet**](CreativeTemplateListApi.md#CreativeTemplateListGet) | **Get** /creative_template_list/get | 获取创意规格列表


# **CreativeTemplateListGet**
> CreativeTemplateListGetResponse CreativeTemplateListGet(ctx, accountId, marketingGoal, marketingTargetType, marketingCarrierType, optional)
获取创意规格列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **marketingGoal** | **string**|  | 
  **marketingTargetType** | **string**|  | 
  **marketingCarrierType** | **string**|  | 
 **optional** | ***CreativeTemplateListApiCreativeTemplateListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativeTemplateListApiCreativeTemplateListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **marketingSubGoal** | **optional.String**|  | 
 **siteSet** | **optional.String**|  | 
 **dynamicAbilityType** | **optional.String**|  | 
 **wechatSceneSpecPosition** | [**optional.Interface of []int64**](int64.md)|  | 
 **creativeTemplateId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CreativeTemplateListGetResponse**](CreativeTemplateListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

