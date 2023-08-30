# TencentAds\CustomAudienceFilesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CustomAudienceFilesAdd**](CustomAudienceFilesApi.md#CustomAudienceFilesAdd) | **Post** /custom_audience_files/add | 上传客户人群数据文件
[**CustomAudienceFilesGet**](CustomAudienceFilesApi.md#CustomAudienceFilesGet) | **Get** /custom_audience_files/get | 获取客户人群数据文件


# **CustomAudienceFilesAdd**
> CustomAudienceFilesAddResponse CustomAudienceFilesAdd(ctx, accountId, audienceId, userIdType, file, optional)
上传客户人群数据文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **audienceId** | **int64**|  | 
  **userIdType** | **string**|  | 
  **file** | ***os.File**|  | 
 **optional** | ***CustomAudienceFilesApiCustomAudienceFilesAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudienceFilesApiCustomAudienceFilesAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **operationType** | **optional.String**|  | 
 **openAppId** | **optional.String**|  | 

### Return type

[**CustomAudienceFilesAddResponse**](CustomAudienceFilesAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CustomAudienceFilesGet**
> CustomAudienceFilesGetResponse CustomAudienceFilesGet(ctx, accountId, optional)
获取客户人群数据文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***CustomAudienceFilesApiCustomAudienceFilesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomAudienceFilesApiCustomAudienceFilesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **audienceId** | **optional.Int64**|  | 
 **customAudienceFileId** | **optional.Int64**|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CustomAudienceFilesGetResponse**](CustomAudienceFilesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

