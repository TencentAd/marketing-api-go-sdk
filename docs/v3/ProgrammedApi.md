# TencentAds\ProgrammedApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgrammedAdd**](ProgrammedApi.md#ProgrammedAdd) | **Get** /programmed/add | 创建模板预览接口
[**ProgrammedGet**](ProgrammedApi.md#ProgrammedGet) | **Get** /programmed/get | 获取模板预览接口
[**ProgrammedUpdate**](ProgrammedApi.md#ProgrammedUpdate) | **Get** /programmed/update | 更新模板预览接口


# **ProgrammedAdd**
> ProgrammedAddResponse ProgrammedAdd(ctx, accountId, adgroupId, createMaterialGroups, optional)
创建模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adgroupId** | **int64**|  | 
  **createMaterialGroups** | [**[]MaterialGroupCreateStruct**](MaterialGroupCreateStruct.md)|  | 
 **optional** | ***ProgrammedApiProgrammedAddOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedApiProgrammedAddOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **autoDerivedProgramCreativeSwitch** | **optional.Bool**|  | 
 **standardSwitch** | **optional.Bool**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedAddResponse**](ProgrammedAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedGet**
> ProgrammedGetResponse ProgrammedGet(ctx, accountId, materialDeriveId, optional)
获取模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **materialDeriveId** | **int64**|  | 
 **optional** | ***ProgrammedApiProgrammedGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedApiProgrammedGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedGetResponse**](ProgrammedGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProgrammedUpdate**
> ProgrammedUpdateResponse ProgrammedUpdate(ctx, accountId, materialDeriveId, optional)
更新模板预览接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **materialDeriveId** | **int64**|  | 
 **optional** | ***ProgrammedApiProgrammedUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedApiProgrammedUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **autoDerivedProgramCreativeSwitch** | **optional.Bool**|  | 
 **standardSwitch** | **optional.Bool**|  | 
 **updateMaterialGroups** | [**optional.Interface of []MaterialGroupUpdateStruct**](MaterialGroupUpdateStruct.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedUpdateResponse**](ProgrammedUpdateResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

