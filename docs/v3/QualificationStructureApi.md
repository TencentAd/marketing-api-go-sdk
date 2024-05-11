# TencentAds\QualificationStructureApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QualificationStructureGet**](QualificationStructureApi.md#QualificationStructureGet) | **Get** /qualification_structure/get | 获取广告主资质结构


# **QualificationStructureGet**
> QualificationStructureGetResponse QualificationStructureGet(ctx, accountId, qualificationCode, optional)
获取广告主资质结构

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **qualificationCode** | **string**|  | 
 **optional** | ***QualificationStructureApiQualificationStructureGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a QualificationStructureApiQualificationStructureGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**QualificationStructureGetResponse**](QualificationStructureGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

