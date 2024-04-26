# TencentAds\ProgrammedMaterialMappingsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgrammedMaterialMappingsGet**](ProgrammedMaterialMappingsApi.md#ProgrammedMaterialMappingsGet) | **Get** /programmed_material_mappings/get | 获取衍生素材映射关系接口


# **ProgrammedMaterialMappingsGet**
> ProgrammedMaterialMappingsGetResponse ProgrammedMaterialMappingsGet(ctx, accountId, dynamicCreativeId, optional)
获取衍生素材映射关系接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **dynamicCreativeId** | **int64**|  | 
 **optional** | ***ProgrammedMaterialMappingsApiProgrammedMaterialMappingsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedMaterialMappingsApiProgrammedMaterialMappingsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedMaterialMappingsGetResponse**](ProgrammedMaterialMappingsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

