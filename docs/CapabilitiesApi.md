# TencentAds\CapabilitiesApi

All URIs are relative to *https://sandbox-api.e.qq.com/v1.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CapabilitiesGet**](CapabilitiesApi.md#CapabilitiesGet) | **Get** /capabilities/get | 查询广告相关权限（待废弃）


# **CapabilitiesGet**
> CapabilitiesGetResponse CapabilitiesGet(ctx, accountId, capability, optional)
查询广告相关权限（待废弃）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **capability** | **string**|  | 
 **optional** | ***CapabilitiesApiCapabilitiesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CapabilitiesApiCapabilitiesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **querySpec** | [**optional.Interface of CapabilitiesGetQuerySpec**](CapabilitiesGetQuerySpec.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**CapabilitiesGetResponse**](CapabilitiesGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

