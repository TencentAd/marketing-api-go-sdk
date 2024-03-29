# TencentAds\DynamicAdImageTemplatesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicAdImageTemplatesGet**](DynamicAdImageTemplatesApi.md#DynamicAdImageTemplatesGet) | **Get** /dynamic_ad_image_templates/get | 获取动态商品图片模板


# **DynamicAdImageTemplatesGet**
> DynamicAdImageTemplatesGetResponse DynamicAdImageTemplatesGet(ctx, accountId, productCatalogId, productMode, dynamicAdTemplateWidth, dynamicAdTemplateHeight, optional)
获取动态商品图片模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
  **productMode** | **string**|  | 
  **dynamicAdTemplateWidth** | **int64**|  | 
  **dynamicAdTemplateHeight** | **int64**|  | 
 **optional** | ***DynamicAdImageTemplatesApiDynamicAdImageTemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicAdImageTemplatesApiDynamicAdImageTemplatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **dynamicAdTemplateOwnershipType** | **optional.String**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **templateIdList** | [**optional.Interface of []int64**](int64.md)|  | 
 **templateName** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DynamicAdImageTemplatesGetResponse**](DynamicAdImageTemplatesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

