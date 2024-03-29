# TencentAds\DynamicAdVideoTemplatesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicAdVideoTemplatesGet**](DynamicAdVideoTemplatesApi.md#DynamicAdVideoTemplatesGet) | **Get** /dynamic_ad_video_templates/get | 获取动态商品视频模板


# **DynamicAdVideoTemplatesGet**
> DynamicAdVideoTemplatesGetResponse DynamicAdVideoTemplatesGet(ctx, accountId, productCatalogId, adcreativeTemplateId, productMode, optional)
获取动态商品视频模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
  **adcreativeTemplateId** | **int64**|  | 
  **productMode** | **string**|  | 
 **optional** | ***DynamicAdVideoTemplatesApiDynamicAdVideoTemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicAdVideoTemplatesApiDynamicAdVideoTemplatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **supportChannel** | **optional.Bool**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **templateIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **templateName** | **optional.String**|  | 
 **dynamicAdTemplateOwnershipType** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DynamicAdVideoTemplatesGetResponse**](DynamicAdVideoTemplatesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

