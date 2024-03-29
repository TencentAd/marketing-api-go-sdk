# TencentAds\LeadsVoipCallTokenApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LeadsVoipCallTokenGet**](LeadsVoipCallTokenApi.md#LeadsVoipCallTokenGet) | **Get** /leads_voip_call_token/get | 获取网络电话token


# **LeadsVoipCallTokenGet**
> LeadsVoipCallTokenGetResponse LeadsVoipCallTokenGet(ctx, accountId, userId, optional)
获取网络电话token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **userId** | **int64**|  | 
 **optional** | ***LeadsVoipCallTokenApiLeadsVoipCallTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LeadsVoipCallTokenApiLeadsVoipCallTokenGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**LeadsVoipCallTokenGetResponse**](LeadsVoipCallTokenGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

