# TencentAds\SaasMiniprogramPrereviewApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SaasMiniprogramPrereviewAdd**](SaasMiniprogramPrereviewApi.md#SaasMiniprogramPrereviewAdd) | **Post** /saas_miniprogram_prereview/add | SaaS小程序送审
[**SaasMiniprogramPrereviewGet**](SaasMiniprogramPrereviewApi.md#SaasMiniprogramPrereviewGet) | **Get** /saas_miniprogram_prereview/get | SaaS小程序送审结果获取


# **SaasMiniprogramPrereviewAdd**
> SaasMiniprogramPrereviewAddResponse SaasMiniprogramPrereviewAdd(ctx, data)
SaaS小程序送审

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**SaasMiniprogramPrereviewAddRequest**](SaasMiniprogramPrereviewAddRequest.md)|  | 

### Return type

[**SaasMiniprogramPrereviewAddResponse**](SaasMiniprogramPrereviewAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SaasMiniprogramPrereviewGet**
> SaasMiniprogramPrereviewGetResponse SaasMiniprogramPrereviewGet(ctx, taskId, optional)
SaaS小程序送审结果获取

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **taskId** | **string**|  | 
 **optional** | ***SaasMiniprogramPrereviewApiSaasMiniprogramPrereviewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SaasMiniprogramPrereviewApiSaasMiniprogramPrereviewGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**SaasMiniprogramPrereviewGetResponse**](SaasMiniprogramPrereviewGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

