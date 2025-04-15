# TencentAds\ProgrammedCommponentPreviewApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgrammedCommponentPreviewAdd**](ProgrammedCommponentPreviewApi.md#ProgrammedCommponentPreviewAdd) | **Post** /programmed_commponent_preview/add | 组件化创意衍生预览创建接口
[**ProgrammedCommponentPreviewDelete**](ProgrammedCommponentPreviewApi.md#ProgrammedCommponentPreviewDelete) | **Post** /programmed_commponent_preview/delete | 组件化创意衍生预览删除接口
[**ProgrammedCommponentPreviewGet**](ProgrammedCommponentPreviewApi.md#ProgrammedCommponentPreviewGet) | **Get** /programmed_commponent_preview/get | 组件化创意衍生预览查询接口
[**ProgrammedCommponentPreviewUpdate**](ProgrammedCommponentPreviewApi.md#ProgrammedCommponentPreviewUpdate) | **Post** /programmed_commponent_preview/update | 组件化创意衍生预览更新接口


# **ProgrammedCommponentPreviewAdd**
> ProgrammedCommponentPreviewAddResponse ProgrammedCommponentPreviewAdd(ctx, data)
组件化创意衍生预览创建接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedCommponentPreviewAddRequest**](ProgrammedCommponentPreviewAddRequest.md)|  | 

### Return type

[**ProgrammedCommponentPreviewAddResponse**](ProgrammedCommponentPreviewAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedCommponentPreviewDelete**
> ProgrammedCommponentPreviewDeleteResponse ProgrammedCommponentPreviewDelete(ctx, data)
组件化创意衍生预览删除接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedCommponentPreviewDeleteRequest**](ProgrammedCommponentPreviewDeleteRequest.md)|  | 

### Return type

[**ProgrammedCommponentPreviewDeleteResponse**](ProgrammedCommponentPreviewDeleteResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedCommponentPreviewGet**
> ProgrammedCommponentPreviewGetResponse ProgrammedCommponentPreviewGet(ctx, accountId, materialDeriveId, optional)
组件化创意衍生预览查询接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **materialDeriveId** | **int64**|  | 
 **optional** | ***ProgrammedCommponentPreviewApiProgrammedCommponentPreviewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedCommponentPreviewApiProgrammedCommponentPreviewGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedCommponentPreviewGetResponse**](ProgrammedCommponentPreviewGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedCommponentPreviewUpdate**
> ProgrammedCommponentPreviewUpdateResponse ProgrammedCommponentPreviewUpdate(ctx, data)
组件化创意衍生预览更新接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedCommponentPreviewUpdateRequest**](ProgrammedCommponentPreviewUpdateRequest.md)|  | 

### Return type

[**ProgrammedCommponentPreviewUpdateResponse**](ProgrammedCommponentPreviewUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

