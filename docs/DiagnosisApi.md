# TencentAds\DiagnosisApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiagnosisGet**](DiagnosisApi.md#DiagnosisGet) | **Get** /diagnosis/get | 广告诊断分析工具


# **DiagnosisGet**
> DiagnosisGetResponse DiagnosisGet(ctx, accountId, adgroupIdList, optional)
广告诊断分析工具

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **adgroupIdList** | [**[]int64**](int64.md)|  | 
 **optional** | ***DiagnosisApiDiagnosisGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiagnosisApiDiagnosisGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **timeRange** | [**optional.Interface of TimeRange**](TimeRange.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**DiagnosisGetResponse**](DiagnosisGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

