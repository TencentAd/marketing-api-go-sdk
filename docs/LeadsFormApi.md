# TencentAds\LeadsFormApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsFormAdd**](LeadsFormApi.md#LeadsFormAdd) | **Post** /leads_form/add | 创建表单组件
[**LeadsFormGet**](LeadsFormApi.md#LeadsFormGet) | **Get** /leads_form/get | 获取表单组件详情


# **LeadsFormAdd**
> LeadsFormAddResponse LeadsFormAdd(ctx, data)
创建表单组件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **data** | [**LeadsFormAddRequest**](LeadsFormAddRequest.md)|  | 

### Return type

[**LeadsFormAddResponse**](LeadsFormAddResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: application/json, application/xml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LeadsFormGet**
> LeadsFormGetResponse LeadsFormGet(ctx, accountId, componentId, optional)
获取表单组件详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **componentId** | **string**|  | 
 **optional** | ***LeadsFormApiLeadsFormGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeadsFormApiLeadsFormGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LeadsFormGetResponse**](LeadsFormGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

