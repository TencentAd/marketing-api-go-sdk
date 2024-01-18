# TencentAds\ProgrammedApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgrammedAdd**](ProgrammedApi.md#ProgrammedAdd) | **Post** /programmed/add | 创建模板预览接口
[**ProgrammedGet**](ProgrammedApi.md#ProgrammedGet) | **Post** /programmed/get | 获取模板预览接口
[**ProgrammedUpdate**](ProgrammedApi.md#ProgrammedUpdate) | **Post** /programmed/update | 更新模板预览接口


# **ProgrammedAdd**
> ProgrammedAddResponse ProgrammedAdd(ctx, data)
创建模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedAddRequest**](ProgrammedAddRequest.md)|  | 

### Return type

[**ProgrammedAddResponse**](ProgrammedAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedGet**
> ProgrammedGetResponse ProgrammedGet(ctx, data)
获取模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedGetRequest**](ProgrammedGetRequest.md)|  | 

### Return type

[**ProgrammedGetResponse**](ProgrammedGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedUpdate**
> ProgrammedUpdateResponse ProgrammedUpdate(ctx, data)
更新模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**ProgrammedUpdateRequest**](ProgrammedUpdateRequest.md)|  | 

### Return type

[**ProgrammedUpdateResponse**](ProgrammedUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

