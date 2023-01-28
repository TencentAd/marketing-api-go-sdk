# TencentAds\FileSchemaApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FileSchemaGet**](FileSchemaApi.md#FileSchemaGet) | **Get** /file_schema/get | 获取文件Schema定义


# **FileSchemaGet**
> FileSchemaGetResponse FileSchemaGet(ctx, accountId, scenes, optional)
获取文件Schema定义

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **scenes** | [**[]string**](string.md)|  | 
 **optional** | ***FileSchemaApiFileSchemaGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a FileSchemaApiFileSchemaGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**FileSchemaGetResponse**](FileSchemaGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

