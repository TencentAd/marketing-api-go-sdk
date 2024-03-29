# TencentAds\XijingTemplateApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingTemplateGet**](XijingTemplateApi.md#XijingTemplateGet) | **Get** /xijing_template/get | 获取蹊径落地页模板


# **XijingTemplateGet**
> XijingTemplateGetResponse XijingTemplateGet(ctx, accountId, templateId, optional)
获取蹊径落地页模板

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **templateId** | **string**|  | 
 **optional** | ***XijingTemplateApiXijingTemplateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a XijingTemplateApiXijingTemplateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**XijingTemplateGetResponse**](XijingTemplateGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

