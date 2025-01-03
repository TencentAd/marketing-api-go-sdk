# TencentAds\ChannelsLivenoticeinfoApi

All URIs are relative to *https://sandbox-api.e.qq.com/v3.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsLivenoticeinfoGet**](ChannelsLivenoticeinfoApi.md#ChannelsLivenoticeinfoGet) | **Get** /channels_livenoticeinfo/get | 获取视频号当前的预约直播信息


# **ChannelsLivenoticeinfoGet**
> ChannelsLivenoticeinfoGetResponse ChannelsLivenoticeinfoGet(ctx, accountId, optional)
获取视频号当前的预约直播信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **int64**|  | 
 **optional** | ***ChannelsLivenoticeinfoApiChannelsLivenoticeinfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsLivenoticeinfoApiChannelsLivenoticeinfoGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **finderUsername** | **optional.String**|  | 
 **nickname** | **optional.String**|  | 
 **wechatChannelsAccountId** | **optional.String**|  | 
 **fields** | [**optional.Interface of []string**](string.md)| 返回参数的字段列表 | 

### Return type

[**ChannelsLivenoticeinfoGetResponse**](ChannelsLivenoticeinfoGetResponse.md)

### Authorization

[accessToken](../README.md#accessToken), [nonce](../README.md#nonce), [timestamp](../README.md#timestamp)

### HTTP request headers

 - **Content-Type**: text/plain
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

