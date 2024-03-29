# TencentAds\LeadsCallVirtualNumberApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsCallVirtualNumberGet**](LeadsCallVirtualNumberApi.md#LeadsCallVirtualNumberGet) | **Get** /leads_call_virtual_number/get | 获取中间号


# **LeadsCallVirtualNumberGet**
> LeadsCallVirtualNumberGetResponse LeadsCallVirtualNumberGet(ctx, accountId, caller, callee, optional)
获取中间号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **caller** | **string**|  | 
  **callee** | **string**|  | 
 **optional** | ***LeadsCallVirtualNumberApiLeadsCallVirtualNumberGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeadsCallVirtualNumberApiLeadsCallVirtualNumberGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **leadsId** | **optional.Int64**|  | 
 **outerLeadsId** | **optional.String**|  | 
 **requestId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LeadsCallVirtualNumberGetResponse**](LeadsCallVirtualNumberGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

