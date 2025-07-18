# TencentAds\ChannelsUserpageobjectsApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsUserpageobjectsGet**](ChannelsUserpageobjectsApi.md#ChannelsUserpageobjectsGet) | **Get** /channels_userpageobjects/get | 获取视频号动态列表


# **ChannelsUserpageobjectsGet**
> ChannelsUserpageobjectsGetResponse ChannelsUserpageobjectsGet(ctx, accountId, optional)
获取视频号动态列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ChannelsUserpageobjectsApiChannelsUserpageobjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsUserpageobjectsApiChannelsUserpageobjectsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **finderUsername** | **optional.String**|  | 
 **nickname** | **optional.String**|  | 
 **lastBuffer** | **optional.String**|  | 
 **count** | **optional.Int64**|  | 
 **wechatChannelsAccountId** | **optional.String**|  | 
 **adContext** | [**optional.Interface of AdContext**](AdContext.md)|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ChannelsUserpageobjectsGetResponse**](ChannelsUserpageobjectsGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

