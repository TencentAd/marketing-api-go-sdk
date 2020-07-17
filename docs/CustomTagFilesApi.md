# TencentAds\CustomTagFilesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomTagFilesAdd**](CustomTagFilesApi.md#CustomTagFilesAdd) | **Post** /custom_tag_files/add | 上传客户标签人群数据文件
[**CustomTagFilesGet**](CustomTagFilesApi.md#CustomTagFilesGet) | **Get** /custom_tag_files/get | 获取客户标签人群文件


# **CustomTagFilesAdd**
> CustomTagFilesAddResponse CustomTagFilesAdd(ctx, accountId, userIdType, tagId, file, optional)
上传客户标签人群数据文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **userIdType** | **string**|  | 
  **tagId** | **int64**|  | 
  **file** | ***os.File**|  | 
 **optional** | ***CustomTagFilesApiCustomTagFilesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomTagFilesApiCustomTagFilesAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **operationType** | **optional.String**|  | 
 **openAppId** | **optional.String**|  | 

### Return type

[**CustomTagFilesAddResponse**](CustomTagFilesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomTagFilesGet**
> CustomTagFilesGetResponse CustomTagFilesGet(ctx, accountId, optional)
获取客户标签人群文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CustomTagFilesApiCustomTagFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomTagFilesApiCustomTagFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomTagFilesGetResponse**](CustomTagFilesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

