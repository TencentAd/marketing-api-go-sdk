# TencentAds\PropertyFilesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PropertyFilesAdd**](PropertyFilesApi.md#PropertyFilesAdd) | **Post** /property_files/add | 上传属性数据文件


# **PropertyFilesAdd**
> PropertyFilesAddResponse PropertyFilesAdd(ctx, accountId, propertySetId, sessionId, fileName, file)
上传属性数据文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **propertySetId** | **int64**|  | 
  **sessionId** | **int64**|  | 
  **fileName** | **string**|  | 
  **file** | ***os.File**|  | 

### Return type

[**PropertyFilesAddResponse**](PropertyFilesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

