# TencentAds\DynamicAdTemplatesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicAdTemplatesGet**](DynamicAdTemplatesApi.md#DynamicAdTemplatesGet) | **Get** /dynamic_ad_templates/get | 获取动态模板信息


# **DynamicAdTemplatesGet**
> DynamicAdTemplatesGetResponse DynamicAdTemplatesGet(ctx, accountId, productCatalogId, dynamicAdTemplateWidth, dynamicAdTemplateHeight, optional)
获取动态模板信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **productCatalogId** | **int64**|  | 
  **dynamicAdTemplateWidth** | **int64**|  | 
  **dynamicAdTemplateHeight** | **int64**|  | 
 **optional** | ***DynamicAdTemplatesApiDynamicAdTemplatesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicAdTemplatesApiDynamicAdTemplatesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **dynamicAdTemplateType** | **optional.String**|  | 
 **dynamicAdTemplateOwnershipType** | **optional.String**|  | 
 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DynamicAdTemplatesGetResponse**](DynamicAdTemplatesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

