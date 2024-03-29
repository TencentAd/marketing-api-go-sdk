# TencentAds\XijingTemplateListApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingTemplateListGet**](XijingTemplateListApi.md#XijingTemplateListGet) | **Get** /xijing_template_list/get | 获取蹊径落地页模板列表


# **XijingTemplateListGet**
> XijingTemplateListGetResponse XijingTemplateListGet(ctx, accountId, pageTemplateId, optional)
获取蹊径落地页模板列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **pageTemplateId** | **string**|  | 
 **optional** | ***XijingTemplateListApiXijingTemplateListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a XijingTemplateListApiXijingTemplateListGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isInteraction** | **optional.Bool**|  | 
 **isPublic** | **optional.Bool**|  | 
 **templateSource** | **optional.String**|  | 
 **pageSize** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**XijingTemplateListGetResponse**](XijingTemplateListGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

