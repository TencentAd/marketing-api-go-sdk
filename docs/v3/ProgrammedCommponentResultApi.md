# TencentAds\ProgrammedCommponentResultApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgrammedCommponentResultGet**](ProgrammedCommponentResultApi.md#ProgrammedCommponentResultGet) | **Get** /programmed_commponent_result/get | 组件化创意衍生成品查询接口


# **ProgrammedCommponentResultGet**
> ProgrammedCommponentResultGetResponse ProgrammedCommponentResultGet(ctx, accountId, materialDeriveId, optional)
组件化创意衍生成品查询接口

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **materialDeriveId** | **int64**|  | 
 **optional** | ***ProgrammedCommponentResultApiProgrammedCommponentResultGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgrammedCommponentResultApiProgrammedCommponentResultGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ProgrammedCommponentResultGetResponse**](ProgrammedCommponentResultGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

