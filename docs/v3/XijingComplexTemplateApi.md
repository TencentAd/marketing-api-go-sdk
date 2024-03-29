# TencentAds\XijingComplexTemplateApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**XijingComplexTemplateGet**](XijingComplexTemplateApi.md#XijingComplexTemplateGet) | **Get** /xijing_complex_template/get | 获取蹊径落地页互动模板配置


# **XijingComplexTemplateGet**
> XijingComplexTemplateGetResponse XijingComplexTemplateGet(ctx, accountId, pageTemplateId, optional)
获取蹊径落地页互动模板配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **pageTemplateId** | **string**|  | 
 **optional** | ***XijingComplexTemplateApiXijingComplexTemplateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a XijingComplexTemplateApiXijingComplexTemplateGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**XijingComplexTemplateGetResponse**](XijingComplexTemplateGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

