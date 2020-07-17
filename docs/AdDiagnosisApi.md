# TencentAds\AdDiagnosisApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdDiagnosisGet**](AdDiagnosisApi.md#AdDiagnosisGet) | **Post** /ad_diagnosis/get | 获取广告诊断信息


# **AdDiagnosisGet**
> AdDiagnosisGetResponse AdDiagnosisGet(ctx, data)
获取广告诊断信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**AdDiagnosisGetRequest**](AdDiagnosisGetRequest.md)|  | 

### Return type

[**AdDiagnosisGetResponse**](AdDiagnosisGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

