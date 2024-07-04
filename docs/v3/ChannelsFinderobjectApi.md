# TencentAds\ChannelsFinderobjectApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsFinderobjectGet**](ChannelsFinderobjectApi.md#ChannelsFinderobjectGet) | **Get** /channels_finderobject/get | 获取视频号动态详情


# **ChannelsFinderobjectGet**
> ChannelsFinderobjectGetResponse ChannelsFinderobjectGet(ctx, accountId, exportId, optional)
获取视频号动态详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
  **exportId** | **string**|  | 
 **optional** | ***ChannelsFinderobjectApiChannelsFinderobjectGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsFinderobjectApiChannelsFinderobjectGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ChannelsFinderobjectGetResponse**](ChannelsFinderobjectGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

