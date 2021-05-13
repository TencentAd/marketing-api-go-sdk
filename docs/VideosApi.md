# TencentAds\VideosApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideosAdd**](VideosApi.md#VideosAdd) | **Post** /videos/add | 添加视频文件
[**VideosDelete**](VideosApi.md#VideosDelete) | **Post** /videos/delete | 删除视频
[**VideosGet**](VideosApi.md#VideosGet) | **Get** /videos/get | 获取视频文件
[**VideosUpdate**](VideosApi.md#VideosUpdate) | **Post** /videos/update | 修改视频信息


# **VideosAdd**
> VideosAddResponse VideosAdd(ctx, accountId, videoFile, signature, optional)
添加视频文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **videoFile** | ***os.File**|  | 
  **signature** | **string**|  | 
 **optional** | ***VideosApiVideosAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideosApiVideosAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **description** | **optional.String**|  | 
 **adcreativeTemplateId** | **optional.Int64**|  | 

### Return type

[**VideosAddResponse**](VideosAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideosDelete**
> VideosDeleteResponse VideosDelete(ctx, data)
删除视频

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**VideosDeleteRequest**](VideosDeleteRequest.md)|  | 

### Return type

[**VideosDeleteResponse**](VideosDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideosGet**
> VideosGetResponse VideosGet(ctx, accountId, optional)
获取视频文件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***VideosApiVideosGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VideosApiVideosGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filtering** | [**optional.Interface of []FilteringStruct**](FilteringStruct.md)|  | 
 **page** | **optional.Int64**|  | 
 **pageSize** | **optional.Int64**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**VideosGetResponse**](VideosGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VideosUpdate**
> VideosUpdateResponse VideosUpdate(ctx, data)
修改视频信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**VideosUpdateRequest**](VideosUpdateRequest.md)|  | 

### Return type

[**VideosUpdateResponse**](VideosUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

