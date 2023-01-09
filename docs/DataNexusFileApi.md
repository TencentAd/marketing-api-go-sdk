# TencentAds\DataNexusFileApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataNexusFileAdd**](DataNexusFileApi.md#DataNexusFileAdd) | **Post** /data_nexus_file/add | 上传文件
[**DataNexusFileGet**](DataNexusFileApi.md#DataNexusFileGet) | **Get** /data_nexus_file/get | 文件查看


# **DataNexusFileAdd**
> DataNexusFileAddResponse DataNexusFileAdd(ctx, accountId, fileName, fileDesc, file, schemaDefine, scenes, optional)
上传文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **fileName** | **string**|  | 
  **fileDesc** | **string**|  | 
  **file** | ***os.File**|  | 
  **schemaDefine** | [**[]SchemeCol**](SchemeCol.md)|  | 
  **scenes** | [**[]SelectScene**](SelectScene.md)|  | 
 **optional** | ***DataNexusFileApiDataNexusFileAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataNexusFileApiDataNexusFileAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **extraInfo** | [**optional.Interface of FileExtraInfo**](FileExtraInfo.md)|  | 

### Return type

[**DataNexusFileAddResponse**](DataNexusFileAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DataNexusFileGet**
> DataNexusFileGetResponse DataNexusFileGet(ctx, accountId, optional)
文件查看

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***DataNexusFileApiDataNexusFileGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DataNexusFileApiDataNexusFileGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fileId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DataNexusFileGetResponse**](DataNexusFileGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

