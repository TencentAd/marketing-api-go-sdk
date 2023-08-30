# TencentAds\TargetcpaTagApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TargetcpaTagGet**](TargetcpaTagApi.md#TargetcpaTagGet) | **Get** /targetcpa_tag/get | 获取分人群出价标签


# **TargetcpaTagGet**
> TargetcpaTagGetResponse TargetcpaTagGet(ctx, accountId, tagTypes, optional)
获取分人群出价标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **tagTypes** | [**[]string**](string.md)|  | 
 **optional** | ***TargetcpaTagApiTargetcpaTagGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetcpaTagApiTargetcpaTagGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**TargetcpaTagGetResponse**](TargetcpaTagGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

